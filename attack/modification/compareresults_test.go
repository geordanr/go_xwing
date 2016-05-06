package modification

import (
	"github.com/geordanr/go_xwing/attack"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/gamestate"
	"github.com/geordanr/go_xwing/ship"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareResults_NoHits(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 2, 3, 3, 0)
	defender := ship.New("Defender", 0, 3, 2, 3, 2)
	state := gamestate.GameState{}
	mod := CompareResults{}

	attackResults := dice.RollAttackDice(3)
	attackResults[0].SetResult(dice.BLANK)
	attackResults[1].SetResult(dice.FOCUS)
	attackResults[2].SetResult(dice.BLANK)
	state.SetAttackResults(&attackResults)

	defenseResults := dice.RollDefenseDice(3)
	defenseResults[0].SetResult(dice.BLANK)
	defenseResults[1].SetResult(dice.FOCUS)
	defenseResults[2].SetResult(dice.BLANK)
	state.SetDefenseResults(&defenseResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, attacker)

	assert.EqualValues(0, state.HitsLanded())
	assert.EqualValues(0, state.CritsLanded())
	assert.True(state.AttackMissed())
}

func TestCompareResults_AttackMissed(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 2, 3, 3, 0)
	defender := ship.New("Defender", 0, 3, 2, 3, 2)
	state := gamestate.GameState{}
	mod := CompareResults{}

	attackResults := dice.RollAttackDice(3)
	attackResults[0].SetResult(dice.BLANK)
	attackResults[1].SetResult(dice.HIT)
	attackResults[2].SetResult(dice.HIT)
	state.SetAttackResults(&attackResults)

	defenseResults := dice.RollDefenseDice(3)
	defenseResults[0].SetResult(dice.BLANK)
	defenseResults[1].SetResult(dice.FOCUS)
	defenseResults[2].SetResult(dice.EVADE)
	state.SetDefenseResults(&defenseResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, attacker)

	assert.EqualValues(1, state.HitsLanded())
	assert.EqualValues(0, state.CritsLanded())
	assert.False(state.AttackMissed())
}

func TestCompareResults_ShieldBeforeHull(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 2, 3, 3, 0)
	defender := ship.New("Defender", 0, 3, 2, 3, 2)
	state := gamestate.GameState{}
	mod := CompareResults{}

	attackResults := dice.RollAttackDice(3)
	attackResults[0].SetResult(dice.HIT)
	attackResults[1].SetResult(dice.HIT)
	attackResults[2].SetResult(dice.CRIT)
	state.SetAttackResults(&attackResults)

	defenseResults := dice.RollDefenseDice(3)
	defenseResults[0].SetResult(dice.BLANK)
	defenseResults[1].SetResult(dice.FOCUS)
	defenseResults[2].SetResult(dice.BLANK)
	state.SetDefenseResults(&defenseResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, attacker)

	assert.EqualValues(2, state.HitsLanded())
	assert.EqualValues(1, state.CritsLanded())
}
