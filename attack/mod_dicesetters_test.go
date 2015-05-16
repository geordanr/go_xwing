package attack

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/geordanr/go_xwing/dice"
    "github.com/geordanr/go_xwing/ship"
)

func TestAttackDiceSetter(t *testing.T) {
    assert := assert.New(t)

    attacker := ship.Ship{
	Name: "Attacker",
    }
    defender := ship.Ship{
	Name: "Defender",
    }

    setter := AttackDiceSetter{
	desiredResults: []dice.Result{
	    dice.HIT,
	    dice.CRIT,
	    dice.FOCUS,
	    dice.BLANK,
	    dice.HIT,
	},
    }

    atk := Attack{
	Attacker: &attacker,
	NumAttackDice: 4,
	AttackerModifications: []Modification{
	    setter,
	},

	Defender: &defender,
	NumDefenseDice: 0,
    }

    hits, crits := atk.Execute()

    assert.EqualValues(5, len(*atk.AttackResults))
    assert.EqualValues(1, crits)
    assert.EqualValues(2, hits)
    assert.EqualValues(1, atk.AttackResults.Focuses())
    assert.EqualValues(1, atk.AttackResults.Blanks())
}

func TestDefenseDiceSetter(t *testing.T) {
    assert := assert.New(t)

    attacker := ship.Ship{
	Name: "Attacker",
    }
    defender := ship.Ship{
	Name: "Defender",
    }

    setter := DefenseDiceSetter{
	desiredResults: []dice.Result{
	    dice.EVADE,
	    dice.FOCUS,
	    dice.BLANK,
	    dice.EVADE,
	},
    }

    atk := Attack{
	Attacker: &attacker,
	NumAttackDice: 4,

	Defender: &defender,
	NumDefenseDice: 0,
	DefenderModifications: []Modification{
	    setter,
	},
    }

    atk.Execute()

    assert.EqualValues(4, len(*atk.DefenseResults))
    assert.EqualValues(1, atk.DefenseResults.Focuses())
    assert.EqualValues(2, atk.DefenseResults.Evades())
    assert.EqualValues(1, atk.DefenseResults.Blanks())
}
