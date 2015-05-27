package main

import (
    "flag"
    "fmt"
    "math/rand"
    // "path/filepath"
    // "runtime"
    "os"
    "time"
    "github.com/geordanr/go_xwing/combat"
)

// Demo of reading JSON, simulating it, and outputting JSON results
func main() {
    // _, thisfile, _, _ := runtime.Caller(0)
    // thisdir := filepath.Dir(thisfile)
    // shipStats, shipResults, err := combat.SimulateFromJSONPath(filepath.Join(thisdir, "..", "combat", "sample.json"))
    // if err != nil {
	// panic(err)
    // }

    jsonpath := flag.String("jsonpath", "", "Path to JSON file to parse")
    flag.Parse()

    if *jsonpath == "" {
	fmt.Println("Specify the path to the JSON file to parse.")
	os.Exit(1)
    }

    rand.Seed(time.Now().UnixNano())
    shipStats, shipResults, err := combat.SimulateFromJSONPath(*jsonpath)
    if err != nil { panic(err) }
    data, err := combat.SimResultsToJSON(shipStats, shipResults)
    if err != nil { panic(err) }

    fmt.Println(string(data))
}
