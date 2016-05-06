package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/interfaces"
)

// SufferDamage applies damage to the defender.
type SufferDamage struct {
	actor constants.ModificationActor
}

func (mod *SufferDamage) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	currentAttack := state.CurrentAttack()
	currentAttack.Defender().SufferDamage(state.HitsLanded() + state.CritsLanded())
}

func (mod SufferDamage) Actor() constants.ModificationActor          { return mod.actor }
func (mod *SufferDamage) SetActor(actor constants.ModificationActor) { mod.actor = actor }
func (mod SufferDamage) String() string                              { return "Suffer Damage" }
func (mod SufferDamage) IsSecondaryWeapon() bool                     { return false }
