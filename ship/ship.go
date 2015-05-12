package ship

type Ship struct {
    Shields uint
    Hull uint
    FocusTokens uint
    EvadeTokens uint
    canAttack bool
}

func (ship *Ship) SpendFocus() bool {
    if ship.FocusTokens > 0 {
	ship.FocusTokens--
	return true
    } else {
	return false
    }
}

func (ship *Ship) SpendEvade() bool {
    if ship.EvadeTokens > 0 {
	ship.EvadeTokens--
	return true
    } else {
	return false
    }
}

func (ship *Ship) IsAlive() bool {
    return ship.Hull > 0
}

func (ship *Ship) SufferDamage(nDamage uint) {
    var i uint
    for i = 0; i < nDamage; i++ {
	if ship.Shields > 0 {
	    ship.Shields--
	} else if ship.Hull > 0 {
	    ship.Hull--
	} else {
	    break
	}
    }
}
