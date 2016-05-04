package modification

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/geordanr/go_xwing/attack"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/gamestate"
	"github.com/geordanr/go_xwing/ship"
)

func TestSpendFocus_OnAttackHasTokens(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := SpendFocus{}

	attacker.SetFocusTokens(2)
	attackResults := dice.RollAttackDice(3)
	attackResults[0].SetResult(dice.BLANK)
	attackResults[1].SetResult(dice.FOCUS)
	attackResults[2].SetResult(dice.HIT)
	state.SetAttackResults(&attackResults)

	defender.SetFocusTokens(3)
	defenseResults := dice.RollDefenseDice(3)
	defenseResults[0].SetResult(dice.BLANK)
	defenseResults[1].SetResult(dice.FOCUS)
	defenseResults[2].SetResult(dice.EVADE)
	state.SetDefenseResults(&defenseResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, attacker)

	assert.EqualValues(1, attacker.FocusTokens())
	assert.EqualValues(1, attackResults.Blanks())
	assert.EqualValues(0, attackResults.Focuses())
	assert.EqualValues(2, attackResults.Hits())
	assert.EqualValues(0, attackResults.Crits())

	assert.EqualValues(3, defender.FocusTokens())
	assert.EqualValues(1, defenseResults.Blanks())
	assert.EqualValues(1, defenseResults.Focuses())
	assert.EqualValues(1, defenseResults.Evades())
}

func TestSpendFocus_OnAttackHasNoTokens(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := SpendFocus{}

	attacker.SetFocusTokens(0)
	attackResults := dice.RollAttackDice(3)
	attackResults[0].SetResult(dice.BLANK)
	attackResults[1].SetResult(dice.FOCUS)
	attackResults[2].SetResult(dice.HIT)
	state.SetAttackResults(&attackResults)

	defender.SetFocusTokens(3)
	defenseResults := dice.RollDefenseDice(3)
	defenseResults[0].SetResult(dice.BLANK)
	defenseResults[1].SetResult(dice.FOCUS)
	defenseResults[2].SetResult(dice.EVADE)
	state.SetDefenseResults(&defenseResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, attacker)

	assert.EqualValues(0, attacker.FocusTokens())
	assert.EqualValues(1, attackResults.Blanks())
	assert.EqualValues(1, attackResults.Focuses())
	assert.EqualValues(1, attackResults.Hits())
	assert.EqualValues(0, attackResults.Crits())

	assert.EqualValues(3, defender.FocusTokens())
	assert.EqualValues(1, defenseResults.Blanks())
	assert.EqualValues(1, defenseResults.Focuses())
	assert.EqualValues(1, defenseResults.Evades())
}

func TestSpendFocus_OnAttackHasNoFocusResults(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := SpendFocus{}

	attacker.SetFocusTokens(2)
	attackResults := dice.RollAttackDice(3)
	attackResults[0].SetResult(dice.BLANK)
	attackResults[1].SetResult(dice.CRIT)
	attackResults[2].SetResult(dice.HIT)
	state.SetAttackResults(&attackResults)

	defender.SetFocusTokens(3)
	defenseResults := dice.RollDefenseDice(3)
	defenseResults[0].SetResult(dice.BLANK)
	defenseResults[1].SetResult(dice.FOCUS)
	defenseResults[2].SetResult(dice.EVADE)
	state.SetDefenseResults(&defenseResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, attacker)

	assert.EqualValues(2, attacker.FocusTokens())
	assert.EqualValues(1, attackResults.Blanks())
	assert.EqualValues(0, attackResults.Focuses())
	assert.EqualValues(1, attackResults.Hits())
	assert.EqualValues(1, attackResults.Crits())

	assert.EqualValues(3, defender.FocusTokens())
	assert.EqualValues(1, defenseResults.Blanks())
	assert.EqualValues(1, defenseResults.Focuses())
	assert.EqualValues(1, defenseResults.Evades())
}

func TestSpendFocus_OnDefenseHasTokens(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := SpendFocus{}

	attacker.SetFocusTokens(2)
	attackResults := dice.RollAttackDice(3)
	attackResults[0].SetResult(dice.BLANK)
	attackResults[1].SetResult(dice.FOCUS)
	attackResults[2].SetResult(dice.HIT)
	state.SetAttackResults(&attackResults)

	defender.SetFocusTokens(3)
	defenseResults := dice.RollDefenseDice(3)
	defenseResults[0].SetResult(dice.BLANK)
	defenseResults[1].SetResult(dice.FOCUS)
	defenseResults[2].SetResult(dice.EVADE)
	state.SetDefenseResults(&defenseResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, defender)

	assert.EqualValues(2, attacker.FocusTokens())
	assert.EqualValues(1, attackResults.Blanks())
	assert.EqualValues(1, attackResults.Focuses())
	assert.EqualValues(1, attackResults.Hits())
	assert.EqualValues(0, attackResults.Crits())

	assert.EqualValues(2, defender.FocusTokens())
	assert.EqualValues(1, defenseResults.Blanks())
	assert.EqualValues(0, defenseResults.Focuses())
	assert.EqualValues(2, defenseResults.Evades())
}

func TestSpendFocus_OnDefenseHasNoTokens(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := SpendFocus{}

	attacker.SetFocusTokens(2)
	attackResults := dice.RollAttackDice(3)
	attackResults[0].SetResult(dice.BLANK)
	attackResults[1].SetResult(dice.FOCUS)
	attackResults[2].SetResult(dice.HIT)
	state.SetAttackResults(&attackResults)

	defender.SetFocusTokens(0)
	defenseResults := dice.RollDefenseDice(3)
	defenseResults[0].SetResult(dice.BLANK)
	defenseResults[1].SetResult(dice.FOCUS)
	defenseResults[2].SetResult(dice.EVADE)
	state.SetDefenseResults(&defenseResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, defender)

	assert.EqualValues(2, attacker.FocusTokens())
	assert.EqualValues(1, attackResults.Blanks())
	assert.EqualValues(1, attackResults.Focuses())
	assert.EqualValues(1, attackResults.Hits())
	assert.EqualValues(0, attackResults.Crits())

	assert.EqualValues(0, defender.FocusTokens())
	assert.EqualValues(1, defenseResults.Blanks())
	assert.EqualValues(1, defenseResults.Focuses())
	assert.EqualValues(1, defenseResults.Evades())
}

func TestSpendFocus_OnDefenseHasNoFocusResults(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := SpendFocus{}

	attacker.SetFocusTokens(2)
	attackResults := dice.RollAttackDice(3)
	attackResults[0].SetResult(dice.BLANK)
	attackResults[1].SetResult(dice.FOCUS)
	attackResults[2].SetResult(dice.HIT)
	state.SetAttackResults(&attackResults)

	defender.SetFocusTokens(3)
	defenseResults := dice.RollDefenseDice(3)
	defenseResults[0].SetResult(dice.BLANK)
	defenseResults[1].SetResult(dice.BLANK)
	defenseResults[2].SetResult(dice.EVADE)
	state.SetDefenseResults(&defenseResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, defender)

	assert.EqualValues(2, attacker.FocusTokens())
	assert.EqualValues(1, attackResults.Blanks())
	assert.EqualValues(1, attackResults.Focuses())
	assert.EqualValues(1, attackResults.Hits())
	assert.EqualValues(0, attackResults.Crits())

	assert.EqualValues(3, defender.FocusTokens())
	assert.EqualValues(2, defenseResults.Blanks())
	assert.EqualValues(0, defenseResults.Focuses())
	assert.EqualValues(1, defenseResults.Evades())
}
