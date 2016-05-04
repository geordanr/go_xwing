// Package step implements a single step of resolving an attack.
package step

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/interfaces"
)

type Step struct {
	Name     string
}

/*
Run loops forever and reads a GameState from the input channel, runs
the modifications appropriate to this step, and sends a new GameState
to the output channel.

When the in channel is closed and no more GameStates are waiting,
true is send on the done channel.
*/
func (step *Step) Run(in <-chan interfaces.GameState, out chan<- interfaces.GameState, done chan<- bool) {
	for {
		state, more := <-in

		if !more {
			close(out)
			done <- true
			return
		}

		currentAttack := state.CurrentAttack()
		stepmods, exists := currentAttack.Modifications()[step.Name]
		if exists {
			for _, mod := range stepmods {
				var ship interfaces.Ship
				switch mod.Actor() {
				case constants.ATTACKER:
					ship = currentAttack.Attacker()
				case constants.DEFENDER:
					ship = currentAttack.Defender()
				}
				mod.ModifyState(state, ship)
			}
		}

		out <- state
	}
}
