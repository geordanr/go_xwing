package main

import (
    "fmt"
    "path/filepath"
    "runtime"
    "github.com/geordanr/go_xwing/combat"
)

// Demo of reading JSON, simulating it, and outputting JSON results
func main() {
    _, thisfile, _, _ := runtime.Caller(0)
    thisdir := filepath.Dir(thisfile)
    shipStats, shipResults, err := combat.SimulateFromJSONPath(filepath.Join(thisdir, "..", "combat", "sample.json"))
    if err != nil {
	panic(err)
    }

    data, err := combat.SimResultsToJSON(shipStats, shipResults)
    if err != nil {
	panic(err)
    }

    fmt.Println(string(data))
}
