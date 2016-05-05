package ship

import "fmt"

type Ship struct {
	name        string
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
func New(name string, attack, agility, hull, shields uint) *Ship {
	var ship Ship

	ship = Ship{
		name:    name,
		attack:  attack,
		agility: agility,
		hull:    hull,
		shields: shields,
	}

	return &ship
}

func (ship *Ship) String() string {
	return fmt.Sprintf("<Ship name='%s' %d/%d/%d/%d focus=%d evade=%d>", ship.name, ship.attack, ship.agility, ship.hull, ship.shields, ship.focusTokens, ship.evadeTokens)
}

func (ship Ship) Name() string  { return ship.name }
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

func (ship *Ship) SufferDamage(nDamage uint) {
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

// func (s *Ship) PerformActions() {
//     for _, action := range(s.Actions) {
// 	action.Perform(s)
//     }
// }

/////////////////////////////////////////////////

// var ShipFactory map[string]func() Ship = map[string]func() Ship{
// 	"E-Wing": func () Ship { return Ship{ Name: " E-Wing", Attack: 3, Agility: 3, Hull: 2, Shields: 3, }},
// 	"M3-A Syck Interceptor": func () Ship { return Ship{ Name: " M3-A Syck Interceptor", Attack: 2, Agility: 3, Hull: 3, Shields: 0, }},
// 	"Star Viper": func () Ship { return Ship{ Name: " Star Viper", Attack: 3, Agility: 3, Hull: 4, Shields: 1, }},
// 	"X-Wing": func () Ship { return Ship{ Name: " X-Wing", Attack: 3, Agility: 2, Hull: 3, Shields: 2, }},
// 	"VT-49 Decimator": func () Ship { return Ship{ Name: " VT-49 Decimator", Attack: 3, Agility: 0, Hull: 12, Shields: 4, }},
// 	"TIE Advanced": func () Ship { return Ship{ Name: " TIE Advanced", Attack: 2, Agility: 3, Hull: 3, Shields: 2, }},
// 	"HWK-290": func () Ship { return Ship{ Name: " HWK-290", Attack: 1, Agility: 2, Hull: 4, Shields: 1, }},
// 	"B-Wing": func () Ship { return Ship{ Name: " B-Wing", Attack: 3, Agility: 1, Hull: 3, Shields: 5, }},
// 	"Lambda-Class Shuttle": func () Ship { return Ship{ Name: " Lambda-Class Shuttle", Attack: 3, Agility: 1, Hull: 5, Shields: 5, }},
// 	"Aggressor": func () Ship { return Ship{ Name: " Aggressor", Attack: 3, Agility: 3, Hull: 4, Shields: 4, }},
// 	"Firespray-31": func () Ship { return Ship{ Name: " Firespray-31", Attack: 3, Agility: 2, Hull: 6, Shields: 4, }},
// 	"TIE Phantom": func () Ship { return Ship{ Name: " TIE Phantom", Attack: 4, Agility: 2, Hull: 2, Shields: 2, }},
// 	"Y-Wing": func () Ship { return Ship{ Name: " Y-Wing", Attack: 2, Agility: 1, Hull: 5, Shields: 3, }},
// 	"YT-1300": func () Ship { return Ship{ Name: " YT-1300", Attack: 3, Agility: 1, Hull: 8, Shields: 5, }},
// 	"TIE Interceptor": func () Ship { return Ship{ Name: " TIE Interceptor", Attack: 3, Agility: 3, Hull: 3, Shields: 0, }},
// 	"TIE Bomber": func () Ship { return Ship{ Name: " TIE Bomber", Attack: 2, Agility: 3, Hull: 6, Shields: 0, }},
// 	"TIE Fighter": func () Ship { return Ship{ Name: " TIE Fighter", Attack: 2, Agility: 3, Hull: 3, Shields: 0, }},
// 	"YT-2400 Freighter": func () Ship { return Ship{ Name: " YT-2400 Freighter", Attack: 2, Agility: 2, Hull: 5, Shields: 5, }},
// 	"Z-95 Headhunter": func () Ship { return Ship{ Name: " Z-95 Headhunter", Attack: 2, Agility: 2, Hull: 2, Shields: 2, }},
// 	"A-Wing": func () Ship { return Ship{ Name: " A-Wing", Attack: 2, Agility: 3, Hull: 2, Shields: 2, }},
// 	"TIE Defender": func () Ship { return Ship{ Name: " TIE Defender", Attack: 3, Agility: 3, Hull: 3, Shields: 3, }},
// }
