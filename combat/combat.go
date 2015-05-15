package combat

import (
    "math"
    "github.com/geordanr/go_xwing/attack"
    "github.com/geordanr/go_xwing/ship"
    "github.com/geordanr/go_xwing/histogram"
)

type Combat struct {
    attacks []attack.Attack
    results map[ship.Ship]simResult
}

type stats struct {
    HitSum uint
    HitSumSquares uint
    CritSum uint
    CritSumSquares uint
}

type simResult struct {
    HitAverage float64
    HitStddev float64
    HitHistogram histogram.IntHistogram
    CritAverage float64
    CritStddev float64
    CritHistogram histogram.IntHistogram
}

type statsByShip map[*ship.Ship]*stats
type resultsByShip map[*ship.Ship]*simResult

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

// Execute executes all attacks in combat, returns stats per ship.
func (cbt *Combat) Execute(combatStats *statsByShip, combatResults *resultsByShip) {
    for _, atk := range(cbt.attacks) {
	hits, crits := atk.Execute()

	(*combatStats)[atk.Attacker].HitSum += hits
	(*combatStats)[atk.Attacker].HitSumSquares += hits * hits
	(*combatStats)[atk.Attacker].CritSum += crits
	(*combatStats)[atk.Attacker].CritSumSquares += crits * crits

	(*combatResults)[atk.Attacker].HitHistogram[int(hits)]++
	(*combatResults)[atk.Attacker].CritHistogram[int(crits)]++
    }
}

func (cbt *Combat) Simulate(iterations int) (*statsByShip, *resultsByShip) {
    // Collect ships for stats map
    combatStats := make(statsByShip)
    combatResults := make(resultsByShip)
    for _, atk := range(cbt.attacks) {
	if _, prs := combatStats[atk.Attacker]; !prs {
	    combatStats[atk.Attacker] = new(stats)
	    combatResults[atk.Attacker] = new(simResult)
	    combatResults[atk.Attacker].HitHistogram = make(histogram.IntHistogram)
	    combatResults[atk.Attacker].CritHistogram = make(histogram.IntHistogram)
	}
    }

    for i := 0; i < iterations; i++ {
	cbt.Execute(&combatStats, &combatResults)
    }

    for attacker, res := range(combatResults) {
	st := combatStats[attacker]

	res.HitAverage = float64(st.HitSum) /  float64(iterations)
	res.HitStddev = math.Sqrt((float64(st.HitSumSquares) / float64(iterations)) - math.Pow(res.HitAverage, 2))

	res.CritAverage = float64(st.CritSum) /  float64(iterations)
	res.CritStddev = math.Sqrt((float64(st.CritSumSquares) / float64(iterations)) - math.Pow(res.CritAverage, 2))
    }

    return &combatStats, &combatResults
}
