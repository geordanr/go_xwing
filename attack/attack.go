package attack

import (
	// "fmt"
	"github.com/geordanr/go_xwing/interfaces"
)

type Attack struct {
	attacker      interfaces.Ship
	defender      interfaces.Ship
	modifications map[string][]interfaces.Modification
}

func New(attacker, defender interfaces.Ship, modifications map[string][]interfaces.Modification) *Attack {
	atk := Attack{
		attacker:      attacker,
		defender:      defender,
		modifications: modifications,
	}
	return &atk
}

func (atk *Attack) Attacker() interfaces.Ship {
	return atk.attacker
}

func (atk *Attack) Defender() interfaces.Ship {
	return atk.defender
}

func (atk *Attack) Modifications() map[string][]interfaces.Modification {
	return atk.modifications
}
