package attack

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/geordanr/go_xwing/dice"
    "github.com/geordanr/go_xwing/ship"
)

func TestCopy(t *testing.T) {
    assert := assert.New(t)

    attacker := ship.Ship{
	Name: "Attacker",
    }
    defender := ship.Ship{
	Name: "Defender",
    }
    src := Attack{
	Attacker: &attacker,
	NumAttackDice: 3,
	AttackerModifications: []Modification{
	    Modifications["Offensive Focus"],
	},

	Defender: &defender,
	NumDefenseDice: 2,
	DefenderModifications: []Modification{
	    Modifications["Defensive Focus"],
	},
    }

    dst := src.Copy()

    assert.Equal(src.Attacker, dst.Attacker, "Attacker should be the same")
    assert.Equal(src.NumAttackDice, dst.NumAttackDice, "NumAttackDice should be the same")
    assert.Equal(len(src.AttackerModifications), len(dst.AttackerModifications), "Length of attack modifications should be the same")
    for i, mod := range(src.AttackerModifications) {
	assert.Equal(mod, dst.AttackerModifications[i], "Attack modification at index %d", i)
    }

    assert.Equal(src.Defender, dst.Defender, "Defender should be the same")
    assert.Equal(src.NumDefenseDice, dst.NumDefenseDice, "NumDefenseDice should be the same")
    assert.Equal(len(src.DefenderModifications), len(dst.DefenderModifications), "Length of defense modifications should be the same")
    for i, mod := range(src.DefenderModifications) {
	assert.Equal(mod, dst.DefenderModifications[i], "Defender modification at index %d", i)
    }
}

func TestExecute(t *testing.T) {
    assert := assert.New(t)

    attacker := ship.Ship{
	Name: "Attacker",
	Hull: 3,
	Shields: 2,
	FocusTokens: 1,
    }
    defender := ship.Ship{
	Name: "Defender",
	Hull: 4,
	Shields: 1,
	EvadeTokens: 1,
    }
    atk := Attack{
	Attacker: &attacker,
	NumAttackDice: 3,
	AttackerModifications: []Modification{
	    AttackDiceSetter{
		desiredResults: []dice.Result{
		    dice.HIT,
		    dice.HIT,
		    dice.CRIT,
		    dice.FOCUS,
		},
	    },
	    Modifications["Offensive Focus"],
	},

	Defender: &defender,
	NumDefenseDice: 3,
	DefenderModifications: []Modification{
	    DefenseDiceSetter{
		desiredResults: []dice.Result{
		    dice.BLANK,
		    dice.FOCUS,
		    dice.EVADE,
		},
	    },
	    Modifications["Use Evade Token"],
	},
    }

    hits, crits := atk.Execute()

    assert.EqualValues(hits, 1)
    assert.EqualValues(crits, 1)
    assert.EqualValues(attacker.FocusTokens, 0)
    assert.EqualValues(defender.EvadeTokens, 0)
    assert.EqualValues(defender.Hull, 3)
    assert.EqualValues(defender.Shields, 0)

}
