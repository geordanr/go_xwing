package modification

import (
	"fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/interfaces"
)

// C-3PO should be used in the Roll Defense Dice step (since it can be Juked).
// It must come after the defense die roll modification.
// It can be instantiated with a guess (default=0).
type C3PO struct {
	actor constants.ModificationActor
	guess uint8
}

func (mod *C3PO) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	results := *state.DefenseResults()
	if results.Evades() == mod.guess {
		evadeDie := new(dice.DefenseDie)
		evadeDie.SetResult(dice.EVADE)
		results = append(results, evadeDie)
		state.SetDefenseResults(&results)
	}
}

func (mod C3PO) Actor() constants.ModificationActor          { return constants.DEFENDER }
func (mod *C3PO) SetActor(actor constants.ModificationActor) {}
func (mod C3PO) String() string                              { return fmt.Sprintf("C-3PO (guess %d)", mod.guess) }
func (mod C3PO) IsSecondaryWeapon() bool                     { return false }
