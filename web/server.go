package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "math/rand"
    "net/http"
    "sort"
    "time"
    "github.com/gocraft/web"
    "github.com/geordanr/go_xwing/attack"
    "github.com/geordanr/go_xwing/combat"
    "github.com/geordanr/go_xwing/ship"
)

type jsonErrorMessage struct {
    Error bool `json:"error"`
    Message string `json:"message"`
}

type genericList struct {
    Data []string `json:"data"`
}

type Context struct {}

func main() {
    rand.Seed(time.Now().UnixNano())
    router := web.New(Context{}).
	Middleware(web.LoggerMiddleware).
	Get("/", (*Context).Root).
	Get("/api/v1/ships", (*Context).Ships).
	Get("/api/v1/actions", (*Context).Actions).
	Get("/api/v1/modifications", (*Context).Modifications).
	Post("/api/v1/combat/listVsList", (*Context).ListVsList).
	Error((*Context).Error)

    http.ListenAndServe("localhost:8080", router)
}

func (*Context) Root(rw web.ResponseWriter, req *web.Request) {
    fmt.Fprintf(rw, "<html><body>Monte Carlo is the best Carlo</body></html>")
}

// ListVersusList receives JSON POST data specifying parameters for simulating a list vs. list combat round.
func (*Context) ListVsList(rw web.ResponseWriter, req *web.Request) {
    rw.Header().Set("Content-Type", "application/json")

    defer req.Body.Close()
    body, err := ioutil.ReadAll(req.Body)
    if err != nil { panic(err) }

    shipStats, shipResults, err := combat.SimulateFromJSON(body)
    if err != nil { panic(err) }

    data, err := combat.SimResultsToJSON(shipStats, shipResults)
    if err != nil { panic(err) }

    fmt.Fprintf(rw, `{"data": "%s"}`, data)
}

func (*Context) Error(rw web.ResponseWriter, req *web.Request, err interface{}) {
    rw.WriteHeader(http.StatusInternalServerError)
    rw.Header().Set("Content-Type", "application/json")
    msg := jsonErrorMessage{
	Error: true,
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

    shiplist := make([]string, 0, len(ship.ShipFactory))
    for shipname, _ := range(ship.ShipFactory) {
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

// Actions returns a JSON list of supported actions.
func (*Context) Actions(rw web.ResponseWriter, req *web.Request) {
    rw.Header().Set("Content-Type", "application/json")

    actionlist := make([]string, 0, len(ship.Actions))
    for actionname, _ := range(ship.Actions) {
	actionlist = append(actionlist, actionname)
    }
    sort.Strings(actionlist)

    d := genericList{Data: actionlist}

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

    modist := make([]string, 0, len(attack.Modifications))
    for modame, _ := range(attack.Modifications) {
	modist = append(modist, modame)
    }
    sort.Strings(modist)

    d := genericList{Data: modist}

    bytes, marshalError := json.Marshal(d)
    if marshalError != nil {
	fmt.Fprintf(rw, `{"error":true,"message":"Error marshaling error message: %s"}`, marshalError)
    } else {
	fmt.Fprintf(rw, "%s", bytes)
    }
}

