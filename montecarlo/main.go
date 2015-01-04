package main

import (
    "fmt"
    "flag"
    "io/ioutil"
    "log"
    "math"
    "math/rand"
    "time"
    "github.com/geordanr/go_xwing/montecarlo/scenario"
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
    CritAverage float64
    CritStddev float64
}

func dieOrDice(n int) (s string) {
    s = "die"
    if n > 1 {
	s = "dice"
    }
    return
}

// Simulate runs the scenario in the JSON iteration times and returns the results.
func SimulateJSON(scenarioJSON []byte, iterations uint) (res simResults) {
    s, err := scenario.ScenarioFromJSON(scenarioJSON)
    if err != nil {
	log.Fatal(err)
    }
    fmt.Printf("Running %d iterations of:\n", iterations)
    fmt.Printf("\t%d attack %s, %d focus token(s)\n", len(*s.AttackResults), dieOrDice(len(*s.AttackResults)), s.NumAttackerFocus)
    for i := range(s.DefenderModifiesAttackDice) {
	fmt.Printf("\t... %s\n", s.DefenderModifiesAttackDice[i])
    }
    for i := range(s.AttackerModifiesAttackDice) {
	fmt.Printf("\t... %s\n", s.AttackerModifiesAttackDice[i])
    }

    fmt.Printf("\t%d defense %s, %d focus token(s), %d evade token(s)\n", len(*s.DefenseResults), dieOrDice(len(*s.DefenseResults)), s.NumDefenderFocus, s.NumDefenderEvade)
    for i := range(s.AttackerModifiesDefenseDice) {
	fmt.Printf("\t... %s\n", s.AttackerModifiesDefenseDice[i])
    }
    for i := range(s.DefenderModifiesDefenseDice) {
	fmt.Printf("\t... %s\n", s.DefenderModifiesDefenseDice[i])
    }

    stats := new(stats)
    var i uint
    for i = 0; i < iterations; i++ {
	s, err = scenario.ScenarioFromJSON(scenarioJSON)
	if err != nil {
	    log.Fatal(err)
	}
	s.Run()
	hits, crits := s.CompareResults()
	stats.hitSum += hits
	stats.hitSumSquares += hits * hits
	stats.critSum += crits
	stats.critSumSquares += crits * crits
    }

    res.HitAverage = float64(stats.hitSum) /  float64(iterations)
    res.HitStddev = math.Sqrt((float64(stats.hitSumSquares) / float64(iterations)) - math.Pow(res.HitAverage, 2))

    res.CritAverage = float64(stats.critSum) /  float64(iterations)
    res.CritStddev = math.Sqrt((float64(stats.critSumSquares) / float64(iterations)) - math.Pow(res.CritAverage, 2))

    return
}

func main() {
    rand.Seed(time.Now().UnixNano())
    numIterations := flag.Int("iter", 1000, "Number of iterations to run")
    jsonPath := flag.String("json", "", "Path to JSON file to parse")

    flag.Parse()

    if *jsonPath == "" {
	log.Fatal("No JSON path given (use --json)")
    }

    log.Printf("Reading %s", *jsonPath)
    data, err := ioutil.ReadFile(*jsonPath)
    if err != nil {
	log.Fatal(err)
    }

    res := SimulateJSON(data, uint(*numIterations))

    fmt.Printf("Hits: %-.3f (stddev=%-.3f)\n", res.HitAverage, res.HitStddev)
    fmt.Printf("Crits: %-.3f (stddev=%-.3f)\n", res.CritAverage, res.CritStddev)
}
