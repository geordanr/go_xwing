package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/interfaces"
)

// Gunner should be used in the Perform Additional Attack step.
type Gunner struct {
	actor constants.ModificationActor
}

func (mod *Gunner) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	currentAttack := state.CurrentAttack()
	if state.AttackMissed() {
		var mods, newMods []interfaces.Modification
		// Copy current attack parameters
		newAtk := currentAttack.Copy()

		// Use the default dice roller (primary)
		mods = newAtk.Modifications()["Roll Attack Dice"]
		newMods = []interfaces.Modification{}
		for i, mod := range mods {
			switch mod.(type) {
			case *RollDice:
				mods[i] = &RollDice{}
			}
		}
		newAtk.Modifications()["Roll Attack Dice"] = newMods

		// Remove Gunner mod
		mods = newAtk.Modifications()["Perform Additional Attack"]
		newMods = []interfaces.Modification{}
		for _, mod := range mods {
			switch mod.(type) {
			case *Gunner:
				// don't copy
			default:
				newMods = append(newMods, mod)
			}
		}
		newAtk.Modifications()["Perform Additional Attack"] = newMods

		// Add mod to prevent attacks after gunner
		mods = newAtk.Modifications()["After Attacking/Defending"]
		newAtk.Modifications()["After Attacking/Defending"] = append(mods, &CannotAttackAgain{})

		// Enqueue attack copy
		state.SetNextAttack(newAtk)
	}
}

func (mod Gunner) Actor() constants.ModificationActor          { return constants.ATTACKER }
func (mod *Gunner) SetActor(actor constants.ModificationActor) {}
func (mod Gunner) String() string                              { return "Gunner" }
