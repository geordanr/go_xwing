# I STILL DON'T KNOW WHAT I'M DOING

[![Build Status](https://travis-ci.org/geordanr/go_xwing.svg?branch=master)](https://travis-ci.org/geordanr/go_xwing)

A second attempt at this Go stuff, as an experiment in code organization.  (No claim is made on whether this organization is good...)

## But seriously, what's going on here?

This is a monte carlo simulator for simulating a single round of combat in X-Wing.

## Installing

    go get github.com/geordanr/go_xwing
    go get ./...

## Running Tests

    go test ./...

## Running

Human-readable console output:

    go run $GOPATH/src/github.com/geordanr/go_xwing/jsonsimdemo/main.go -shipjson ships.json -simjson jsonsimdemo/sample.json

Web server accepting JSON and outputting JSON:

    go run $GOPATH/src/github.com/geordanr/go_xwing/web/server.go -shipjson ships.json

Then send a request:

    curl -H 'Content-Type: application-json' -d @jsonsimdemo/sample.json http://localhost:8080/api/v1/sim
