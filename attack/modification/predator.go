package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/dice/filters"
	"github.com/geordanr/go_xwing/interfaces"
)

// Predator should be used in the Modify Attack Dice step.
type Predator struct {
	actor constants.ModificationActor
}

func (mod *Predator) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	currentAttack := state.CurrentAttack()
	defender := currentAttack.Defender()
	nReroll := uint(1)
	if defender.Skill() <= 2 {
		nReroll = 2
	}
	results := state.AttackResults()
	if ship.FocusTokens() > 0 {
		results.RerollUpto(nReroll, filters.Blanks)
	} else {
		results.RerollUpto(nReroll, filters.BlanksAndFocuses)
	}
	state.SetAttackResults(results)
}

func (mod Predator) Actor() constants.ModificationActor          { return constants.ATTACKER }
func (mod *Predator) SetActor(actor constants.ModificationActor) {}
func (mod Predator) String() string                              { return "Predator" }
