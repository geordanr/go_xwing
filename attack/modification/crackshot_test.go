package modification

import (
	"github.com/geordanr/go_xwing/attack"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/gamestate"
	"github.com/geordanr/go_xwing/ship"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrackShot(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := CrackShot{}

	attackResults := dice.RollAttackDice(3)
	attackResults[0].SetResult(dice.BLANK)
	attackResults[1].SetResult(dice.FOCUS)
	attackResults[2].SetResult(dice.HIT)
	state.SetAttackResults(&attackResults)

	defenseResults := dice.RollDefenseDice(3)
	defenseResults[0].SetResult(dice.BLANK)
	defenseResults[1].SetResult(dice.FOCUS)
	defenseResults[2].SetResult(dice.EVADE)
	state.SetDefenseResults(&defenseResults)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, attacker)

	defenseResults = *state.DefenseResults()
	assert.EqualValues(1, defenseResults.Blanks())
	assert.EqualValues(1, defenseResults.Focuses())
	assert.EqualValues(0, defenseResults.Evades())
}
