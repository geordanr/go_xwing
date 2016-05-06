package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/interfaces"
)

// CrackShot should be used in the Compare Results step, before the default handler.
type CrackShot struct {
	actor constants.ModificationActor
}

func (mod *CrackShot) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	results := state.DefenseResults()
	results.ConvertUpto(1, dice.EVADE, dice.CANCELED)
	// TODO crack shot only once per sim
}

func (mod CrackShot) Actor() constants.ModificationActor          { return constants.ATTACKER }
func (mod *CrackShot) SetActor(actor constants.ModificationActor) {}
func (mod CrackShot) String() string                              { return "Crack Shot" }
func (mod CrackShot) IsSecondaryWeapon() bool                     { return false }
