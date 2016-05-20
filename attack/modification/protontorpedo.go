package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/interfaces"
)

/*
ProtonTorpedo should be used in the Declare Target step.
It replaces the default dice roller with a custom roller that rolls 4 dice,
and installs a modification that converts a focus at the Modify Attack Dice step.
*/
type ProtonTorpedo struct {
	actor constants.ModificationActor
}

func (mod *ProtonTorpedo) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	currentAttack := state.CurrentAttack()
	mods := currentAttack.Modifications()

	// Replace default die roller
	rollPlasmaDice := RollDice{}
	rollPlasmaDice.SetNumDice(4)
	rollPlasmaDice.SetActor(constants.ATTACKER)
	rollAtkDiceMods := mods["Roll Attack Dice"]
	for i, mod := range rollAtkDiceMods {
		if mod.String() == "Roll Attack Dice" {
			rollAtkDiceMods[i] = &rollPlasmaDice
		}
	}
	if len(rollAtkDiceMods) == 0 {
		// just in case
		mods["Roll Attack Dice"] = []interfaces.Modification{&rollPlasmaDice}
	}

	// Install mod to convert attack results first
	// TODO - gunner etc. needs to be able to remove this.
	stepMods, exists := mods["Modify Attack Dice"]
	if !exists {
		stepMods = []interfaces.Modification{}
	}
	newMods := []interfaces.Modification{
		&ConvertResults{
			actor:     constants.ATTACKER,
			from:      dice.FOCUS,
			to:        dice.CRIT,
			upto:      1,
			transient: true,
		},
	}
	mods["Modify Attack Dice"] = append(newMods, stepMods...)
}

func (mod ProtonTorpedo) Actor() constants.ModificationActor          { return constants.ATTACKER }
func (mod *ProtonTorpedo) SetActor(actor constants.ModificationActor) {}
func (mod ProtonTorpedo) String() string                              { return "Proton Torpedo" }
func (mod ProtonTorpedo) IsSecondaryWeapon() bool                     { return true }
