package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/interfaces"
)

type PerformAttackTwice struct {
	actor constants.ModificationActor
}

func (mod *PerformAttackTwice) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	if state.PerformAttackTwice() {
		state.ResetTransientState()
		// fmt.Println("Would set next step to 'Roll Attack Dice'")
	}
}

func (mod PerformAttackTwice) Actor() constants.ModificationActor          { return constants.ATTACKER }
func (mod *PerformAttackTwice) SetActor(actor constants.ModificationActor) {}
func (mod PerformAttackTwice) String() string                              { return "Perform Attack Twice" }
