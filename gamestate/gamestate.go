package gamestate

import (
	// "fmt"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/interfaces"
)

type GameState struct {
	teams          map[string]interfaces.Team
	attackQueue    []interfaces.Attack
	attackResults  *dice.Results
	defenseResults *dice.Results
	nextAttackStep interfaces.Step
}

func (state *GameState) EnqueueAttack(atk interfaces.Attack) {
	state.attackQueue = append(state.attackQueue, atk)
}

func (state GameState) CurrentAttack() interfaces.Attack {
	return state.attackQueue[len(state.attackQueue)-1]
}

func (state GameState) AttackResults() *dice.Results {
	return state.attackResults
}

func (state *GameState) SetAttackResults(r *dice.Results) {
	state.attackResults = r
}

func (state GameState) DefenseResults() *dice.Results {
	return state.defenseResults
}

func (state *GameState) SetDefenseResults(r *dice.Results) {
	state.defenseResults = r
}

func (state GameState) NextAttackStep() interfaces.Step {
	return state.nextAttackStep
}

func (state *GameState) SetNextAttackStep(step interfaces.Step) {
	state.nextAttackStep = step
}
