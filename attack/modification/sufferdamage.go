package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/interfaces"
)

// SufferDamage examines results after modifications, deals damage
// if necessary, and sets the attackMissed flag appropriately.
type SufferDamage struct {
	actor constants.ModificationActor
}

func (mod *SufferDamage) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	currentAttack := state.CurrentAttack()
	currentAttack.Defender().SufferDamage(state.HitsLanded() + state.CritsLanded())
}

func (mod SufferDamage) Actor() constants.ModificationActor          { return mod.actor }
func (mod *SufferDamage) SetActor(actor constants.ModificationActor) { mod.actor = actor }
func (mod SufferDamage) String() string                              { return "Compare Results" }
