package modification

import (
	"github.com/geordanr/go_xwing/attack"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/gamestate"
	"github.com/geordanr/go_xwing/ship"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccuracyCorrector(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0, 0)
	state := gamestate.GameState{}
	mod := AccuracyCorrector{}

	for i := 0; i < 100; i++ {
		attackResults := dice.RollAttackDice(10)
		state.SetAttackResults(&attackResults)

		state.EnqueueAttack(attack.New(attacker, defender, nil))
		mod.ModifyState(&state, attacker)

		attackResults = *state.AttackResults()
		if attackResults.Hits()+attackResults.Crits() < 2 {
			assert.EqualValues(0, attackResults.Blanks())
			assert.EqualValues(0, attackResults.Focuses())
			assert.EqualValues(2, attackResults.Hits())
			assert.EqualValues(0, attackResults.Crits())
		}
	}
}
