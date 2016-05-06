package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/interfaces"
)

// CrackShot should be used in the Compare defenseResults step, before the default handler.
type CrackShot struct {
	actor constants.ModificationActor
	used  bool
}

func (mod *CrackShot) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	if mod.used {
		return
	}
	attackResults := state.AttackResults()
	defenseResults := state.DefenseResults()
	if defenseResults.Evades() > 0 && (attackResults.Hits()+attackResults.Crits()) > 0 {
		mod.used = true
		defenseResults.ConvertUpto(1, dice.EVADE, dice.CANCELED)
	}
}

func (mod CrackShot) Actor() constants.ModificationActor          { return constants.ATTACKER }
func (mod *CrackShot) SetActor(actor constants.ModificationActor) {}
func (mod CrackShot) String() string                              { return "Crack Shot" }
func (mod CrackShot) IsSecondaryWeapon() bool                     { return false }
