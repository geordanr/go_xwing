package dice

import (
	"fmt"
	"math/rand"
)

type Result uint8

const (
	CANCELED Result = iota
	BLANK
	FOCUS
	HIT
	CRIT
	EVADE
)

func (face Result) String() string {
	switch face {
	case BLANK:
		return "Blank"
	case FOCUS:
		return "Focus"
	case HIT:
		return "Hit"
	case CRIT:
		return "Crit"
	case EVADE:
		return "Evade"
	default:
		return "Invalid"
	}
}

type Die struct {
	result      Result
	locked      bool
	wasRerolled bool
}

type Rollable interface {
	Roll() Rollable
	Result() Result
	SetResult(Result) Result
	Locked() bool
	Lock() bool
	Unlock() bool
	Reroll() Rollable
	Rerolled() bool
	IsRerollable() bool
	String() string
}

func (die Die) String() string {
	s := fmt.Sprintf("[%s]", die.result)
	if die.wasRerolled {
		s = fmt.Sprintf("%s (rerolled)", s)
	}
	if die.locked {
		s = fmt.Sprintf("%s (locked)", s)
	}
	return s
}

type AttackDie struct {
	Die
}

func (die *AttackDie) Roll() Rollable {
	if die.IsRerollable() {
		face := uint8(rand.Int31n(8))
		switch {
		case face < 2:
			die.result = BLANK
		case face < 4:
			die.result = FOCUS
		case face < 7:
			die.result = HIT
		default:
			die.result = CRIT
		}
	}
	return die
}

func (die *AttackDie) Reroll() Rollable {
	die.Roll()
	die.wasRerolled = true
	return die
}

func (die *AttackDie) Result() Result {
	return die.result
}

func (die *AttackDie) SetResult(result Result) Result {
	if result != EVADE {
		die.result = result
	}
	return die.Result()
}

func (die *AttackDie) Locked() bool {
	return die.locked
}

func (die *AttackDie) Lock() bool {
	die.locked = true
	return die.locked
}

func (die *AttackDie) Unlock() bool {
	die.locked = false
	return die.locked
}

func (die *AttackDie) Rerolled() bool {
	return die.wasRerolled
}

func (die *AttackDie) IsRerollable() bool {
	return !(die.wasRerolled || die.locked)
}

type DefenseDie struct {
	Die
}

func (die *DefenseDie) Roll() Rollable {
	if die.IsRerollable() {
		face := uint8(rand.Int31n(8))
		switch {
		case face < 3:
			die.result = BLANK
		case face < 5:
			die.result = FOCUS
		default:
			die.result = EVADE
		}
	}
	return die
}

func (die *DefenseDie) Result() Result {
	return die.result
}

func (die *DefenseDie) SetResult(result Result) Result {
	if result != HIT && result != CRIT {
		die.result = result
	}
	return die.Result()
}

func (die *DefenseDie) Locked() bool {
	return die.locked
}

func (die *DefenseDie) Lock() bool {
	die.locked = true
	return die.locked
}

func (die *DefenseDie) Unlock() bool {
	die.locked = false
	return die.locked
}

func (die *DefenseDie) Rerolled() bool {
	return die.wasRerolled
}

func (die *DefenseDie) IsRerollable() bool {
	return !(die.wasRerolled || die.locked)
}

func (die *DefenseDie) Reroll() Rollable {
	die.Roll()
	die.wasRerolled = true
	return die
}
