package modification

import (
	"github.com/geordanr/go_xwing/attack"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/gamestate"
	"github.com/geordanr/go_xwing/ship"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertResults_AttackerUpTo(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := ConvertResults{
		from:  dice.BLANK,
		to:    dice.CRIT,
		actor: constants.ATTACKER,
		upto:  2,
	}

	attackResults := dice.RollAttackDice(4)
	attackResults[0].SetResult(dice.BLANK)
	attackResults[1].SetResult(dice.BLANK)
	attackResults[2].SetResult(dice.HIT)
	attackResults[3].SetResult(dice.BLANK)
	state.SetAttackResults(&attackResults)

	defenseResults := dice.RollDefenseDice(3)
	defenseResults[0].SetResult(dice.BLANK)
	defenseResults[1].SetResult(dice.FOCUS)
	defenseResults[2].SetResult(dice.EVADE)
	state.SetDefenseResults(&defenseResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, attacker)

	attackResults = *state.AttackResults()
	assert.EqualValues(1, attackResults.Blanks())
	assert.EqualValues(0, attackResults.Focuses())
	assert.EqualValues(1, attackResults.Hits())
	assert.EqualValues(2, attackResults.Crits())
}

func TestConvertResults_AttackerAll(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := ConvertResults{
		from:  dice.HIT,
		to:    dice.FOCUS,
		actor: constants.ATTACKER,
		all:   true,
	}

	attackResults := dice.RollAttackDice(4)
	attackResults[0].SetResult(dice.HIT)
	attackResults[1].SetResult(dice.CRIT)
	attackResults[2].SetResult(dice.HIT)
	attackResults[3].SetResult(dice.BLANK)
	state.SetAttackResults(&attackResults)

	defenseResults := dice.RollDefenseDice(3)
	defenseResults[0].SetResult(dice.BLANK)
	defenseResults[1].SetResult(dice.FOCUS)
	defenseResults[2].SetResult(dice.EVADE)
	state.SetDefenseResults(&defenseResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, attacker)

	attackResults = *state.AttackResults()
	assert.EqualValues(1, attackResults.Blanks())
	assert.EqualValues(2, attackResults.Focuses())
	assert.EqualValues(0, attackResults.Hits())
	assert.EqualValues(1, attackResults.Crits())
}

func TestConvertResults_DefenderUpTo(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := ConvertResults{
		from:  dice.BLANK,
		to:    dice.EVADE,
		actor: constants.DEFENDER,
		upto:  2,
	}

	attackResults := dice.RollAttackDice(4)
	attackResults[0].SetResult(dice.BLANK)
	attackResults[1].SetResult(dice.BLANK)
	attackResults[2].SetResult(dice.HIT)
	attackResults[3].SetResult(dice.BLANK)
	state.SetAttackResults(&attackResults)

	defenseResults := dice.RollDefenseDice(4)
	defenseResults[0].SetResult(dice.BLANK)
	defenseResults[1].SetResult(dice.FOCUS)
	defenseResults[2].SetResult(dice.BLANK)
	defenseResults[3].SetResult(dice.BLANK)
	state.SetDefenseResults(&defenseResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, attacker)

	defenseResults = *state.DefenseResults()
	assert.EqualValues(1, defenseResults.Blanks())
	assert.EqualValues(1, defenseResults.Focuses())
	assert.EqualValues(2, defenseResults.Evades())
}

func TestConvertResults_DefenderAll(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := ConvertResults{
		from:  dice.FOCUS,
		to:    dice.EVADE,
		actor: constants.DEFENDER,
		all:   true,
	}

	attackResults := dice.RollAttackDice(4)
	attackResults[0].SetResult(dice.HIT)
	attackResults[1].SetResult(dice.CRIT)
	attackResults[2].SetResult(dice.HIT)
	attackResults[3].SetResult(dice.BLANK)
	state.SetAttackResults(&attackResults)

	defenseResults := dice.RollDefenseDice(4)
	defenseResults[0].SetResult(dice.FOCUS)
	defenseResults[1].SetResult(dice.FOCUS)
	defenseResults[2].SetResult(dice.BLANK)
	defenseResults[3].SetResult(dice.EVADE)
	state.SetDefenseResults(&defenseResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, attacker)

	defenseResults = *state.DefenseResults()
	assert.EqualValues(1, defenseResults.Blanks())
	assert.EqualValues(0, defenseResults.Focuses())
	assert.EqualValues(3, defenseResults.Evades())
}
