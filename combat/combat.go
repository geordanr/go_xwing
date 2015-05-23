package combat

import (
    "github.com/geordanr/go_xwing/attack"
    "github.com/geordanr/go_xwing/ship"
    "github.com/geordanr/go_xwing/histogram"
)

type Combat struct {
    attacks []attack.Attack
    results map[ship.Ship]simResult
}
// New takes a function that returns a list of Attacks and constructs a Combat from it.
func New(atkFactory func() []attack.Attack) *Combat {
    cbt := Combat{
	attacks: atkFactory(),
    }

    return &cbt
}

func (cbt Combat) Results() map[ship.Ship]simResult {
    return cbt.results
}
// Simulate creates a combat from the given attack factory and runs it the specified number of times.
func Simulate(atkFactory func() []attack.Attack, iterations int) (*statsByShipName, *resultsByShipName) {

    // Collect ships for stats map
    // Create a combat to analyze the combatants
    cbt := New(atkFactory)
    combatStats := make(statsByShipName)
    combatResults := make(resultsByShipName)
    for _, atk := range(cbt.attacks) {
	if _, prs := combatStats[atk.Attacker.Name]; !prs {
	    combatStats[atk.Attacker.Name] = new(stats)
	    combatResults[atk.Attacker.Name] = new(simResult)
	    combatResults[atk.Attacker.Name].HitHistogram = make(histogram.IntHistogram)
	    combatResults[atk.Attacker.Name].CritHistogram = make(histogram.IntHistogram)
	    combatResults[atk.Attacker.Name].HullHistogram = make(histogram.IntHistogram)
	    combatResults[atk.Attacker.Name].ShieldHistogram = make(histogram.IntHistogram)
	}

	if _, prs := combatStats[atk.Defender.Name]; !prs {
	    combatStats[atk.Defender.Name] = new(stats)
	    combatResults[atk.Defender.Name] = new(simResult)
	    combatResults[atk.Defender.Name].HitHistogram = make(histogram.IntHistogram)
	    combatResults[atk.Defender.Name].CritHistogram = make(histogram.IntHistogram)
	    combatResults[atk.Defender.Name].HullHistogram = make(histogram.IntHistogram)
	    combatResults[atk.Defender.Name].ShieldHistogram = make(histogram.IntHistogram)
	}
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
    combatants := make(map[string]*ship.Ship)
    for _, atk := range(cbt.attacks) {
	// Update combatants
	if _, prs := combatants[atk.Attacker.Name]; !prs {
	    combatants[atk.Attacker.Name] = atk.Attacker
	}
	if _, prs := combatants[atk.Defender.Name]; !prs {
	    combatants[atk.Defender.Name] = atk.Defender
	}

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
    for name, combatant := range(combatants) {
	(*combatStats)[name].HullSum += combatant.Hull
	(*combatStats)[name].HullSumSquares += combatant.Hull * combatant.Hull
	(*combatStats)[name].ShieldSum += combatant.Shields
	(*combatStats)[name].ShieldSumSquares += combatant.Shields * combatant.Shields
	(*combatResults)[name].HullHistogram[int(combatant.Hull)]++
	(*combatResults)[name].ShieldHistogram[int(combatant.Shields)]++
    }
}

