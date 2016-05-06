package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/interfaces"
)

// AdvancedTargetingComputer should be used in the Modify Attack Dice step.
// This does not prevent target locks from being spent afterward in the Modify Attack Dice step.  So don't do it!
type AdvancedTargetingComputer struct {
	actor constants.ModificationActor
}

func (mod *AdvancedTargetingComputer) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	results := state.AttackResults()
	if ship.TargetLock() == state.CurrentAttack().Defender().Name() {
		d := dice.AttackDie{}
		d.SetResult(dice.CRIT)
		*results = append(*results, &d)
	}
	state.SetAttackResults(results)
}

func (mod AdvancedTargetingComputer) Actor() constants.ModificationActor          { return constants.ATTACKER }
func (mod *AdvancedTargetingComputer) SetActor(actor constants.ModificationActor) {}
func (mod AdvancedTargetingComputer) String() string                              { return "Advanced Targeting Computer" }
