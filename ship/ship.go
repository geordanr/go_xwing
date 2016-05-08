package ship

import (
	"fmt"
	"github.com/geordanr/go_xwing/interfaces"
)

type Ship struct {
	name        string
	skill       uint
	attack      uint
	agility     uint
	hull        uint
	shields     uint
	focusTokens uint
	evadeTokens uint
	targetLock  string
	// Actions []Action
	cannotAttack bool // default zero value means we can attack
}

// New returns a pointer to a ship with the given stats.
func New(name string, skill, attack, agility, hull, shields uint) *Ship {
	var ship Ship

	ship = Ship{
		name:    name,
		skill:   skill,
		attack:  attack,
		agility: agility,
		hull:    hull,
		shields: shields,
	}

	return &ship
}

func (ship *Ship) String() string {
	return fmt.Sprintf("<Ship name='%s' PS%d %d/%d/%d/%d focus=%d evade=%d>", ship.name, ship.skill, ship.attack, ship.agility, ship.hull, ship.shields, ship.focusTokens, ship.evadeTokens)
}

func (ship Ship) Name() string  { return ship.name }
func (ship Ship) Skill() uint   { return ship.skill }
func (ship Ship) Attack() uint  { return ship.attack }
func (ship Ship) Agility() uint { return ship.agility }
func (ship Ship) Hull() uint    { return ship.hull }
func (ship Ship) Shields() uint { return ship.shields }

func (ship Ship) FocusTokens() uint      { return ship.focusTokens }
func (ship *Ship) SetFocusTokens(n uint) { ship.focusTokens = n }
func (ship Ship) EvadeTokens() uint      { return ship.evadeTokens }
func (ship *Ship) SetEvadeTokens(n uint) { ship.evadeTokens = n }

func (ship Ship) TargetLock() string         { return ship.targetLock }
func (ship *Ship) SetTargetLock(name string) { ship.targetLock = name }

func (ship *Ship) SpendFocus() bool {
	if ship.focusTokens > 0 {
		ship.focusTokens--
		return true
	} else {
		return false
	}
}

func (ship *Ship) SpendEvade() bool {
	if ship.evadeTokens > 0 {
		ship.evadeTokens--
		return true
	} else {
		return false
	}
}

func (ship Ship) IsAlive() bool {
	return ship.hull > 0
}

func (ship Ship) CanAttack() bool {
	return !ship.cannotAttack
}

func (ship *Ship) SetCanAttack(b bool) {
	ship.cannotAttack = !b
}

func (ship *Ship) SufferDamage(hits uint, crits uint) {
	ship.applyDamage(hits)
	// TODO handle crits differently?
	ship.applyDamage(crits)
	// fmt.Printf("%s (%p) damaged\n", ship.name, ship)
}

func (ship *Ship) applyDamage(nDamage uint) {
	var i uint
	for i = 0; i < nDamage; i++ {
		if ship.shields > 0 {
			ship.shields--
		} else if ship.hull > 0 {
			ship.hull--
		} else {
			break
		}
	}
}

func (ship Ship) Copy() interfaces.Ship {
	newShip := &Ship{}
	*newShip = ship
	return newShip
}
