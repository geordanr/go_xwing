package runner

import (
	// "fmt"
	"github.com/geordanr/go_xwing/attack"
	"github.com/geordanr/go_xwing/attack/modification"
	"github.com/geordanr/go_xwing/attack/step"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/gamestate"
	"github.com/geordanr/go_xwing/interfaces"
	"github.com/geordanr/go_xwing/ship"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwinLaserTurret(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 2, 1, 5, 3)
	defender := ship.New("Defender", 3, 3, 3, 0)
	state := gamestate.GameState{}
	mods := map[string][]interfaces.Modification{}

	tlt := modification.TwinLaserTurret{}
	mods["Declare Target"] = []interfaces.Modification{&tlt}

	// Add modifications to record how many times we attacked
	// and how many times the after attacking/defending step triggered.
	compareResultsCounter := new(countSteps)
	mods["Compare Results"] = []interfaces.Modification{
		&modification.CompareResults{},
		compareResultsCounter,
	}

	finalCounter := new(countSteps)
	// Does "after attacking/defending" happen before or after "perform attack twice"?
	mods["Perform Additional Attack"] = []interfaces.Modification{finalCounter}

	state.EnqueueAttack(attack.New(attacker, defender, mods))

	nStates := 1
	runner := New(step.All, nStates)
	output := make(chan interfaces.GameState, nStates)
	go runner.Run(output)

	runner.InjectState(&state)
	<-output

	atkResults := state.AttackResults()
	assert.EqualValues(3, atkResults.Blanks()+atkResults.Focuses()+atkResults.Hits()+atkResults.Crits())
	assert.EqualValues(2, compareResultsCounter.Value())
	assert.EqualValues(1, finalCounter.Value())
}

type countSteps int

func (mod *countSteps) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	*mod++
}

func (mod countSteps) Actor() constants.ModificationActor          { return constants.IGNORE }
func (mod *countSteps) SetActor(actor constants.ModificationActor) {}
func (mod countSteps) String() string                              { return "Count Steps" }
func (mod countSteps) Mods() []interfaces.Modification             { return nil }
func (mod countSteps) Value() int                                  { return int(mod) }
func (mod countSteps) IsSecondaryWeapon() bool                     { return false }
