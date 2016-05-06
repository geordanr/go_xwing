package modification

import (
	"github.com/geordanr/go_xwing/attack"
	"github.com/geordanr/go_xwing/gamestate"
	"github.com/geordanr/go_xwing/ship"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeclareTarget_Normal(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 2, 3, 3, 0)
	defender := ship.New("Defender", 0, 3, 2, 3, 2)
	state := gamestate.GameState{}
	mod := DeclareTarget{}

	state.SetNextAttackStep("next")

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, nil)

	assert.NotEqual("", state.NextAttackStep())
}

func TestDeclareTarget_CannotAttack(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 2, 3, 3, 0)
	defender := ship.New("Defender", 0, 3, 2, 3, 2)
	state := gamestate.GameState{}
	mod := DeclareTarget{}

	state.SetNextAttackStep("next")
	attacker.SetCanAttack(false)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, nil)

	assert.Equal("", state.NextAttackStep())
}

func TestDeclareTarget_DefenderDead(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 2, 3, 3, 0)
	defender := ship.New("Defender", 0, 3, 2, 0, 0)
	state := gamestate.GameState{}
	mod := DeclareTarget{}

	state.SetNextAttackStep("next")

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, nil)

	assert.Equal("", state.NextAttackStep())
}
