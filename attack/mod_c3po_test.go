package attack

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/geordanr/go_xwing/dice"
    "github.com/geordanr/go_xwing/ship"
)

func TestC3PO_guessNoneAndBeWrong(t *testing.T) {
    assert := assert.New(t)

    attacker := ship.Ship{
	Name: "Attacker",
    }
    defender := ship.Ship{
	Name: "Defender",
    }
    atk := Attack{
	Attacker: &attacker,
	NumAttackDice: 0,

	Defender: &defender,
	NumDefenseDice: 4,
	DefenderModifications: []Modification{
	    DefenseDiceSetter{
		desiredResults: []dice.Result{
		    dice.EVADE,
		    dice.EVADE,
		    dice.FOCUS,
		    dice.BLANK,
		},
	    },
	    Modifications["C-3PO (guess 0)"],
	},
    }

    atk.Execute()
    assert.EqualValues(2, atk.DefenseResults.Evades(), "Should still have 2 evades")
}

func TestC3PO_guessNoneAndBeRight(t *testing.T) {
    assert := assert.New(t)

    attacker := ship.Ship{
	Name: "Attacker",
    }
    defender := ship.Ship{
	Name: "Defender",
    }
    atk := Attack{
	Attacker: &attacker,
	NumAttackDice: 0,

	Defender: &defender,
	NumDefenseDice: 4,
	DefenderModifications: []Modification{
	    DefenseDiceSetter{
		desiredResults: []dice.Result{
		    dice.BLANK,
		    dice.FOCUS,
		    dice.FOCUS,
		    dice.BLANK,
		},
	    },
	    Modifications["C-3PO (guess 0)"],
	},
    }

    atk.Execute()
    assert.EqualValues(1, atk.DefenseResults.Evades(), "Should now have 1 evade")
}

func TestC3PO_guessOneAndBeWrong(t *testing.T) {
    assert := assert.New(t)

    attacker := ship.Ship{
	Name: "Attacker",
    }
    defender := ship.Ship{
	Name: "Defender",
    }
    atk := Attack{
	Attacker: &attacker,
	NumAttackDice: 0,

	Defender: &defender,
	NumDefenseDice: 4,
	DefenderModifications: []Modification{
	    DefenseDiceSetter{
		desiredResults: []dice.Result{
		    dice.EVADE,
		    dice.EVADE,
		    dice.FOCUS,
		    dice.BLANK,
		},
	    },
	    Modifications["C-3PO (guess 1)"],
	},
    }

    atk.Execute()
    assert.EqualValues(2, atk.DefenseResults.Evades(), "Should still have 2 evades")
}

func TestC3PO_guessOneAndBeRight(t *testing.T) {
    assert := assert.New(t)

    attacker := ship.Ship{
	Name: "Attacker",
    }
    defender := ship.Ship{
	Name: "Defender",
    }
    atk := Attack{
	Attacker: &attacker,
	NumAttackDice: 0,

	Defender: &defender,
	NumDefenseDice: 4,
	DefenderModifications: []Modification{
	    DefenseDiceSetter{
		desiredResults: []dice.Result{
		    dice.EVADE,
		    dice.FOCUS,
		    dice.FOCUS,
		    dice.BLANK,
		},
	    },
	    Modifications["C-3PO (guess 1)"],
	},
    }

    atk.Execute()
    assert.EqualValues(2, atk.DefenseResults.Evades(), "Should now have 2 evades")
}
