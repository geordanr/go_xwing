package ship

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpendFocus(t *testing.T) {
	ship := Ship{}

	ret := ship.SpendFocus()
	if ret {
		t.Errorf("Should return false when spending focus on ship with no focus tokens")
	}

	ship.focusTokens = 2
	ret = ship.SpendFocus()
	if !ret {
		t.Errorf("Should have succeded spending focus")
	}
	if ship.focusTokens != 1 {
		t.Errorf("Ship has %d focus tokens left when it should have 1", ship.focusTokens)
	}
}

func TestSpendEvade(t *testing.T) {
	ship := Ship{}

	ret := ship.SpendEvade()
	if ret {
		t.Errorf("Should return false when spending Evade on ship with no Evade tokens")
	}

	ship.evadeTokens = 2
	ret = ship.SpendEvade()
	if !ret {
		t.Errorf("Should have succeded spending Evade")
	}
	if ship.evadeTokens != 1 {
		t.Errorf("Ship has %d Evade tokens left when it should have 1", ship.evadeTokens)
	}
}

func TestIsAlive(t *testing.T) {
	ship := Ship{}

	if ship.IsAlive() {
		t.Errorf("Ship shouldn't be alive, has %d hull", ship.hull)
	}

	ship.hull = 1

	if !ship.IsAlive() {
		t.Errorf("Ship should be alive, has %d hull", ship.hull)
	}
}

func TestSufferDamage(t *testing.T) {
	ship := Ship{shields: 4, hull: 4}

	ship.SufferDamage(2, 1)
	if ship.shields != 1 {
		t.Errorf("Expected 1 shield left, has %d", ship.shields)
	}
	if ship.hull != 4 {
		t.Errorf("Expected 4 hull left, has %d", ship.hull)
	}

	ship.SufferDamage(1, 1)
	if ship.shields != 0 {
		t.Errorf("Expected 0 shields left, has %d", ship.shields)
	}
	if ship.hull != 3 {
		t.Errorf("Expected 3 hull left, has %d", ship.hull)
	}

	ship = Ship{shields: 4, hull: 4}
	ship.SufferDamage(4, 4)
	if ship.IsAlive() {
		t.Errorf("Ship should be dead, has %d hull", ship.hull)
	}

	ship = Ship{shields: 4, hull: 4}
	ship.SufferDamage(5, 5)
	if ship.IsAlive() {
		t.Errorf("Ship should be dead, has %d hull", ship.hull)
	}

}

func TestCopy(t *testing.T) {
	assert := assert.New(t)

	wedge := Ship{
		name:         "Wedge Antilles",
		skill:        9,
		attack:       3,
		agility:      2,
		hull:         3,
		shields:      2,
		focusTokens:  1,
		evadeTokens:  0,
		targetLock:   "Colonel Vessery",
		cannotAttack: false,
	}

	copyInterface := wedge.Copy()
	cp, ok := copyInterface.(*Ship)
	assert.True(ok)

	assert.Equal(wedge.name, cp.name)
	assert.Equal(wedge.skill, cp.skill)
	assert.Equal(wedge.attack, cp.attack)
	assert.Equal(wedge.agility, cp.agility)
	assert.Equal(wedge.hull, cp.hull)
	assert.Equal(wedge.shields, cp.shields)
	assert.Equal(wedge.focusTokens, cp.focusTokens)
	assert.Equal(wedge.evadeTokens, cp.evadeTokens)
	assert.Equal(wedge.targetLock, cp.targetLock)
	assert.Equal(wedge.cannotAttack, cp.cannotAttack)

	cp.name = "Miranda Doni"
	cp.skill = 8
	cp.attack = 2
	cp.agility = 1
	cp.hull = 5
	cp.shields = 4
	cp.focusTokens = 0
	cp.evadeTokens = 1
	cp.targetLock = "Howlrunner"
	cp.cannotAttack = true

	assert.NotEqual(wedge.name, cp.name)
	assert.NotEqual(wedge.skill, cp.skill)
	assert.NotEqual(wedge.attack, cp.attack)
	assert.NotEqual(wedge.agility, cp.agility)
	assert.NotEqual(wedge.hull, cp.hull)
	assert.NotEqual(wedge.shields, cp.shields)
	assert.NotEqual(wedge.focusTokens, cp.focusTokens)
	assert.NotEqual(wedge.evadeTokens, cp.evadeTokens)
	assert.NotEqual(wedge.targetLock, cp.targetLock)
	assert.NotEqual(wedge.cannotAttack, cp.cannotAttack)

	wedge.name = "Colonel Vessery"
	wedge.skill = 6
	wedge.attack = 3
	wedge.agility = 3
	wedge.hull = 3
	wedge.shields = 3
	wedge.focusTokens = 2
	wedge.evadeTokens = 2
	wedge.targetLock = ""
	wedge.cannotAttack = false

	assert.NotEqual(wedge.name, cp.name)
	assert.NotEqual(wedge.skill, cp.skill)
	assert.NotEqual(wedge.attack, cp.attack)
	assert.NotEqual(wedge.agility, cp.agility)
	assert.NotEqual(wedge.hull, cp.hull)
	assert.NotEqual(wedge.shields, cp.shields)
	assert.NotEqual(wedge.focusTokens, cp.focusTokens)
	assert.NotEqual(wedge.evadeTokens, cp.evadeTokens)
	assert.NotEqual(wedge.targetLock, cp.targetLock)
	assert.NotEqual(wedge.cannotAttack, cp.cannotAttack)
}
