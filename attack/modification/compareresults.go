package modification

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/interfaces"
	"math"
)

// CompareResults examines results after modifications
// and sets the attackMissed flag appropriately.
type CompareResults struct {
	actor constants.ModificationActor
}

func (mod *CompareResults) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	hits := uint(state.AttackResults().Hits())
	crits := uint(state.AttackResults().Crits())
	evades := uint(state.DefenseResults().Evades())
	// fmt.Printf("%s: Compare: hits=%d, crits=%d, evades=%d\n", state.CurrentAttack(), hits, crits, evades)
	// Spend evade results on hits first
	evadesSpentOnHits := uint(math.Min(float64(hits), float64(evades)))
	// fmt.Printf("Evades spent on hits: %d\n", evadesSpentOnHits)
	hits -= evadesSpentOnHits
	evades -= evadesSpentOnHits
	// fmt.Printf("Hits now %d\nEvades now %d\n", hits, evades)
	// Then crits
	crits = uint(math.Max(0, float64(crits-evades)))
	// fmt.Printf("Crits now %d\n", crits)
	state.SetHitsLanded(hits)
	state.SetCritsLanded(crits)
	// fmt.Println("Landed", hits, "hits and", crits, "crits")
	state.SetAttackMissed(hits+crits == 0)
}

func (mod CompareResults) Actor() constants.ModificationActor          { return mod.actor }
func (mod *CompareResults) SetActor(actor constants.ModificationActor) { mod.actor = actor }
func (mod CompareResults) String() string                              { return "Compare Results" }
func (mod CompareResults) IsSecondaryWeapon() bool                     { return false }
