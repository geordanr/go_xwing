// Package attack represents the parameters for a single attack.
package attack

import (
	"fmt"
	"github.com/geordanr/go_xwing/interfaces"
)

type Attack struct {
	attacker      interfaces.Ship
	defender      interfaces.Ship
	modifications map[string][]interfaces.Modification
}

// New returns a new Attack with the given attacker, defender, and map of step names to list of modifications.
func New(attacker, defender interfaces.Ship, modifications map[string][]interfaces.Modification) *Attack {
	atk := Attack{
		attacker:      attacker,
		defender:      defender,
		modifications: modifications,
	}
	return &atk
}

// Attacker returns the attacking ship in this attack.
func (atk *Attack) Attacker() interfaces.Ship {
	return atk.attacker
}

// Defender returns the defending ship in this attack.
func (atk *Attack) Defender() interfaces.Ship {
	return atk.defender
}

// Modifications returns the map of step names to modifications.
func (atk *Attack) Modifications() map[string][]interfaces.Modification {
	return atk.modifications
}

// Copy returns a copy of this Attack.
func (atk *Attack) Copy() interfaces.Attack {
	cp := Attack{
		attacker:      atk.attacker,
		defender:      atk.defender,
		modifications: map[string][]interfaces.Modification{},
	}
	// Since there's no map copy...
	for stepName, mods := range atk.modifications {
		newMods := make([]interfaces.Modification, len(mods))
		copy(newMods, mods)
		cp.modifications[stepName] = newMods
	}
	return &cp
}

func (atk Attack) String() string {
	return fmt.Sprintf("<Attack attacker=%s defender=%s>", atk.attacker, atk.defender)
}
