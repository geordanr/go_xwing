package attack

import (
    // "fmt"
    "github.com/geordanr/go_xwing/ship"
)

type Team struct {
    Name string
    Initiative int
    Ships []ship.Ship
}
