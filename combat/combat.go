package combat

import (
    "github.com/geordanr/go_xwing/attack"
    "github.com/geordanr/go_xwing/ship"
    "github.com/geordanr/go_xwing/histogram"
)

type Combat struct {
    attacks []attack.Attack
    results map[*ship.Ship]simResult
    combatants map[string]*ship.Ship
}
// New takes a function that returns a list of Attacks and constructs a Combat from it.
func New(atkFactory func() []attack.Attack) *Combat {
    cbt := Combat{
	attacks: atkFactory(),
    }

    cbt.combatants = make(map[string]*ship.Ship)
    for _, atk := range(cbt.attacks) {
	// Update combatants
	if _, prs := cbt.combatants[atk.Attacker.Name]; !prs {
	    cbt.combatants[atk.Attacker.Name] = atk.Attacker
	}
	if _, prs := cbt.combatants[atk.Defender.Name]; !prs {
	    cbt.combatants[atk.Defender.Name] = atk.Defender
	}
    }

    return &cbt
}

func (cbt Combat) Results() map[*ship.Ship]simResult {
    return cbt.results
}
// Simulate creates a combat from the given attack factory and runs it the specified number of times.
func Simulate(atkFactory func() []attack.Attack, iterations int) (*statsByShipName, *resultsByShipName) {

    // Collect ships for stats map
    // Create a combat to analyze the combatants
    cbt := New(atkFactory)
    combatStats := make(statsByShipName)
    combatResults := make(resultsByShipName)
    for name := range(cbt.combatants) {
	combatStats[name] = new(stats)
	combatResults[name] = new(simResult)
	combatResults[name].HitHistogram = make(histogram.IntHistogram)
	combatResults[name].CritHistogram = make(histogram.IntHistogram)
	combatResults[name].HullHistogram = make(histogram.IntHistogram)
	combatResults[name].ShieldHistogram = make(histogram.IntHistogram)
    }

    for i := 0; i < iterations; i++ {
	// Create a fresh combat
	cbt = New(atkFactory)
	cbt.Execute(&combatStats, &combatResults)
    }

    for combatant, res := range(combatResults) {
	st := combatStats[combatant]
	st.ComputeStandardDeviations(res)
    }

    return &combatStats, &combatResults
}

// Execute executes all attacks in combat, returns stats per ship.
func (cbt *Combat) Execute(combatStats *statsByShipName, combatResults *resultsByShipName) {
    // Execute actions
    for _, combatant := range(cbt.combatants) {
	combatant.PerformActions()
    }

    for _, atk := range(cbt.attacks) {
	hits, crits := atk.Execute()

	(*combatStats)[atk.Attacker.Name].Iterations++
	(*combatStats)[atk.Attacker.Name].HitSum += hits
	(*combatStats)[atk.Attacker.Name].HitSumSquares += hits * hits
	(*combatStats)[atk.Attacker.Name].CritSum += crits
	(*combatStats)[atk.Attacker.Name].CritSumSquares += crits * crits

	(*combatResults)[atk.Attacker.Name].HitHistogram[int(hits)]++
	(*combatResults)[atk.Attacker.Name].CritHistogram[int(crits)]++
    }
    // Tally up final health
    for name, combatant := range(cbt.combatants) {
	(*combatStats)[name].HullSum += combatant.Hull
	(*combatStats)[name].HullSumSquares += combatant.Hull * combatant.Hull
	(*combatStats)[name].ShieldSum += combatant.Shields
	(*combatStats)[name].ShieldSumSquares += combatant.Shields * combatant.Shields
	(*combatResults)[name].HullHistogram[int(combatant.Hull)]++
	(*combatResults)[name].ShieldHistogram[int(combatant.Shields)]++
    }
}

