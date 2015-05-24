package attack

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/geordanr/go_xwing/dice"
    "github.com/geordanr/go_xwing/ship"
)

func TestLandoCrew(t *testing.T) {
    assert := assert.New(t)

    for i := 0; i < 10000; i++ {
	attacker := ship.Ship{
	    Name: "Attacker",
	    Hull: 1,
	}
	defender := ship.Ship{
	    Name: "Defender",
	    Hull: 1,
	}
	atk := Attack{
	    Attacker: &attacker,
	    NumAttackDice: 1,
	    AttackerModifications: []Modification{
		AttackDiceSetter{
		    desiredResults: []dice.Result{
			dice.HIT,
		    },
		},
	    },

	    Defender: &defender,
	    NumDefenseDice: 1,
	    DefenderModifications: []Modification{
		DefenseDiceSetter{
		    desiredResults: []dice.Result{
			dice.BLANK,
		    },
		},
		Modifications["Lando (Crew)"],
	    },
	}

	atk.Execute()
	assert.True(len(*atk.DefenseResults) < 4)
	if len(*atk.DefenseResults) > 1 {
	    assert.EqualValues(1, atk.DefenseResults.Evades() + atk.DefenseResults.Focuses())
	}
	if len(*atk.DefenseResults) > 2 {
	    assert.EqualValues(2, atk.DefenseResults.Evades() + atk.DefenseResults.Focuses())
	}
    }
}
