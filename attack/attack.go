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
func (atk *Attack) SetModifications(mods map[string][]interfaces.Modification) {
	atk.modifications = mods
}

// Copy returns a copy of this Attack, skipping transient mods.
func (atk *Attack) Copy() interfaces.Attack {
	cp := Attack{
		attacker:      atk.attacker.Copy(),
		defender:      atk.defender.Copy(),
		modifications: map[string][]interfaces.Modification{},
	}
	// Since there's no map copy...
	for stepName, mods := range atk.modifications {
		newMods := []interfaces.Modification{}
		// Don't copy transients
		for _, mod := range mods {
			t, ok := mod.(interfaces.Transient)
			if !ok || !t.IsTransient() {
				newMods = append(newMods, mod)
			}
		}
		cp.modifications[stepName] = newMods
	}
	return &cp
}

func (atk Attack) String() string {
	return fmt.Sprintf("<Attack attacker=%s defender=%s>", atk.attacker.Name(), atk.defender.Name())
}
