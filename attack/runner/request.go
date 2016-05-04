package runner

import (
	"fmt"
	"github.com/geordanr/go_xwing/interfaces"
)

// Request represents a game state going to or from a step for processing.
type Request struct {
	state interfaces.GameState
	step  interfaces.Step
}

// State returns the game state.
func (r Request) State() interfaces.GameState { return r.state }

func (r *Request) SetState(state interfaces.GameState) { r.state = state }

// Step returns the step that processed this request.
func (r Request) Step() interfaces.Step { return r.step }

func (r *Request) SetStep(step interfaces.Step) { r.step = step }

func (r Request) String() string { return fmt.Sprintf("<Request state=%s step=%s>", r.state, r.step) }
