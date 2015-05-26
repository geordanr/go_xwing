package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "math/rand"
    "net/http"
    "time"
    "github.com/gocraft/web"
    "github.com/geordanr/go_xwing/combat"
)

type jsonErrorMessage struct {
    Error bool `json:"error"`
    Message string `json:"message"`
}

type Context struct {}

func main() {
    rand.Seed(time.Now().UnixNano())
    router := web.New(Context{}).
	Middleware(web.LoggerMiddleware).
	Get("/", (*Context).Root).
	Post("/api/v1/combat/listVsList", (*Context).ListVsList).
	Error((*Context).Error)

    http.ListenAndServe("localhost:8080", router)
}

func (*Context) Root(rw web.ResponseWriter, req *web.Request) {
    fmt.Fprintf(rw, "<html><body>Monte Carlo is the best Carlo</body></html>")
}

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
	fmt.Printf("Marshal error %s\n", marshalError)
	fmt.Fprintf(rw, `{"error":true,"message":"Error marshaling error message: %s"}`, marshalError)
    } else {
	fmt.Printf("Bytes %s\n", bytes)
	fmt.Fprintf(rw, "%s", bytes)
    }
}
