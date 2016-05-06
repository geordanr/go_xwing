package modification

import (
	"fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/interfaces"
	"math"
)

// RollDice rolls dice for the attacker or defender, using their attack
// or agility values as appropriate, and sets the state's dice results.
type RollDice struct {
	actor      constants.ModificationActor
	numDice    uint8
	useNumDice bool
}

func (mod *RollDice) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	var results dice.Results
	var nDice uint8
	currentAttack := state.CurrentAttack()
	if ship == currentAttack.Attacker() {
		if mod.useNumDice {
			nDice = mod.numDice
		} else {
			nDice = uint8(math.Max(0, float64(ship.Attack())+float64(state.AttackDiceModifier())))
		}
		results = dice.RollAttackDice(nDice)
		state.SetAttackResults(&results)
	} else if ship == currentAttack.Defender() {
		if mod.useNumDice {
			nDice = mod.numDice
		} else {
			nDice = uint8(math.Max(0, float64(ship.Agility())+float64(state.DefenseDiceModifier())))
		}
		results = dice.RollDefenseDice(nDice)
		state.SetDefenseResults(&results)
	}
}

func (mod RollDice) Actor() constants.ModificationActor          { return mod.actor }
func (mod *RollDice) SetActor(actor constants.ModificationActor) { mod.actor = actor }
func (mod RollDice) String() string {
	var dieType string
	switch mod.actor {
	case constants.ATTACKER:
		dieType = "Attack"
	case constants.DEFENDER:
		dieType = "Defense"
	default:
		dieType = "Unknown"
	}

	if mod.useNumDice {
		return fmt.Sprintf("Roll %d %s Dice", mod.numDice, dieType)
	} else {
		return fmt.Sprintf("Roll %s Dice", dieType)
	}
}
func (mod RollDice) IsSecondaryWeapon() bool { return false }

// SetNumDice sets the number of dice to be rolled, instead of using the
// attacker's attack value or the defender's agility value modified by
// any other dice count modifiers.
func (mod *RollDice) SetNumDice(numDice uint8) { mod.numDice = numDice; mod.useNumDice = true }
