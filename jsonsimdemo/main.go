package main

// import _ "net/http/pprof"
// import "net/http"
// import "log"
// import "github.com/pkg/profile"

import (
	"flag"
	"fmt"
	"math/rand"
	// "path/filepath"
	// "runtime"
	"github.com/geordanr/go_xwing/serialization"
	"github.com/geordanr/go_xwing/ship"
	"github.com/geordanr/go_xwing/shipstats"
	"os"
	"time"
)

// Demo of reading JSON, simulating it, and outputting JSON results
func main() {
	// _, thisfile, _, _ := runtime.Caller(0)
	// thisdir := filepath.Dir(thisfile)
	// ship.Stats, shipResults, err := combat.SimulateFromsimjson(filepath.Join(thisdir, "..", "combat", "sample.json"))
	// if err != nil {
	// 	panic(err)
	// }

	// defer profile.Start(profile.ProfilePath("."), profile.CPUProfile).Stop()

	rand.Seed(time.Now().UnixNano())

	shipjson, simjson := parseArgs()

	shipFactory, err := serialization.ShipsFromJSONPath(*shipjson)
	if err != nil {
		panic(err)
	}

	output, err := serialization.FromJSONPath(*simjson, shipFactory)
	if err != nil {
		panic(err)
	}

	// map from ship ID to stats struct
	cbtStats := map[string]*shipstats.Stats{}

	for {
		state, more := <-output
		if !more {
			break
		}
		for name, cbt := range state.Combatants() {
			s, exists := cbtStats[name]
			if !exists {
				cbtStats[name] = shipstats.New()
				s = cbtStats[name]
			}
			// fmt.Printf("Update ship stats for %s (%p) %s\n", cbt.Name(), cbt, cbt)
			s.Update(cbt.(*ship.Ship))
		}
	}

	for name, s := range cbtStats {
		fmt.Println()
		fmt.Println(name)
		fmt.Println("---")
		fmt.Println(s)
	}

	// log.Println(http.ListenAndServe("localhost:8080", nil))
}

func parseArgs() (*string, *string) {
	shipjson := flag.String("shipjson", "", "Path to JSON file of ship data")
	simjson := flag.String("simjson", "", "Path to JSON file of sim parameters")

	flag.Parse()

	if *shipjson == "" {
		fmt.Println("Specify the path to the ship JSON file.")
		os.Exit(1)
	}

	if *simjson == "" {
		fmt.Println("Specify the path to the sim parameter JSON file.")
		os.Exit(1)
	}

	return shipjson, simjson
}
