package dice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRollable_AttackDie(t *testing.T) {
	assert := assert.New(t)

	var attackDie AttackDie
	var blanks, focuses, hits, crits int
	attackDie.Roll()

	for i := 0; i < 1000; i++ {
		switch attackDie.Result() {
		case BLANK:
			blanks++
		case FOCUS:
			focuses++
		case HIT:
			hits++
		case CRIT:
			crits++
		}
	}

	assert.InEpsilon(int(1000*2.0/8), blanks, 50)
	assert.InEpsilon(int(1000*2.0/8), focuses, 50)
	assert.InEpsilon(int(1000*3.0/8), hits, 50)
	assert.InEpsilon(int(1000*1.0/8), crits, 50)
}

func TestRollable_DefenseDie(t *testing.T) {
	assert := assert.New(t)

	var defenseDie DefenseDie
	var blanks, focuses, evades int
	defenseDie.Roll()

	for i := 0; i < 1000; i++ {
		switch defenseDie.Result() {
		case BLANK:
			blanks++
		case FOCUS:
			focuses++
		case EVADE:
			evades++
		}
	}

	assert.InEpsilon(int(1000*3.0/8), blanks, 50)
	assert.InEpsilon(int(1000*2.0/8), focuses, 50)
	assert.InEpsilon(int(1000*3.0/8), evades, 50)
}
