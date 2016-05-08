package gamestate

import (
	// "fmt"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/interfaces"
)

// GameState contains state that is transient for attacks.
// It should be reset at the beginning of each new attack (but not
// necessarily for things like repeated attacks).
//
// This may need to store transient per round stuff.
type GameState struct {
	combatants  map[string]interfaces.Ship
	attackQueue []interfaces.Attack
	// transient per attack
	attackResults       *dice.Results
	defenseResults      *dice.Results
	nextAttackStepName  string
	performAttackTwice  bool // can change mid-attack
	attackMissed        bool
	attackDiceModifier  int // amount to increase or decrease rolled attack dice
	defenseDiceModifier int // amount to increase or decrease rolled defense dice
	hitsLanded          uint
	critsLanded         uint
}

// Combatants returns a map of ship names to ships.
func (state GameState) Combatants() map[string]interfaces.Ship { return state.combatants }
func (state *GameState) SetCombatants(combatants map[string]interfaces.Ship) {
	state.combatants = combatants
}

// EnqueueAttack adds an attack to the front of the attack queue.
// Enqueue is the wrong term (wrong end).  Must come up with better name.
func (state *GameState) EnqueueAttack(atk interfaces.Attack) {
	state.attackQueue = append(state.attackQueue, atk)
}

// DequeueAttack removes the attack at the front of the attack queue.
// Returns true if there are still attacks afterward.
func (state *GameState) DequeueAttack() bool {
	if len(state.attackQueue) > 0 {
		state.attackQueue = state.attackQueue[:len(state.attackQueue)-1]
		return len(state.attackQueue) > 0
	} else {
		return false
	}
}

// SetNextAttack makes the given attack be the next one in the queue after the current attack.
func (state *GameState) SetNextAttack(atk interfaces.Attack) {
	curAtk := state.CurrentAttack()
	newQueue := []interfaces.Attack{}
	newQueue = append(newQueue, state.attackQueue[:len(state.attackQueue)-1]...)
	newQueue = append(newQueue, atk, curAtk)
	state.attackQueue = newQueue
}

// CurrentAttack returns the attack at the front of the queue.
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

func (state GameState) NextAttackStep() string {
	return state.nextAttackStepName
}

func (state *GameState) SetNextAttackStep(stepName string) {
	state.nextAttackStepName = stepName
}

func (state GameState) PerformAttackTwice() bool {
	return state.performAttackTwice
}

func (state *GameState) SetPerformAttackTwice(performAttackTwice bool) {
	state.performAttackTwice = performAttackTwice
}

func (state GameState) AttackMissed() bool {
	return state.attackMissed
}

func (state *GameState) SetAttackMissed(attackMissed bool) {
	state.attackMissed = attackMissed
}

func (state GameState) AttackDiceModifier() int {
	return state.attackDiceModifier
}

func (state *GameState) SetAttackDiceModifier(attackDiceModifier int) {
	state.attackDiceModifier = attackDiceModifier
}

func (state GameState) DefenseDiceModifier() int {
	return state.defenseDiceModifier
}

func (state *GameState) SetDefenseDiceModifier(defenseDiceModifier int) {
	state.defenseDiceModifier = defenseDiceModifier
}

func (state GameState) HitsLanded() uint {
	return state.hitsLanded
}

func (state *GameState) SetHitsLanded(hitsLanded uint) {
	state.hitsLanded = hitsLanded
}

func (state GameState) CritsLanded() uint {
	return state.critsLanded
}

func (state *GameState) SetCritsLanded(critsLanded uint) {
	state.critsLanded = critsLanded
}

func (state *GameState) ResetTransientState() {
	state.attackResults = nil
	state.defenseResults = nil
	state.nextAttackStepName = ""
	state.performAttackTwice = false
	state.attackMissed = false
	state.attackDiceModifier = 0
	state.defenseDiceModifier = 0
	state.hitsLanded = 0
	state.critsLanded = 0
}

func (state GameState) Copy() interfaces.GameState {
	var newState *GameState
	newState = &GameState{}
	*newState = state

	// copy over non-values
	// first, copy the combatants
	newState.combatants = make(map[string]interfaces.Ship)
	for name, combatant := range state.combatants {
		newState.combatants[name] = combatant.Copy()
	}
	// then remap them onto the attacks
	newState.attackQueue = make([]interfaces.Attack, len(state.attackQueue))
	for i, atk := range state.attackQueue {
		newState.attackQueue[i] = atk.Copy()
		attacker := atk.Attacker()
		defender := atk.Defender()
		newState.attackQueue[i].SetAttacker(newState.combatants[attacker.Name()])
		newState.attackQueue[i].SetDefender(newState.combatants[defender.Name()])
	}

	return newState
}
