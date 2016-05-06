package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/interfaces"
)

// DealDamage applies damage to the defender.
type DealDamage struct {
	actor constants.ModificationActor
}

func (mod *DealDamage) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	currentAttack := state.CurrentAttack()
	currentAttack.Defender().SufferDamage(state.HitsLanded(), state.CritsLanded())
}

func (mod DealDamage) Actor() constants.ModificationActor          { return mod.actor }
func (mod *DealDamage) SetActor(actor constants.ModificationActor) { mod.actor = actor }
func (mod DealDamage) String() string                              { return "Deal Damage" }
func (mod DealDamage) IsDamageDealer() bool                        { return true }
