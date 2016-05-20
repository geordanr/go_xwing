package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/interfaces"
)

/*
PlasmaTorpedo should be used in the Declare Target step.
It replaces the default dice roller with a custom roller that rolls 4 dice,
and installs a modification that removes an extra shield at the Deal Damage step.
*/
type PlasmaTorpedo struct {
	actor constants.ModificationActor
}

func (mod *PlasmaTorpedo) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
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

	// Install mod to add plasma damage
	stepMods := mods["Deal Damage"]
	if len(stepMods) == 0 {
		mods["Deal Damage"] = append(stepMods, new(DealDamage))
		stepMods = mods["Deal Damage"]
	}
	mods["Deal Damage"] = append(stepMods, &applyPlasmaDamage{transient: true})
}

func (mod PlasmaTorpedo) Actor() constants.ModificationActor          { return constants.ATTACKER }
func (mod *PlasmaTorpedo) SetActor(actor constants.ModificationActor) {}
func (mod PlasmaTorpedo) String() string                              { return "Plasma Torpedo" }
func (mod PlasmaTorpedo) IsSecondaryWeapon() bool                     { return true }

// helper modification to apply plasma damage if it hits

type applyPlasmaDamage struct {
	actor     constants.ModificationActor
	transient bool
}

func (mod *applyPlasmaDamage) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	defender := state.CurrentAttack().Defender()
	if !state.AttackMissed() && defender.Shields() > 0 {
		defender.SufferDamage(1, 0)
	}
}

func (mod applyPlasmaDamage) Actor() constants.ModificationActor          { return constants.ATTACKER }
func (mod *applyPlasmaDamage) SetActor(actor constants.ModificationActor) {}
func (mod applyPlasmaDamage) String() string                              { return "Apply Plasma Damage" }
func (mod applyPlasmaDamage) IsTransient() bool                           { return mod.transient }
