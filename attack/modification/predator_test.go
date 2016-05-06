package modification

import (
	"github.com/geordanr/go_xwing/attack"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/gamestate"
	"github.com/geordanr/go_xwing/ship"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPredator_LowPSNoFocusToken(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := Predator{}

	attackResults := dice.RollAttackDice(3)
	attackResults[0].SetResult(dice.BLANK)
	attackResults[1].SetResult(dice.FOCUS)
	attackResults[2].SetResult(dice.HIT)
	state.SetAttackResults(&attackResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, attacker)

	attackResults = *state.AttackResults()
	assert.True(attackResults[0].Rerolled())
	assert.True(attackResults[1].Rerolled())
	assert.False(attackResults[2].Rerolled())
}

func TestPredator_LowPSWithFocusToken(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := Predator{}

	attacker.SetFocusTokens(1)
	attackResults := dice.RollAttackDice(3)
	attackResults[0].SetResult(dice.BLANK)
	attackResults[1].SetResult(dice.FOCUS)
	attackResults[2].SetResult(dice.HIT)
	state.SetAttackResults(&attackResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, attacker)

	attackResults = *state.AttackResults()
	assert.True(attackResults[0].Rerolled())
	assert.False(attackResults[1].Rerolled())
	assert.False(attackResults[2].Rerolled())
}

func TestPredator_HighPSNoFocusToken(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0, 0)
	defender := ship.New("Defender", 9, 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := Predator{}

	attackResults := dice.RollAttackDice(3)
	attackResults[0].SetResult(dice.BLANK)
	attackResults[1].SetResult(dice.FOCUS)
	attackResults[2].SetResult(dice.HIT)
	state.SetAttackResults(&attackResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, attacker)

	attackResults = *state.AttackResults()
	nRerolled := 0
	for _, res := range attackResults {
		if res.Rerolled() {
			nRerolled++
		}
	}
	assert.Equal(1, nRerolled)

}

func TestPredator_HighPSWithFocusToken(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0, 0)
	defender := ship.New("Defender", 9, 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := Predator{}

	attacker.SetFocusTokens(1)
	attackResults := dice.RollAttackDice(4)
	attackResults[0].SetResult(dice.BLANK)
	attackResults[1].SetResult(dice.BLANK)
	attackResults[2].SetResult(dice.FOCUS)
	attackResults[3].SetResult(dice.HIT)
	state.SetAttackResults(&attackResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, attacker)

	attackResults = *state.AttackResults()
	nRerolled := 0
	for _, res := range attackResults {
		if res.Rerolled() {
			nRerolled++
		}
	}
	assert.Equal(1, nRerolled)
}
