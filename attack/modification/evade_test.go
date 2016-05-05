package modification

import (
	"github.com/geordanr/go_xwing/attack"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/gamestate"
	"github.com/geordanr/go_xwing/ship"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpendEvade_NoTokens(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := SpendEvade{}

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
	mod.ModifyState(&state, defender)

	assert.EqualValues(0, defender.EvadeTokens())
	assert.EqualValues(1, state.DefenseResults().Blanks())
	assert.EqualValues(1, state.DefenseResults().Focuses())
	assert.EqualValues(1, state.DefenseResults().Evades())
}

func TestSpendEvade_HasTokens(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := SpendEvade{}

	attackResults := dice.RollAttackDice(3)
	attackResults[0].SetResult(dice.BLANK)
	attackResults[1].SetResult(dice.HIT)
	attackResults[2].SetResult(dice.HIT)
	state.SetAttackResults(&attackResults)

	defender.SetEvadeTokens(1)
	defenseResults := dice.RollDefenseDice(3)
	defenseResults[0].SetResult(dice.BLANK)
	defenseResults[1].SetResult(dice.FOCUS)
	defenseResults[2].SetResult(dice.EVADE)
	state.SetDefenseResults(&defenseResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, defender)

	assert.EqualValues(0, defender.EvadeTokens())
	assert.EqualValues(1, state.DefenseResults().Blanks())
	assert.EqualValues(1, state.DefenseResults().Focuses())
	assert.EqualValues(2, state.DefenseResults().Evades())
}

func TestSpendEvade_NoNeedToSpend(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := SpendEvade{}

	attackResults := dice.RollAttackDice(3)
	attackResults[0].SetResult(dice.BLANK)
	attackResults[1].SetResult(dice.BLANK)
	attackResults[2].SetResult(dice.HIT)
	state.SetAttackResults(&attackResults)

	defender.SetEvadeTokens(1)
	defenseResults := dice.RollDefenseDice(3)
	defenseResults[0].SetResult(dice.BLANK)
	defenseResults[1].SetResult(dice.FOCUS)
	defenseResults[2].SetResult(dice.EVADE)
	state.SetDefenseResults(&defenseResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, defender)

	assert.EqualValues(1, defender.EvadeTokens())
	assert.EqualValues(1, state.DefenseResults().Blanks())
	assert.EqualValues(1, state.DefenseResults().Focuses())
	assert.EqualValues(1, state.DefenseResults().Evades())
}

func TestSpendEvade_SpendEnough(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := SpendEvade{}

	attackResults := dice.RollAttackDice(3)
	attackResults[0].SetResult(dice.BLANK)
	attackResults[1].SetResult(dice.HIT)
	attackResults[2].SetResult(dice.HIT)
	state.SetAttackResults(&attackResults)

	defender.SetEvadeTokens(3)
	defenseResults := dice.RollDefenseDice(3)
	defenseResults[0].SetResult(dice.BLANK)
	defenseResults[1].SetResult(dice.FOCUS)
	defenseResults[2].SetResult(dice.BLANK)
	state.SetDefenseResults(&defenseResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, defender)

	assert.EqualValues(1, defender.EvadeTokens())
	assert.EqualValues(2, state.DefenseResults().Blanks())
	assert.EqualValues(1, state.DefenseResults().Focuses())
	assert.EqualValues(2, state.DefenseResults().Evades())
}
