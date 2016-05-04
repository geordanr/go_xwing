// Package step implements a single step of resolving an attack.
package step

import (
	// "fmt"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/interfaces"
)

type Step struct {
	name string
	mods []interfaces.Modification // default mods to run if not specified in attack
	next string // default next step to perform
}

/*
Run loops forever and reads a StepRequest from the input channel, runs
the modifications appropriate to this step, and sends a new StepRequest
to the output channel.

When the in channel is closed and no more StepRequests are waiting,
true is sent on the done channel.
*/
func (step Step) Run(in <-chan interfaces.StepRequest, out chan<- interfaces.StepRequest, done chan<- bool) {
	for {
		req, more := <-in
		if !more {
			close(out)
			done <- true
			return
		}

		state := req.State()
		// A nil next state means use the step default.
		state.SetNextAttackStep(nil)

		currentAttack := state.CurrentAttack()
		stepmods, exists := currentAttack.Modifications()[step.Name()]
		if !exists {
			if stepmods == nil {
				stepmods = []interfaces.Modification{}
			} else {
				stepmods = step.mods
			}
		}

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

		// No one overrode the next step, so use the default.
		if state.NextAttackStep() == nil {
			state.SetNextAttackStep(step.Next())
		}

		req.SetState(state)
		req.SetStep(&step)
		out <- req
	}
}

// Next returns the default next step to perform.
// This is used by the attack runner to figure out what step to perform
// next if one isn't provided to the GameState by a modification.
func (step Step) Next() interfaces.Step { return Steps[step.next] }

func (step Step) Name() string         { return step.name }
func (step *Step) SetName(name string) { step.name = name }