package main

import (
    "fmt"
    "math"
    "math/rand"
    "time"
    "github.com/geordanr/go_xwing/attack"
    "github.com/geordanr/go_xwing/histogram"
    "github.com/geordanr/go_xwing/ship"
)

type stats struct {
    hitSum uint
    hitSumSquares uint
    critSum uint
    critSumSquares uint
}

type simResults struct {
    HitAverage float64
    HitStddev float64
    HitHistogram histogram.IntHistogram
    CritAverage float64
    CritStddev float64
    CritHistogram histogram.IntHistogram
}

func main() {
    rand.Seed(time.Now().UnixNano())
    iterations := 10000

    stats := new(stats)
    res := new(simResults)
    res.HitHistogram = make(histogram.IntHistogram)
    res.CritHistogram = make(histogram.IntHistogram)
    for i := 0; i < iterations; i++ {
	hits, crits := gunner()
	res.HitHistogram[int(hits)]++
	res.CritHistogram[int(crits)]++

	stats.hitSum += hits
	stats.hitSumSquares += hits * hits
	stats.critSum += crits
	stats.critSumSquares += crits * crits
    }

    res.HitAverage = float64(stats.hitSum) /  float64(iterations)
    res.HitStddev = math.Sqrt((float64(stats.hitSumSquares) / float64(iterations)) - math.Pow(res.HitAverage, 2))

    res.CritAverage = float64(stats.critSum) /  float64(iterations)
    res.CritStddev = math.Sqrt((float64(stats.critSumSquares) / float64(iterations)) - math.Pow(res.CritAverage, 2))
    fmt.Printf("Hits: %-.3f (stddev=%-.3f)\n", res.HitAverage, res.HitStddev)
    fmt.Println(res.HitHistogram)
    fmt.Printf("Crits: %-.3f (stddev=%-.3f)\n", res.CritAverage, res.CritStddev)
    fmt.Println(res.CritHistogram)
}

func gunner() (hits, crits uint) {
    atk := attack.Attack{
	Attacker: &ship.Ship{FocusTokens: 0},
	NumAttackDice: 3,
	AttackerModifications: []attack.Modification{
	    // attack.Modifications["Predator (high PS)"],
	    // attack.Modifications["Chiraneau"],
	    attack.Modifications["Accuracy Corrector"],
	    // attack.Modifications["Offensive Focus"],
	},

	Defender: &ship.Ship{FocusTokens: 1},
	NumDefenseDice: 0,
	DefenderModifications: []attack.Modification{
	    attack.Modifications["Defensive Focus"],
	},
    }

    hits, crits = atk.Execute()

    if hits + crits == 0 {
	gunnerAtk := atk.Copy()

	h, c := gunnerAtk.Execute()
	hits += h
	crits += c
    }

    return hits, crits
}
