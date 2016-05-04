package runner

import (
	// "fmt"
	"github.com/geordanr/go_xwing/attack"
	"github.com/geordanr/go_xwing/attack/modification"
	"github.com/geordanr/go_xwing/attack/step"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/gamestate"
	"github.com/geordanr/go_xwing/interfaces"
	"github.com/geordanr/go_xwing/ship"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0)
	state := gamestate.GameState{}
	attackerSpendFocus := modification.SpendFocus{}
	attackerSpendFocus.SetActor(constants.ATTACKER)
	mods := map[string][]interfaces.Modification{
		"Modify Attack Dice": []interfaces.Modification{
			&attackerSpendFocus,
		},
	}
	runner := New(step.Steps)

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

	state.EnqueueAttack(attack.New(attacker, defender, mods))

	runner.Run(&state)

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
