package ship

import "github.com/geordanr/go_xwing/dice"

type Action interface {
    Perform(*Ship)
    String() string
}

////////////////////////////

type landoCrew struct {}
func (landoCrew) Perform(s *Ship) {
    landoResults := dice.RollDefenseDice(2)

    s.EvadeTokens += uint(landoResults.Evades())
    s.FocusTokens += uint(landoResults.Focuses())
}
func (landoCrew) String() string { return "Lando (Crew)" }

var Actions map[string]Action = map[string]Action{
    "Lando (Crew)": new(landoCrew),
}
