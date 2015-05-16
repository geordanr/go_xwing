package attack

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/geordanr/go_xwing/dice"
    "github.com/geordanr/go_xwing/ship"
)

func TestUseEvadeToken_withoutEvadeToken(t *testing.T) {
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
	    Modifications["Use Evade Token"],
	},
    }

    atk.Execute()
    assert.EqualValues(0, atk.DefenseResults.Evades(), "Should not have evade results without evade token")
}

func TestUseEvadeToken_withEvadeTokens(t *testing.T) {
    assert := assert.New(t)

    attacker := ship.Ship{
	Name: "Attacker",
	Hull: 1,
    }
    defender := ship.Ship{
	Name: "Defender",
	Hull: 1,
	EvadeTokens: 2,
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
	    Modifications["Use Evade Token"],
	},
    }

    atk.Execute()
    assert.EqualValues(1, atk.DefenseResults.Evades(), "Should have 1 evade result")
    assert.EqualValues(1, defender.EvadeTokens, "Should have 1 evade token left")
}

func TestUseEvadeToken_spendWhenNecessary(t *testing.T) {
    assert := assert.New(t)

    attacker := ship.Ship{
	Name: "Attacker",
	Hull: 1,
    }
    defender := ship.Ship{
	Name: "Defender",
	Hull: 1,
	EvadeTokens: 2,
    }
    atk := Attack{
	Attacker: &attacker,
	NumAttackDice: 1,
	AttackerModifications: []Modification{
	    AttackDiceSetter{
		desiredResults: []dice.Result{
		    dice.BLANK,
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
	    Modifications["Use Evade Token"],
	},
    }

    atk.Execute()
    assert.EqualValues(0, atk.DefenseResults.Evades(), "Should have no evade results")
    assert.EqualValues(2, defender.EvadeTokens, "Should have 2 evade tokens left")
}
