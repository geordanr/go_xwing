package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/interfaces"
)

type SpendFocus struct {
	actor constants.ModificationActor
}

func (mod *SpendFocus) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	if ship.FocusTokens() > 0 {
		currentAttack := state.CurrentAttack()
		if ship == currentAttack.Attacker() {
			results := state.AttackResults()
			if results.Focuses() > 0 {
				// fmt.Println(ship.Name(), "spent focus")
				ship.SpendFocus()
				results.ConvertAll(dice.FOCUS, dice.HIT)
			}
		} else if ship == currentAttack.Defender() {
			atkResults := state.AttackResults()
			defResults := state.DefenseResults()
			if defResults.Focuses() > 0 && (atkResults.Hits()+atkResults.Crits()) > defResults.Evades() {
				// fmt.Println(ship.Name(), "spent focus")
				ship.SpendFocus()
				defResults.ConvertAll(dice.FOCUS, dice.EVADE)
			}
		}
	}
}

func (mod SpendFocus) Actor() constants.ModificationActor          { return mod.actor }
func (mod *SpendFocus) SetActor(actor constants.ModificationActor) { mod.actor = actor }
func (mod SpendFocus) String() string                              { return "Spend Focus Token" }
