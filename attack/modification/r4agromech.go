package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/interfaces"
)

// R4Agromech should be used in the Declare Target or Modify Attack Dice step.
type R4Agromech struct {
	actor constants.ModificationActor
}

func (mod *R4Agromech) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	if ship.FocusTokens() > 0 {
		ship.SpendFocus()
		ship.SetTargetLock(state.CurrentAttack().Defender().Name())
	}
}

func (mod R4Agromech) Actor() constants.ModificationActor          { return constants.ATTACKER }
func (mod *R4Agromech) SetActor(actor constants.ModificationActor) {}
func (mod R4Agromech) String() string                              { return "R4 Agromech" }
