package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/interfaces"
)

// CannotAttackAgain should be used in the After Attacking/Defending step.
// It sets the cannot attack flag on the attacker so that subsequent attacks
// from the attacker immediately terminate.
type CannotAttackAgain struct {
	actor constants.ModificationActor
}

func (mod *CannotAttackAgain) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	currentAttack := state.CurrentAttack()
	currentAttack.Attacker().SetCanAttack(false)
}

func (mod CannotAttackAgain) Actor() constants.ModificationActor          { return constants.ATTACKER }
func (mod *CannotAttackAgain) SetActor(actor constants.ModificationActor) {}
func (mod CannotAttackAgain) String() string                              { return "Cannot Attack Again" }
