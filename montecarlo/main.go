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

// Simulate runs the scenario in the JSON iteration times and returns the results.
func SimulateJSON(scenarioJSON []byte, iterations uint) (res simResults) {
    log.Printf("Running %d iterations...", iterations)

    stats := new(stats)
    var i uint
    for i = 0; i < iterations; i++ {
	s, err := scenario.ScenarioFromJSON(scenarioJSON)
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

    fmt.Printf("Hits: %-.2f (stddev=%-.2f)\n", res.HitAverage, res.HitStddev)
    fmt.Printf("Crits: %-.2f (stddev=%-.2f)\n", res.CritAverage, res.CritStddev)
}
