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

func TestRun_ContrivedFocus(t *testing.T) {
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
	runner := New(step.All, 1)
	output := make(chan interfaces.GameState)
	go runner.Run(output)

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

	runner.InjectState(&state)
	<-output

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

func TestRun_MultipleStates(t *testing.T) {
	// assert := assert.New(t)

	makeState := func() *gamestate.GameState {
		attacker := ship.New("Attacker", 3, 2, 3, 2)
		defender := ship.New("Defender", 2, 3, 3, 0)
		state := gamestate.GameState{}
		attackerSpendFocus := modification.SpendFocus{}
		attackerSpendFocus.SetActor(constants.ATTACKER)
		attackerRollDice := modification.RollDice{}
		attackerRollDice.SetActor(constants.ATTACKER)
		defenderRollDice := modification.RollDice{}
		defenderRollDice.SetActor(constants.DEFENDER)
		compareResults := modification.CompareResults{}
		mods := map[string][]interfaces.Modification{
			"Modify Attack Dice": []interfaces.Modification{
				&attackerSpendFocus,
			},
			"Roll Attack Dice": []interfaces.Modification{
				&attackerRollDice,
			},
			"Roll Defense Dice": []interfaces.Modification{
				&defenderRollDice,
			},
			"Compare Results": []interfaces.Modification{
				&compareResults,
			},
		}

		attacker.SetFocusTokens(2)
		defender.SetFocusTokens(3)
		state.EnqueueAttack(attack.New(attacker, defender, mods))
		state.EnqueueAttack(attack.New(defender, attacker, mods))

		return &state
	}

	nStates := 100
	runner := New(step.All, nStates)
	output := make(chan interfaces.GameState, nStates)
	go runner.Run(output)

	for i := 0; i < nStates; i++ {
		runner.InjectState(makeState())
	}

	for i := 0; i < nStates; i++ {
		<-output
	}
}
