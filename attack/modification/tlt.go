package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/interfaces"
)

// TwinLaserTurret should be used in the Declare Attack step.
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
	rollAtkDiceMods := mods["Roll Attack Dice"]
	for i, mod := range rollAtkDiceMods {
		if mod.String() == "Roll Attack Dice" {
			rollAtkDiceMods[i] = &rollTLTDice
		}
	}
	if len(rollAtkDiceMods) == 0 {
		// just in case
		mods["Roll Attack Dice"] = []interfaces.Modification{&rollTLTDice}
	}

	// Instead of the default deal damage handler, use our own
	dealDamageMods := mods["Deal Damage"]
	for i, mod := range dealDamageMods {
		switch mod.(type) {
		case *DealDamage:
			dealDamageMods[i] = &DealTLTDamage{}
		}
	}
	if len(dealDamageMods) == 0 {
		// just in case
		mods["Deal Damage"] = []interfaces.Modification{&DealTLTDamage{}}
	}

	// Signal the default Perform Attack Twice handler to activate
	state.SetPerformAttackTwice(true)
}

func (mod TwinLaserTurret) Actor() constants.ModificationActor          { return constants.ATTACKER }
func (mod *TwinLaserTurret) SetActor(actor constants.ModificationActor) {}
func (mod TwinLaserTurret) String() string                              { return "Twin Laser Turret" }
func (mod TwinLaserTurret) IsSecondaryWeapon() bool                     { return true }

// DealTLTDamage is installed by TwinLaserTurret and deals one damage to the defender if the attack hit.
type DealTLTDamage struct{}

func (mod *DealTLTDamage) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	currentAttack := state.CurrentAttack()
	if state.HitsLanded()+state.CritsLanded() > 0 {
		currentAttack.Defender().SufferDamage(1, 0)
	}
}
func (mod DealTLTDamage) Actor() constants.ModificationActor          { return constants.DEFENDER }
func (mod *DealTLTDamage) SetActor(actor constants.ModificationActor) {}
func (mod DealTLTDamage) String() string                              { return "Deal TLT Damage" }
func (mod DealTLTDamage) IsSecondaryWeapon() bool                     { return false }
func (mod DealTLTDamage) IsDamageDealer() bool                        { return true }
