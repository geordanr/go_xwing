package modification

import (
	"github.com/geordanr/go_xwing/attack"
	"github.com/geordanr/go_xwing/gamestate"
	"github.com/geordanr/go_xwing/interfaces"
	"github.com/geordanr/go_xwing/ship"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGunner_Missed(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 2, 3, 3, 0)
	defender := ship.New("Defender", 0, 3, 2, 3, 2)
	state := gamestate.GameState{}
	mod := Gunner{}

	state.SetAttackMissed(true)

	state.EnqueueAttack(attack.New(attacker, defender, map[string][]interfaces.Modification{}))
	mod.ModifyState(&state, nil)

	assert.True(state.DequeueAttack())
	atk := state.CurrentAttack()
	stepMods, exists := atk.Modifications()["After Attacking/Defending"]
	assert.True(exists)

	found := false
	for _, mod := range stepMods {
		switch mod.(type) {
		case *CannotAttackAgain:
			found = true
		}
	}
	assert.True(found)
}

func TestGunner_Hit(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 2, 3, 3, 0)
	defender := ship.New("Defender", 0, 3, 2, 3, 2)
	state := gamestate.GameState{}
	mod := Gunner{}

	state.EnqueueAttack(attack.New(attacker, defender, map[string][]interfaces.Modification{}))
	mod.ModifyState(&state, nil)

	assert.False(state.DequeueAttack())
}
