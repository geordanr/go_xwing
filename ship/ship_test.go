package ship

import "testing"

func TestSpendFocus(t *testing.T) {
    ship := Ship{}

    ret := ship.SpendFocus()
    if ret {
	t.Errorf("Should return false when spending focus on ship with no focus tokens")
    }

    ship.FocusTokens = 2
    ret = ship.SpendFocus()
    if !ret {
	t.Errorf("Should have succeded spending focus")
    }
    if ship.FocusTokens != 1 {
	t.Errorf("Ship has %d focus tokens left when it should have 1", ship.FocusTokens)
    }
}

func TestSpendEvade(t *testing.T) {
    ship := Ship{}

    ret := ship.SpendEvade()
    if ret {
	t.Errorf("Should return false when spending Evade on ship with no Evade tokens")
    }

    ship.EvadeTokens = 2
    ret = ship.SpendEvade()
    if !ret {
	t.Errorf("Should have succeded spending Evade")
    }
    if ship.EvadeTokens != 1 {
	t.Errorf("Ship has %d Evade tokens left when it should have 1", ship.EvadeTokens)
    }
}

func TestIsAlive(t *testing.T) {
    ship := Ship{}

    if ship.IsAlive() {
	t.Errorf("Ship shouldn't be alive, has %d hull", ship.Hull)
    }

    ship.Hull = 1

    if !ship.IsAlive() {
	t.Errorf("Ship should be alive, has %d hull", ship.Hull)
    }
}

func TestSufferDamage(t *testing.T) {
    ship := Ship{Shields: 4, Hull: 4}

    ship.SufferDamage(3)
    if ship.Shields != 1 {
	t.Errorf("Expected 1 shield left, has %d", ship.Shields)
    }
    if ship.Hull != 4 {
	t.Errorf("Expected 4 hull left, has %d", ship.Hull)
    }

    ship.SufferDamage(2)
    if ship.Shields != 0 {
	t.Errorf("Expected 0 shields left, has %d", ship.Shields)
    }
    if ship.Hull != 3 {
	t.Errorf("Expected 3 hull left, has %d", ship.Hull)
    }

    ship = Ship{Shields: 4, Hull: 4}
    ship.SufferDamage(8)
    if ship.IsAlive() {
	t.Errorf("Ship should be dead, has %d hull", ship.Hull)
    }

    ship = Ship{Shields: 4, Hull: 4}
    ship.SufferDamage(10)
    if ship.IsAlive() {
	t.Errorf("Ship should be dead, has %d hull", ship.Hull)
    }

}
