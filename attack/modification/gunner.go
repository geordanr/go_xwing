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

		// Remove any declared secondary weapons
		mods = newAtk.Modifications()["Declare Target"]
		newMods = []interfaces.Modification{}
		for _, mod := range mods {
			_, ok := mod.(SecondaryWeapon)
			if !ok {
				newMods = append(newMods, mod)
			}
		}
		newAtk.Modifications()["Declare Target"] = newMods

		// Use the default dice roller (primary)
		mods = newAtk.Modifications()["Roll Attack Dice"]
		newMods = []interfaces.Modification{}
		for _, mod := range mods {
			_, ok := mod.(*RollDice)
			if ok {
				newMods = append(newMods, &RollDice{actor: constants.ATTACKER})
			} else {
				newMods = append(newMods, mod)
			}
		}
		newAtk.Modifications()["Roll Attack Dice"] = newMods

		// Use the default damage dealer
		mods = newAtk.Modifications()["Deal Damage"]
		newMods = []interfaces.Modification{}
		for _, mod := range mods {
			switch mod.(type) {
			case DamageDealer:
				_, ok := mod.(*DealDamage)
				if !ok {
					// It's not the default damage dealer
					mod = new(DealDamage)
				}
				// since we can't fallthrough in a type switch
				newMods = append(newMods, mod)
			default:
				newMods = append(newMods, mod)
			}
		}
		newAtk.Modifications()["Perform Additional Attack"] = newMods

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
