package modification

import (
	"github.com/geordanr/go_xwing/attack"
	"github.com/geordanr/go_xwing/gamestate"
	"github.com/geordanr/go_xwing/ship"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCannotAttackAgain(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 2, 3, 3, 0)
	defender := ship.New("Defender", 0, 3, 2, 3, 2)
	state := gamestate.GameState{}
	mod := CannotAttackAgain{}

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, nil)

	assert.False(attacker.CanAttack())
}
