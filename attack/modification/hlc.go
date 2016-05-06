package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/interfaces"
)

/*
HeavyLaserCannon should be used in the Declare Target step.
It replaces the default dice roller with a custom roller that rolls 4
dice, and also installs a modification in the Modify Attack Dice step
that immediately converts crits to hits.
*/
type HeavyLaserCannon struct {
	actor constants.ModificationActor
}

func (mod *HeavyLaserCannon) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	currentAttack := state.CurrentAttack()
	mods := currentAttack.Modifications()

	// Replace default die roller
	rollHLCDice := RollDice{}
	rollHLCDice.SetNumDice(4)
	rollHLCDice.SetActor(constants.ATTACKER)
	rollAtkDiceMods := mods["Roll Attack Dice"]
	for i, mod := range rollAtkDiceMods {
		if mod.String() == "Roll Attack Dice" {
			rollAtkDiceMods[i] = &rollHLCDice
		}
	}
	if len(rollAtkDiceMods) == 0 {
		// just in case
		mods["Roll Attack Dice"] = []interfaces.Modification{&rollHLCDice}
	}

	// Install mod to convert attack results first
	// TODO - gunner etc. needs to be able to remove this.
	stepMods, exists := mods["Modify Attack Dice"]
	if !exists {
		stepMods = []interfaces.Modification{}
	}
	newMods := []interfaces.Modification{
		&ConvertResults{
			actor: constants.ATTACKER,
			from:  dice.CRIT,
			to:    dice.HIT,
			all:   true,
		},
	}
	mods["Modify Attack Dice"] = append(newMods, stepMods...)
}

func (mod HeavyLaserCannon) Actor() constants.ModificationActor          { return constants.ATTACKER }
func (mod *HeavyLaserCannon) SetActor(actor constants.ModificationActor) {}
func (mod HeavyLaserCannon) String() string                              { return "Heavy Laser Cannon" }
func (mod HeavyLaserCannon) IsSecondaryWeapon() bool                     { return true }
