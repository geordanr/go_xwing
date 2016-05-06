package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/interfaces"
)

// DeclareTarget is the default handler for the Declare Target step.
// If the attacker cannot attack or if the defender is destroyed, it
// immediately terminates the attack without resolving.
type DeclareTarget struct {
	actor constants.ModificationActor
}

func (mod *DeclareTarget) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	currentAttack := state.CurrentAttack()
	if !currentAttack.Attacker().CanAttack() || !currentAttack.Defender().IsAlive() {
		state.SetNextAttackStep("")
	}
}

func (mod DeclareTarget) Actor() constants.ModificationActor          { return constants.ATTACKER }
func (mod *DeclareTarget) SetActor(actor constants.ModificationActor) {}
func (mod DeclareTarget) String() string                              { return "Declare Target" }
func (mod DeclareTarget) IsSecondaryWeapon() bool                     { return false }
