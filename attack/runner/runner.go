// Package runner implements the logic that runs attack steps.
package runner

import (
	// "fmt"
	"github.com/geordanr/go_xwing/interfaces"
)

type Runner struct {
	sendChans map[string]chan interfaces.StepRequest
	recvChan  chan interfaces.StepRequest
	doneChans map[string]chan bool
}

// New returns a new instance of the step runner, using the Step map given.
// Bufsize should be the number of states expected to be passed through
// the runner.
func New(steps map[string]interfaces.Step, bufsize int) *Runner {
	recvChan := make(chan interfaces.StepRequest, bufsize)

	r := Runner{
		sendChans: make(map[string]chan interfaces.StepRequest),
		recvChan:  recvChan,
		doneChans: make(map[string]chan bool),
	}

	for name, step := range steps {
		r.sendChans[name] = make(chan interfaces.StepRequest, bufsize)
		r.doneChans[name] = make(chan bool, bufsize)
		go step.Run(r.sendChans[name], recvChan, r.doneChans[name])
	}

	return &r
}

// Run processes game states until every injected game state has an empty attack queue.
// Completed game states are sent to the given channel.
func (r Runner) Run(output chan<- interfaces.GameState) {
	for {
		req := <-r.recvChan
		state := req.State()
		if state.NextAttackStep() == "" {
			// End of attack sequence, process next attack
			more := state.DequeueAttack()
			if more {
				// Send the new state back into the machine
				state.ResetTransientState()
				r.InjectState(state)
			} else {
				// No more attacks, this state is done
				output <- state
			}
		} else {
			// Send the state to the next step
			newReq := Request{state: state}
			// This depends on the names of the steps being the same as the keys in the step map
			r.sendChans[state.NextAttackStep()] <- &newReq
		}
	}
}

// InjectState sends the given state into the runner for processing.
func (r Runner) InjectState(state interfaces.GameState) {
	req := Request{state: state}
	r.sendChans["__START__"] <- &req
}
