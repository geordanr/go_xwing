package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/interfaces"
)

// ConvertResults is a utility modification for changing die results.
// It can be instantiated with what results to convert to what, up to the amount given.  If all matching results should be converted, set all to true.
// The result pool being modified depends on the actor.
// This can be transient; set transient to true if the instantiated modification should not be included in an attack copy.
type ConvertResults struct {
	actor     constants.ModificationActor
	from      dice.Result
	to        dice.Result
	upto      uint
	all       bool
	transient bool
}

func (mod *ConvertResults) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	var results *dice.Results
	switch mod.actor {
	case constants.ATTACKER:
		results = state.AttackResults()
	case constants.DEFENDER:
		results = state.DefenseResults()
	}
	if mod.all {
		results.ConvertAll(mod.from, mod.to)
	} else {
		results.ConvertUpto(mod.upto, mod.from, mod.to)
	}
	// switch mod.actor {
	// case constants.ATTACKER:
	// 	state.SetAttackResults(results)
	// case constants.DEFENDER:
	// 	state.SetDefenseResults(results)
	// }
}

func (mod ConvertResults) Actor() constants.ModificationActor          { return mod.actor }
func (mod *ConvertResults) SetActor(actor constants.ModificationActor) { mod.actor = actor }
func (mod ConvertResults) String() string                              { return "Convert Results" }
func (mod ConvertResults) IsTransient() bool                           { return mod.transient }
