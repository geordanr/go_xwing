package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/interfaces"
	"math"
)

// RollDice rolls dice for the attacker or defender, using their attack
// or agility values as appropriate, and sets the state's dice results.
type RollDice struct {
	actor constants.ModificationActor
}

func (mod *RollDice) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	var results dice.Results
	currentAttack := state.CurrentAttack()
	if ship == currentAttack.Attacker() {
		nDice := uint8(math.Max(0, float64(ship.Attack())+float64(state.AttackDiceModifier())))
		results = dice.RollAttackDice(nDice)
		state.SetAttackResults(&results)
	} else if ship == currentAttack.Defender() {
		nDice := uint8(math.Max(0, float64(ship.Agility())+float64(state.DefenseDiceModifier())))
		results = dice.RollDefenseDice(nDice)
		state.SetDefenseResults(&results)
	}
}

func (mod RollDice) Actor() constants.ModificationActor          { return mod.actor }
func (mod *RollDice) SetActor(actor constants.ModificationActor) { mod.actor = actor }
func (mod RollDice) String() string                              { return "Roll Dice" }
