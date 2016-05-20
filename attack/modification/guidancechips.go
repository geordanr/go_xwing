package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/interfaces"
)

// GuidanceChips should be used in the Modify Attack Dice step.
// Assumes you're firing ordnance (up to you to ensure that's the case).
type GuidanceChips struct {
	actor constants.ModificationActor
}

func (mod *GuidanceChips) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	var newResult dice.Result
	results := state.AttackResults()
	if ship.Attack() < 3 {
		newResult = dice.HIT
		if results.Blanks() > 0 {
			results.ConvertUpto(1, dice.BLANK, newResult)
		} else {
			results.ConvertUpto(1, dice.FOCUS, newResult)
		}
	} else {
		newResult = dice.CRIT
		if results.Blanks() > 0 {
			results.ConvertUpto(1, dice.BLANK, newResult)
		} else if results.Focuses() > 0 {
			results.ConvertUpto(1, dice.FOCUS, newResult)
		} else {
			results.ConvertUpto(1, dice.HIT, newResult)
		}
	}
}

func (mod GuidanceChips) Actor() constants.ModificationActor          { return constants.ATTACKER }
func (mod *GuidanceChips) SetActor(actor constants.ModificationActor) {}
func (mod GuidanceChips) String() string                              { return "Guidance Chips" }
