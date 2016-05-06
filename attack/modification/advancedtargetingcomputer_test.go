package modification

import (
	"github.com/geordanr/go_xwing/attack"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/gamestate"
	"github.com/geordanr/go_xwing/ship"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdvancedTargetingComputer_HasLock(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := AdvancedTargetingComputer{}

	attacker.SetTargetLock("Defender")
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
	assert.EqualValues(1, attackResults.Crits())
}

func TestAdvancedTargetingComputer_HasNoLock(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := AdvancedTargetingComputer{}

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
}
