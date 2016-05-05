package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/interfaces"
)

// SpendEvade spends just enough evade tokens to cancel hits.
// Applies only to the defender.
type SpendEvade struct {
	actor constants.ModificationActor
}

func (mod *SpendEvade) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	currentAttack := state.CurrentAttack()
	defender := currentAttack.Defender()
	if defender.EvadeTokens() > 0 {
		attackResults := state.AttackResults()
		for defender.EvadeTokens() > 0 && (attackResults.Hits()+attackResults.Crits()) > (state.DefenseResults().Evades()) {
			defender.SpendEvade()
			evadeDie := new(dice.DefenseDie)
			evadeDie.SetResult(dice.EVADE)
			results := *state.DefenseResults()
			results = append(results, evadeDie)
			state.SetDefenseResults(&results)
		}
	}
}

func (mod SpendEvade) Actor() constants.ModificationActor          { return mod.actor }
func (mod *SpendEvade) SetActor(actor constants.ModificationActor) { mod.actor = actor }

func (mod SpendEvade) String() string {
	return "Spend Evade"
}
