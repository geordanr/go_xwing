package ship

import "fmt"

type Ship struct {
    Name string
    Attack uint
    Agility uint
    Hull uint
    Shields uint
    FocusTokens uint
    EvadeTokens uint
    canAttack bool
}

func (ship *Ship) String() string {
    return fmt.Sprintf("<Ship name='%s' %d/%d/%d/%d focus=%d evade=%s>", ship.Name, ship.Attack, ship.Agility, ship.Hull, ship.Shields, ship.FocusTokens, ship.EvadeTokens)
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

/////////////////////////////////////////////////

var ShipFactory map[string]func() Ship = map[string]func() Ship{
    "B-Wing": func () Ship {
	return Ship{
	    Name: "B-Wing",
	    Attack: 3,
	    Agility: 1,
	    Hull: 3,
	    Shields: 5,
	    canAttack: true,
	}
    },
    "TIE Interceptor": func () Ship {
	return Ship{
	    Name: "TIE Interceptor",
	    Attack: 3,
	    Agility: 3,
	    Hull: 3,
	    Shields: 0,
	    canAttack: true,
	}
    },
}
