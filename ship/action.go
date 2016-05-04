package ship

import "github.com/geordanr/go_xwing/dice"

// Actions are functions performed once per combat before attacks are executed.
// Useful for things that cannot be represented by simple changes to ship stats.
type Action interface {
	Perform(*Ship)
	String() string
}

var Actions map[string]Action = map[string]Action{
	"Lando (Crew)": new(landoCrew),
}

////////////////////////////

type landoCrew struct{}

func (landoCrew) Perform(s *Ship) {
	landoResults := dice.RollDefenseDice(2)

	s.evadeTokens += uint(landoResults.Evades())
	s.focusTokens += uint(landoResults.Focuses())
}
func (landoCrew) String() string { return "Lando (Crew)" }
