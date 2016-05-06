package runner

import (
	// "fmt"
	"github.com/geordanr/go_xwing/attack"
	"github.com/geordanr/go_xwing/attack/modification"
	"github.com/geordanr/go_xwing/attack/step"
	"github.com/geordanr/go_xwing/gamestate"
	"github.com/geordanr/go_xwing/interfaces"
	"github.com/geordanr/go_xwing/ship"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeavyLaserCannon_NeverCrits(t *testing.T) {
	assert := assert.New(t)

	makeState := func() *gamestate.GameState {
		attacker := ship.New("Attacker", 0, 2, 1, 5, 3)
		defender := ship.New("Defender", 0, 3, 0, 3, 0)
		state := gamestate.GameState{}
		mods := map[string][]interfaces.Modification{}

		hlc := modification.HeavyLaserCannon{}
		mods["Declare Target"] = []interfaces.Modification{&hlc}

		state.EnqueueAttack(attack.New(attacker, defender, mods))

		return &state
	}

	nStates := 1000
	runner := New(step.All, 8)
	output := make(chan interfaces.GameState, 8)
	go runner.Run(output)

	for i := 0; i < nStates; i++ {
		go func() {
			runner.InjectState(makeState())
		}()
	}

	var nCrits uint8 = 0
	for i := 0; i < nStates; i++ {
		state := <-output
		results := state.AttackResults()
		nCrits += results.Crits()
	}

	assert.EqualValues(0, nCrits)
}

// There is a very slim chance that this could fail!
func TestHeavyLaserCannon_GunnerCouldCrit(t *testing.T) {
	assert := assert.New(t)

	makeState := func() *gamestate.GameState {
		attacker := ship.New("Attacker", 0, 3, 1, 5, 3)
		defender := ship.New("Defender", 0, 3, 0, 3, 0)
		state := gamestate.GameState{}
		mods := map[string][]interfaces.Modification{}

		hlc := modification.HeavyLaserCannon{}
		gunner := modification.Gunner{}
		mods["Declare Target"] = []interfaces.Modification{&hlc}
		mods["Perform Additional Attack"] = []interfaces.Modification{&gunner}

		defender.SetEvadeTokens(2)
		state.EnqueueAttack(attack.New(attacker, defender, mods))

		return &state
	}

	nStates := 1000
	runner := New(step.All, 8)
	output := make(chan interfaces.GameState, 8)
	go runner.Run(output)

	for i := 0; i < nStates; i++ {
		go func() {
			runner.InjectState(makeState())
		}()
	}

	var nCrits uint8 = 0
	for i := 0; i < nStates; i++ {
		state := <-output
		results := state.AttackResults()
		nCrits += results.Crits()
	}

	assert.True(nCrits > 0)
}
