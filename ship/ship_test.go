package ship

import "testing"

func TestSpendFocus(t *testing.T) {
    ship := Ship{}

    ret := ship.SpendFocus()
    if ret {
	t.Errorf("Should return false when spending focus on ship with no focus tokens")
    }

    ship.nFocusTokens = 2
    ret = ship.SpendFocus()
    if !ret {
	t.Errorf("Should have succeded spending focus")
    }
    if ship.nFocusTokens != 1 {
	t.Errorf("Ship has %d focus tokens left when it should have 1", ship.nFocusTokens)
    }
}

func TestSpendEvade(t *testing.T) {
    ship := Ship{}

    ret := ship.SpendEvade()
    if ret {
	t.Errorf("Should return false when spending Evade on ship with no Evade tokens")
    }

    ship.nEvadeTokens = 2
    ret = ship.SpendEvade()
    if !ret {
	t.Errorf("Should have succeded spending Evade")
    }
    if ship.nEvadeTokens != 1 {
	t.Errorf("Ship has %d Evade tokens left when it should have 1", ship.nEvadeTokens)
    }
}

func TestIsAlive(t *testing.T) {
    ship := Ship{}

    if ship.IsAlive() {
	t.Errorf("Ship shouldn't be alive, has %d hull", ship.nHull)
    }

    ship.nHull = 1

    if !ship.IsAlive() {
	t.Errorf("Ship should be alive, has %d hull", ship.nHull)
    }
}

func TestSufferDamage(t *testing.T) {
    ship := Ship{nShields: 4, nHull: 4}

    ship.SufferDamage(3)
    if ship.nShields != 1 {
	t.Errorf("Expected 1 shield left, has %d", ship.nShields)
    }
    if ship.nHull != 4 {
	t.Errorf("Expected 4 hull left, has %d", ship.nHull)
    }

    ship.SufferDamage(2)
    if ship.nShields != 0 {
	t.Errorf("Expected 0 shields left, has %d", ship.nShields)
    }
    if ship.nHull != 3 {
	t.Errorf("Expected 3 hull left, has %d", ship.nHull)
    }

    ship = Ship{nShields: 4, nHull: 4}
    ship.SufferDamage(8)
    if ship.IsAlive() {
	t.Errorf("Ship should be dead, has %d hull", ship.nHull)
    }

    ship = Ship{nShields: 4, nHull: 4}
    ship.SufferDamage(10)
    if ship.IsAlive() {
	t.Errorf("Ship should be dead, has %d hull", ship.nHull)
    }

}
