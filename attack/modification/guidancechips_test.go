package modification

import (
	"github.com/geordanr/go_xwing/attack"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/gamestate"
	"github.com/geordanr/go_xwing/ship"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGuidanceChips_attack2_convertBlank(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 2, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := GuidanceChips{}

	attackResults := dice.RollAttackDice(3)
	attackResults[0].SetResult(dice.BLANK)
	attackResults[1].SetResult(dice.FOCUS)
	attackResults[2].SetResult(dice.HIT)
	state.SetAttackResults(&attackResults)

	defenseResults := dice.RollDefenseDice(3)
	state.SetDefenseResults(&defenseResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, attacker)

	attackResults = *state.AttackResults()
	assert.EqualValues(0, attackResults.Blanks())
	assert.EqualValues(1, attackResults.Focuses())
	assert.EqualValues(2, attackResults.Hits())
	assert.EqualValues(0, attackResults.Crits())
}

func TestGuidanceChips_attack2_convertFocus(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 2, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := GuidanceChips{}

	attackResults := dice.RollAttackDice(3)
	attackResults[0].SetResult(dice.FOCUS)
	attackResults[1].SetResult(dice.FOCUS)
	attackResults[2].SetResult(dice.HIT)
	state.SetAttackResults(&attackResults)

	defenseResults := dice.RollDefenseDice(3)
	state.SetDefenseResults(&defenseResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, attacker)

	attackResults = *state.AttackResults()
	assert.EqualValues(0, attackResults.Blanks())
	assert.EqualValues(1, attackResults.Focuses())
	assert.EqualValues(2, attackResults.Hits())
	assert.EqualValues(0, attackResults.Crits())
}

func TestGuidanceChips_attack3_convertBlank(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 3, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := GuidanceChips{}

	attackResults := dice.RollAttackDice(3)
	attackResults[0].SetResult(dice.BLANK)
	attackResults[1].SetResult(dice.FOCUS)
	attackResults[2].SetResult(dice.HIT)
	state.SetAttackResults(&attackResults)

	defenseResults := dice.RollDefenseDice(3)
	state.SetDefenseResults(&defenseResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, attacker)

	attackResults = *state.AttackResults()
	assert.EqualValues(0, attackResults.Blanks())
	assert.EqualValues(1, attackResults.Focuses())
	assert.EqualValues(1, attackResults.Hits())
	assert.EqualValues(1, attackResults.Crits())
}

func TestGuidanceChips_attack3_convertFocus(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 3, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := GuidanceChips{}

	attackResults := dice.RollAttackDice(3)
	attackResults[0].SetResult(dice.FOCUS)
	attackResults[1].SetResult(dice.FOCUS)
	attackResults[2].SetResult(dice.HIT)
	state.SetAttackResults(&attackResults)

	defenseResults := dice.RollDefenseDice(3)
	state.SetDefenseResults(&defenseResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, attacker)

	attackResults = *state.AttackResults()
	assert.EqualValues(0, attackResults.Blanks())
	assert.EqualValues(1, attackResults.Focuses())
	assert.EqualValues(1, attackResults.Hits())
	assert.EqualValues(1, attackResults.Crits())
}

func TestGuidanceChips_attack3_convertHit(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 3, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := GuidanceChips{}

	attackResults := dice.RollAttackDice(3)
	attackResults[0].SetResult(dice.HIT)
	attackResults[1].SetResult(dice.HIT)
	attackResults[2].SetResult(dice.HIT)
	state.SetAttackResults(&attackResults)

	defenseResults := dice.RollDefenseDice(3)
	state.SetDefenseResults(&defenseResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, attacker)

	attackResults = *state.AttackResults()
	assert.EqualValues(0, attackResults.Blanks())
	assert.EqualValues(0, attackResults.Focuses())
	assert.EqualValues(2, attackResults.Hits())
	assert.EqualValues(1, attackResults.Crits())
}
