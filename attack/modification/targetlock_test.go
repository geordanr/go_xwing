package modification

import (
	"github.com/geordanr/go_xwing/attack"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/gamestate"
	"github.com/geordanr/go_xwing/ship"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpendTargetLock_HasNoLock(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := SpendTargetLock{}

	attackResults := dice.RollAttackDice(3)
	attackResults[0].SetResult(dice.BLANK)
	attackResults[1].SetResult(dice.FOCUS)
	attackResults[2].SetResult(dice.HIT)
	state.SetAttackResults(&attackResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, attacker)

	attackResults = *state.AttackResults()
	assert.EqualValues(1, attackResults.Blanks())
	assert.EqualValues(1, attackResults.Focuses())
	assert.EqualValues(1, attackResults.Hits())
	assert.EqualValues(0, attackResults.Crits())

	for _, result := range attackResults {
		assert.False(result.Rerolled())
	}
}

func TestSpendTargetLock_NoFocusToken(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := SpendTargetLock{}

	attacker.SetTargetLock("Defender")
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

func TestSpendTargetLock_HasFocusToken(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := SpendTargetLock{}

	attacker.SetTargetLock("Defender")
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
