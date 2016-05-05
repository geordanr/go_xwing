package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/interfaces"
)

// TwinLaserTurret should be used in the Declare Attack Dice step.
// It sets the number of dice to roll for the attacker and sets the
// perform attack twice flag to be used by the default Perform Attack
// Twice modification.
type TwinLaserTurret struct {
	actor constants.ModificationActor
}

func (mod *TwinLaserTurret) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	currentAttack := state.CurrentAttack()
	mods := currentAttack.Modifications()

	// Set the number of attack dice to roll
	rollTLTDice := RollDice{}
	rollTLTDice.SetNumDice(3)
	rollTLTDice.SetActor(constants.ATTACKER)
	// Instead of the default dice roller, use our roller
	mods["Roll Attack Dice"] = []interfaces.Modification{&rollTLTDice}

	// Signal the default Perform Attack Twice handler to activate
	state.SetPerformAttackTwice(true)
}

func (mod TwinLaserTurret) Actor() constants.ModificationActor          { return constants.ATTACKER }
func (mod *TwinLaserTurret) SetActor(actor constants.ModificationActor) {}
func (mod TwinLaserTurret) String() string                              { return "Twin Laser Turret" }
