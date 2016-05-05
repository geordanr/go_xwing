package interfaces

import (
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/dice"
)

type Attack interface {
	Attacker() Ship
	Defender() Ship
	Modifications() map[string]([]Modification)
}

type GameState interface {
	EnqueueAttack(Attack)
	DequeueAttack() bool
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
}

type Modification interface {
	Actor() constants.ModificationActor
	SetActor(constants.ModificationActor)
	ModifyState(GameState, Ship)
}

type Ship interface {
	Name() string
	Attack() uint
	Agility() uint
	Hull() uint
	Shields() uint

	FocusTokens() uint
	SetFocusTokens(uint)
	EvadeTokens() uint
	SetEvadeTokens(uint)

	SpendFocus() bool
	SpendEvade() bool
	IsAlive() bool
	CanAttack() bool
	SufferDamage(uint)
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
