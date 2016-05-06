package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/interfaces"
)

// AccuracyCorrector should be used in the Modify Attack Dice step.
type AccuracyCorrector struct {
	actor constants.ModificationActor
}

func (mod *AccuracyCorrector) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	results := state.AttackResults()
	for _, result := range *results {
		result.SetResult(dice.CANCELED)
	}
	for i := 0; i < 2; i++ {
		d := dice.AttackDie{}
		d.SetResult(dice.HIT)
		*results = append(*results, &d)
	}
	state.SetAttackResults(results)
}

func (mod AccuracyCorrector) Actor() constants.ModificationActor          { return constants.ATTACKER }
func (mod *AccuracyCorrector) SetActor(actor constants.ModificationActor) {}
func (mod AccuracyCorrector) String() string                              { return "Accuracy Corrector" }
