package attack

import (
	"github.com/stretchr/testify/assert"
	"testing"
	// "github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/interfaces"
	"github.com/geordanr/go_xwing/ship"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0)
	modifications := map[string][]interfaces.Modification{}

	atk := New(attacker, defender, modifications)

	assert.Equal(atk.attacker, attacker)
	assert.Equal(atk.defender, defender)
	assert.Equal(atk.modifications, modifications)
}

func TestAttacker(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0)
	modifications := map[string][]interfaces.Modification{}

	atk := New(attacker, defender, modifications)

	assert.Equal(atk.Attacker(), attacker)
}

func TestDefender(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0)
	modifications := map[string][]interfaces.Modification{}

	atk := New(attacker, defender, modifications)

	assert.Equal(atk.Defender(), defender)
}

func TestModifications(t *testing.T) {
	assert := assert.New(t)

	attacker := ship.New("Attacker", 0, 0, 0, 0)
	defender := ship.New("Defender", 0, 0, 0, 0)
	modifications := map[string][]interfaces.Modification{}

	atk := New(attacker, defender, modifications)

	assert.Equal(atk.Modifications(), modifications)
}
