package main

// import _ "net/http/pprof"
// import "net/http"
// import "log"

import (
	"flag"
	"fmt"
	"math/rand"
	// "path/filepath"
	// "runtime"
	"github.com/geordanr/go_xwing/serialization"
	"github.com/geordanr/go_xwing/ship"
	"github.com/geordanr/go_xwing/stats"
	"os"
	"time"
)

type shipStats struct {
	hull    stats.Integers
	shields stats.Integers
}

func (s *shipStats) update(ship *ship.Ship) {
	s.hull.Update(int(ship.Hull()))
	s.shields.Update(int(ship.Shields()))
}

func (s shipStats) String() string {
	return fmt.Sprintf("Hull\t: average=%2.3f stddev=%2.3f\nShields\t: average=%2.3f stddev=%2.3f\n", s.hull.Average(), s.hull.Stddev(), s.shields.Average(), s.shields.Stddev())
}

// Demo of reading JSON, simulating it, and outputting JSON results
func main() {
	// _, thisfile, _, _ := runtime.Caller(0)
	// thisdir := filepath.Dir(thisfile)
	// shipStats, shipResults, err := combat.SimulateFromsimjson(filepath.Join(thisdir, "..", "combat", "sample.json"))
	// if err != nil {
	// 	panic(err)
	// }

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
	cbtStats := map[string]*shipStats{}

	for {
		state, more := <-output
		if !more {
			break
		}
		for name, cbt := range state.Combatants() {
			s, exists := cbtStats[name]
			if !exists {
				cbtStats[name] = new(shipStats)
				s = cbtStats[name]
			}
			s.update(cbt.(*ship.Ship))
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
