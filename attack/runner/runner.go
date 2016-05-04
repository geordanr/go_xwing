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
func New(steps map[string]interfaces.Step) *Runner {
	recvChan := make(chan interfaces.StepRequest)

	r := Runner{
		sendChans: make(map[string]chan interfaces.StepRequest),
		recvChan:  recvChan,
		doneChans: make(map[string]chan bool),
	}

	for name, step := range steps {
		r.sendChans[name] = make(chan interfaces.StepRequest)
		r.doneChans[name] = make(chan bool)
		go step.Run(r.sendChans[name], recvChan, r.doneChans[name])
	}

	return &r
}

func (r Runner) Run(state interfaces.GameState) interfaces.GameState {
	req := Request{state: state}
	r.sendChans["__START__"] <- &req
	for {
		req := <-r.recvChan
		state := req.State()
		// Are we at an end state?
		if state.NextAttackStep() == nil {
			return state
		} else {
			// Send the state to the next step
			newReq := Request{state: state}
			r.sendChans[state.NextAttackStep().Name()] <- &newReq
		}
	}
}
