package attack
import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/geordanr/go_xwing/dice"
	"github.com/geordanr/go_xwing/ship"
)


func makeLukeAttack( attackerDice []dice.Result, defenderDice []dice.Result ) Attack {

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
		NumAttackDice: 3,
		AttackerModifications: []Modification{
			AttackDiceSetter{
				desiredResults: attackerDice,
			},
		},

		Defender: &defender,
		NumDefenseDice: 2,
		DefenderModifications: []Modification{
			DefenseDiceSetter{ desiredResults: defenderDice },
			Modifications["Luke Skywalker (Pilot)"],

		},
	}
	return atk
}


func TestLukeSkywalker( t * testing.T ) {

	defenderDice := []dice.Result{
		dice.BLANK,
		dice.BLANK,
	}

	attackDice := []dice.Result{
		dice.HIT,
		dice.HIT,
		dice.HIT,
	}

	atk := makeLukeAttack( attackDice, defenderDice )

	atk.Execute()

	if atk.DefenseResults.Focuses() > 0 {
		t.Errorf("Should have no focuses")
	}
	if atk.DefenseResults.Evades() > 0 {
		t.Errorf("Should have no evades")
	}
	defenderDice[0] = dice.FOCUS
	defenderDice[1] = dice.BLANK

	atk2 := makeLukeAttack( attackDice, defenderDice )


	atk2.Execute()

	if atk2.DefenseResults.Focuses() > 0 {
		t.Errorf("Should have no focuses")
	}
	if atk2.DefenseResults.Evades() != 1 {
		t.Errorf("Should have evade")
	}

	defenderDice[0] = dice.FOCUS
	defenderDice[1] = dice.FOCUS

	atk3 := makeLukeAttack( attackDice, defenderDice )

	atk3.Execute()

	if atk3.DefenseResults.Focuses() != 1 {
		t.Errorf("Should only one focuses")
	}
	if atk2.DefenseResults.Evades() != 1 {
		t.Errorf("Should have only evade")
	}


}

func TestAutothrusters_withBlanks(t *testing.T) {
    assert := assert.New(t)

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

	Defender: &defender,
	NumDefenseDice: 4,
	DefenderModifications: []Modification{
	    DefenseDiceSetter{
		desiredResults: []dice.Result{
		    dice.BLANK,
		    dice.BLANK,
		    dice.FOCUS,
		    dice.EVADE,
		},
	    },
	    Modifications["Autothrusters"],
	},
    }

    atk.Execute()
    assert.EqualValues(2, atk.DefenseResults.Evades())
    assert.EqualValues(1, atk.DefenseResults.Blanks())
    assert.EqualValues(1, atk.DefenseResults.Focuses())
}

func TestAutothrusters_withoutBlanks(t *testing.T) {
    assert := assert.New(t)

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

	Defender: &defender,
	NumDefenseDice: 1,
	DefenderModifications: []Modification{
	    DefenseDiceSetter{
		desiredResults: []dice.Result{
		    dice.FOCUS,
		    dice.EVADE,
		},
	    },
	    Modifications["Autothrusters"],
	},
    }

    atk.Execute()
    assert.EqualValues(1, atk.DefenseResults.Evades())
    assert.EqualValues(0, atk.DefenseResults.Blanks())
    assert.EqualValues(1, atk.DefenseResults.Focuses())
}

func TestManglerCannon(t *testing.T) {
    assert := assert.New(t)

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
	AttackerModifications: []Modification{
	    AttackDiceSetter{
		desiredResults: []dice.Result{
		    dice.BLANK,
		    dice.HIT,
		    dice.CRIT,
		    dice.FOCUS,
		},
	    },
	    Modifications["Mangler Cannon"],
	},

	Defender: &defender,
    }

    atk.Execute()
    assert.EqualValues(0, atk.AttackResults.Hits())
    assert.EqualValues(2, atk.AttackResults.Crits())
    assert.EqualValues(1, atk.AttackResults.Blanks())
    assert.EqualValues(1, atk.AttackResults.Focuses())
}

func TestHeavyLaserCannon(t *testing.T) {
    assert := assert.New(t)

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
	AttackerModifications: []Modification{
	    AttackDiceSetter{
		desiredResults: []dice.Result{
		    dice.BLANK,
		    dice.HIT,
		    dice.CRIT,
		    dice.FOCUS,
		},
	    },
	    Modifications["Heavy Laser Cannon"],
	},

	Defender: &defender,
    }

    atk.Execute()
    assert.EqualValues(2, atk.AttackResults.Hits())
    assert.EqualValues(0, atk.AttackResults.Crits())
    assert.EqualValues(1, atk.AttackResults.Blanks())
    assert.EqualValues(1, atk.AttackResults.Focuses())
}
