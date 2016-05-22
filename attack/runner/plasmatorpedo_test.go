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
	"reflect"
	"testing"
)

func TestPlasmaTorpedo_missed(t *testing.T) {
	assert := assert.New(t)

	makeState := func() *gamestate.GameState {
		attacker := ship.New("Attacker", 0, 2, 1, 5, 3)
		defender := ship.New("Defender", 0, 3, 3, 3, 3)
		state := gamestate.GameState{}
		state.SetCombatants(map[string]interfaces.Ship{
			"Attacker": attacker,
			"Defender": defender,
		})
		mods := map[string][]interfaces.Modification{}

		defender.SetEvadeTokens(4)

		mods["Declare Target"] = []interfaces.Modification{new(modification.PlasmaTorpedo)}
		mods["Modify Defense Dice"] = []interfaces.Modification{new(modification.SpendEvade)}

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

	hullLeft := map[uint]bool{}
	for i := 0; i < nStates; i++ {
		state := <-output
		defender := state.Combatants()["Defender"]
		hullLeft[defender.Shields()] = true
	}

	// No damage should ever be done
	keys := reflect.ValueOf(hullLeft).MapKeys()
	assert.Equal(1, len(keys))
	assert.Contains(hullLeft, uint(3))
}

func TestPlasmaTorpedo_hits(t *testing.T) {
	assert := assert.New(t)

	makeState := func() *gamestate.GameState {
		attacker := ship.New("Attacker", 0, 1, 1, 5, 3)
		defender := ship.New("Defender", 0, 3, 0, 3, 3)
		state := gamestate.GameState{}
		state.SetCombatants(map[string]interfaces.Ship{
			"Attacker": attacker,
			"Defender": defender,
		})
		mods := map[string][]interfaces.Modification{}

		mods["Declare Target"] = []interfaces.Modification{new(modification.PlasmaTorpedo)}
		mods["Modify Defense Dice"] = []interfaces.Modification{new(modification.SpendEvade)}

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

	hullLeft := map[uint]bool{}
	for i := 0; i < nStates; i++ {
		state := <-output
		defender := state.Combatants()["Defender"]
		hullLeft[defender.Shields()] = true
	}

	// Either no shield damage was done, or at least 2 shield damage was done
	assert.NotContains(hullLeft, uint(2))
}

func TestPlasmaTorpedo_hitDoesntDamageHull(t *testing.T) {
	assert := assert.New(t)

	makeState := func() *gamestate.GameState {
		attacker := ship.New("Attacker", 0, 1, 1, 5, 3)
		defender := ship.New("Defender", 0, 3, 0, 10, 1)
		state := gamestate.GameState{}
		state.SetCombatants(map[string]interfaces.Ship{
			"Attacker": attacker,
			"Defender": defender,
		})
		mods := map[string][]interfaces.Modification{}

		mods["Declare Target"] = []interfaces.Modification{new(modification.PlasmaTorpedo)}
		mods["Modify Defense Dice"] = []interfaces.Modification{new(modification.SpendEvade)}

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

	minHullLeft := uint(11)
	for i := 0; i < nStates; i++ {
		state := <-output
		defender := state.Combatants()["Defender"]
		if defender.Hull() < minHullLeft {
			minHullLeft = defender.Hull()
		}
	}

	// Can only leave at most 7 hull
	assert.True(minHullLeft > 6)
}
