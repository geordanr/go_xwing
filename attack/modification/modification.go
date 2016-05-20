/*
Package modification implements modification functions used when executing
the steps of an attack.

A modification has a ModifyState method which takes two arguments, the
game state and the ship to affect.  Modifications also have an actor
which is used to signify what role the ship is playing (usually the
attacker or the defender).

Modifications may access and modify any part of the state, including
attack results, the attack queue, and even what modifications will be
applied to downstream attack steps.
*/
package modification

import (
	"github.com/geordanr/go_xwing/constants"
)

type modification struct {
	actor constants.ModificationActor
}

func (mod modification) Actor() constants.ModificationActor         { return mod.actor }
func (mod modification) SetActor(actor constants.ModificationActor) { mod.actor = actor }
