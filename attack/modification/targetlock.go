package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/dice/filters"
	"github.com/geordanr/go_xwing/interfaces"
)

type SpendTargetLock struct {
	actor constants.ModificationActor
}

func (mod *SpendTargetLock) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	currentAttack := state.CurrentAttack()
	results := state.AttackResults()
	// only spend target lock if we need to
	if ship.TargetLock() == currentAttack.Defender().Name() {
		var spentLock bool
		// don't spend if there are only eyes to reroll and we have focus
		if ship.FocusTokens() > 0 && results.Blanks() > 0 {
			spentLock = true
			results.Reroll(filters.Blanks)
			state.SetAttackResults(results)
		} else if ship.FocusTokens() == 0 && (results.Blanks() > 0 || results.Focuses() > 0) {
			spentLock = true
			results.Reroll(filters.BlanksAndFocuses)
			state.SetAttackResults(results)
		}

		if spentLock {
			ship.SetTargetLock("")
		}
	}

}

func (mod SpendTargetLock) Actor() constants.ModificationActor          { return constants.ATTACKER }
func (mod *SpendTargetLock) SetActor(actor constants.ModificationActor) {}

func (mod SpendTargetLock) String() string {
	return "Spend Target Lock"
}
