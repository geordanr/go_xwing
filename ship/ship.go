package ship

type Ship struct {
    nShields uint
    nHull uint
    nFocusTokens uint
    nEvadeTokens uint
    canAttack bool
}

func (ship *Ship) SpendFocus() bool {
    if ship.nFocusTokens > 0 {
	ship.nFocusTokens--
	return true
    } else {
	return false
    }
}

func (ship *Ship) SpendEvade() bool {
    if ship.nEvadeTokens > 0 {
	ship.nEvadeTokens--
	return true
    } else {
	return false
    }
}

func (ship *Ship) IsAlive() bool {
    return ship.nHull > 0
}

func (ship *Ship) SufferDamage(nDamage uint) {
    var i uint
    for i = 0; i < nDamage; i++ {
	if ship.nShields > 0 {
	    ship.nShields--
	} else if ship.nHull > 0 {
	    ship.nHull--
	} else {
	    break
	}
    }
}
