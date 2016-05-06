package runner

import (
	// "fmt"
	"github.com/geordanr/go_xwing/attack"
	"github.com/geordanr/go_xwing/attack/modification"
	"github.com/geordanr/go_xwing/attack/step"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/gamestate"
	"github.com/geordanr/go_xwing/interfaces"
	"github.com/geordanr/go_xwing/ship"
	// "github.com/stretchr/testify/assert"
	"testing"
)

func TestRun_MultipleStates(t *testing.T) {
	// assert := assert.New(t)

	makeState := func() *gamestate.GameState {
		attacker := ship.New("Attacker", 0, 3, 2, 3, 2)
		defender := ship.New("Defender", 0, 2, 3, 3, 0)
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
	bufsz := 8
	runner := New(step.All, bufsz)
	output := make(chan interfaces.GameState, bufsz)
	go runner.Run(output)

	for i := 0; i < nStates; i++ {
		runner.InjectState(makeState())
	}

	for i := 0; i < nStates; i++ {
		<-output
	}
}
