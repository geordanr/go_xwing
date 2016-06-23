package interfaces

import (
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/dice"
)

type Attack interface {
	Attacker() Ship
	SetAttacker(Ship)
	Defender() Ship
	SetDefender(Ship)
	Modifications() map[string]([]Modification)
	SetModifications(map[string]([]Modification))
	Copy() Attack
}

type GameState interface {
	Combatants() map[string]Ship
	SetCombatants(map[string]Ship)
	EnqueueAttack(Attack)
	DequeueAttack() bool
	SetNextAttack(Attack)
	CurrentAttack() Attack
	AttackResults() *dice.Results
	SetAttackResults(*dice.Results)
	DefenseResults() *dice.Results
	SetDefenseResults(*dice.Results)
	NextAttackStep() string
	SetNextAttackStep(string)
	PerformAttackTwice() bool
	SetPerformAttackTwice(bool)
	AttackDiceModifier() int
	SetAttackDiceModifier(int)
	DefenseDiceModifier() int
	SetDefenseDiceModifier(int)
	AttackMissed() bool
	SetAttackMissed(bool)
	ResetTransientState()
	HitsLanded() uint
	SetHitsLanded(uint)
	CritsLanded() uint
	SetCritsLanded(uint)
	Copy() GameState
	ImportCombatants(map[string]Ship, bool)
	HasDeadCombatant() bool
}

type Modification interface {
	Actor() constants.ModificationActor
	SetActor(constants.ModificationActor)
	ModifyState(GameState, Ship)
	String() string
}

type Ship interface {
	Name() string
	Skill() uint
	Attack() uint
	Agility() uint
	Hull() uint
	Shields() uint

	FocusTokens() uint
	SetFocusTokens(uint)
	EvadeTokens() uint
	SetEvadeTokens(uint)
	TargetLock() string
	SetTargetLock(string)
	CopyTokensFrom(Ship)

	SpendFocus() bool
	SpendEvade() bool
	IsAlive() bool
	CanAttack() bool
	SetCanAttack(bool)
	SufferDamage(uint, uint)

	Copy() Ship
}

type Step interface {
	Next() string
	Run(<-chan StepRequest, chan<- StepRequest, chan<- bool)
	Name() string
	SetName(string)
	Mods() []Modification
}

type StepRequest interface {
	State() GameState
	SetState(GameState)
	Step() Step
	SetStep(Step)
}

type Transient interface {
	IsTransient() bool
}
