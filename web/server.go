package main

import (
	// "path/filepath"
	// "runtime"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"sort"
	"time"

	"github.com/geordanr/go_xwing/attack/modification"
	"github.com/geordanr/go_xwing/attack/step"
	"github.com/geordanr/go_xwing/serialization"
	"github.com/geordanr/go_xwing/ship"
	"github.com/geordanr/go_xwing/shipstats"
	"github.com/gocraft/web"
)

type jsonErrorMessage struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type genericList struct {
	Data []string `json:"data"`
}

// map ship name to map of collected stat (e.g. "hull", "shields")
type simResults map[string]map[string]shipStatJSONSchema

// map stat to average, stddev
type shipStatsJSONSchema map[string]shipStatJSONSchema
type shipStatJSONSchema struct {
	Histogram [][]float64 `json:"histogram"`
	Average   float64     `json:"average"`
	Stddev    float64     `json:"stddev"`
}

type Context struct{}

func main() {
	rand.Seed(time.Now().UnixNano())

	shipjson, port := parseArgs()
	factory, err := serialization.ShipsFromJSONPath(*shipjson)
	shipFactory = factory
	if err != nil {
		panic(err)
	}

	currentRoot, _ := os.Getwd()
	router := web.New(Context{}).
		Middleware(web.LoggerMiddleware).
		Middleware(corsMiddleware).
		Middleware(web.StaticMiddleware(path.Join(currentRoot, "web", "public"), web.StaticOption{IndexFile: "index.html"})).
		// Get("/", (*Context).Root).
		Get("/api/v1/modifications", (*Context).Modifications).
		Get("/api/v1/ships", (*Context).Ships).
		Get("/api/v1/steps", (*Context).Steps).
		Post("/api/v1/sim", (*Context).Simulate).
		Error((*Context).Error)

	log.Println("Listening on port", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), router)
}

func parseArgs() (*string, *int) {
	shipjson := flag.String("shipjson", "", "Path to JSON file of ship data")
	port := flag.Int("port", 80, "Port to listen on")

	flag.Parse()

	if *shipjson == "" {
		fmt.Println("Specify the path to the ship JSON file.")
		os.Exit(1)
	}

	return shipjson, port
}

var shipFactory map[string]func(string, uint) *ship.Ship

func corsMiddleware(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	next(rw, req)
}

func (*Context) Root(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprintf(rw, "<html><body>Monte Carlo is the best Carlo</body></html>")
}

// Simulate receives JSON POST data specifying parameters for simulating a list vs. list combat round.
func (*Context) Simulate(rw web.ResponseWriter, req *web.Request) {
	rw.Header().Set("Content-Type", "application/json")

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	output, err := serialization.FromJSON(body, shipFactory)
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

	results := make(simResults)
	for name, s := range cbtStats {
		results[name] = make(shipStatsJSONSchema)
		for statName, combinedStats := range s.StatMap() {
			results[name][statName] = shipStatJSONSchema{
				Average:   combinedStats.Stats.Average(),
				Stddev:    combinedStats.Stats.Stddev(),
				Histogram: combinedStats.Histogram.NormalizedHighchartsSeries(),
			}
		}
	}

	resultsJson, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(rw, string(resultsJson))
}

func (*Context) Error(rw web.ResponseWriter, req *web.Request, err interface{}) {
	rw.WriteHeader(http.StatusInternalServerError)
	rw.Header().Set("Content-Type", "application/json")
	msg := jsonErrorMessage{
		Error:   true,
		Message: err.(error).Error(),
	}
	bytes, marshalError := json.Marshal(msg)
	if marshalError != nil {
		fmt.Fprintf(rw, `{"error":true,"message":"Error marshaling error message: %s"}`, marshalError)
	} else {
		fmt.Fprintf(rw, "%s", bytes)
	}
}

// Ships returns a JSON list of supported ships.
func (*Context) Ships(rw web.ResponseWriter, req *web.Request) {
	rw.Header().Set("Content-Type", "application/json")

	shiplist := make([]string, 0, len(shipFactory))
	for shipname, _ := range shipFactory {
		shiplist = append(shiplist, shipname)
	}
	sort.Strings(shiplist)

	d := genericList{Data: shiplist}

	bytes, marshalError := json.Marshal(d)
	if marshalError != nil {
		fmt.Fprintf(rw, `{"error":true,"message":"Error marshaling error message: %s"}`, marshalError)
	} else {
		fmt.Fprintf(rw, "%s", bytes)
	}
}

// Modifications returns a JSON list of supported modifications.
func (*Context) Modifications(rw web.ResponseWriter, req *web.Request) {
	rw.Header().Set("Content-Type", "application/json")

	modList := make([]string, 0, len(modification.All))
	for modName, _ := range modification.All {
		modList = append(modList, modName)
	}
	sort.Strings(modList)

	d := genericList{Data: modList}

	bytes, marshalError := json.Marshal(d)
	if marshalError != nil {
		fmt.Fprintf(rw, `{"error":true,"message":"Error marshaling error message: %s"}`, marshalError)
	} else {
		fmt.Fprintf(rw, "%s", bytes)
	}
}

// Steps returns a JSON list of attack steps in order.
func (*Context) Steps(rw web.ResponseWriter, req *web.Request) {
	rw.Header().Set("Content-Type", "application/json")

	stepList := make([]string, 0, len(step.All)-1) // not including the start step
	s := step.All["__START__"]
	for s.Next() != "" {
		s = step.All[s.Next()]
		stepList = append(stepList, s.Name())
	}

	d := genericList{Data: stepList}

	bytes, marshalError := json.Marshal(d)
	if marshalError != nil {
		fmt.Fprintf(rw, `{"error":true,"message":"Error marshaling error message: %s"}`, marshalError)
	} else {
		fmt.Fprintf(rw, "%s", bytes)
	}
}
