package modification

import (
	"github.com/geordanr/go_xwing/attack"
	"github.com/geordanr/go_xwing/gamestate"
	"github.com/geordanr/go_xwing/ship"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDealDamage_NoHits(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 2, 3, 3, 0)
	defender := ship.New("Defender", 3, 2, 3, 2)
	state := gamestate.GameState{}
	mod := DealDamage{}

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, nil)

	assert.EqualValues(3, defender.Hull())
	assert.EqualValues(2, defender.Shields())
}

func TestDealDamage_ShieldBeforeHull(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 2, 3, 3, 0)
	defender := ship.New("Defender", 3, 2, 3, 2)
	state := gamestate.GameState{}
	mod := DealDamage{}

	state.SetHitsLanded(2)
	state.SetCritsLanded(1)

	state.EnqueueAttack(attack.New(attacker, defender, nil))
	mod.ModifyState(&state, nil)

	assert.EqualValues(2, defender.Hull())
	assert.EqualValues(0, defender.Shields())
}
