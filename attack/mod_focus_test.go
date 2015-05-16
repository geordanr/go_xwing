package attack

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/geordanr/go_xwing/dice"
    "github.com/geordanr/go_xwing/ship"
)

func TestUseOffensiveFocus_withoutFocuses(t *testing.T) {
    assert := assert.New(t)

    attacker := ship.Ship{
	Name: "Attacker",
    }
    defender := ship.Ship{
	Name: "Defender",
    }
    atk := Attack{
	Attacker: &attacker,
	NumAttackDice: 4,
	AttackerModifications: []Modification{
	    AttackDiceSetter{
		desiredResults: []dice.Result{
		    dice.BLANK,
		    dice.HIT,
		    dice.CRIT,
		    dice.FOCUS,
		},
	    },
	    Modifications["Offensive Focus"],
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
    assert.EqualValues(1, atk.AttackResults.Hits(), "Should still have 1 hit")
    assert.EqualValues(0, attacker.FocusTokens, "Should still have no focus tokens")
}

func TestUseOffensiveFocus_withFocuses(t *testing.T) {
    assert := assert.New(t)

    attacker := ship.Ship{
	Name: "Attacker",
	FocusTokens: 2,
    }
    defender := ship.Ship{
	Name: "Defender",
    }
    atk := Attack{
	Attacker: &attacker,
	NumAttackDice: 4,
	AttackerModifications: []Modification{
	    AttackDiceSetter{
		desiredResults: []dice.Result{
		    dice.BLANK,
		    dice.HIT,
		    dice.CRIT,
		    dice.FOCUS,
		},
	    },
	    Modifications["Offensive Focus"],
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
    assert.EqualValues(2, atk.AttackResults.Hits(), "Should now have 2 hits")
    assert.EqualValues(1, attacker.FocusTokens, "Should not have 1 focus token")
}

func TestUseDefensiveFocus_withoutFocuses(t *testing.T) {
    assert := assert.New(t)

    attacker := ship.Ship{
	Name: "Attacker",
    }
    defender := ship.Ship{
	Name: "Defender",
    }
    atk := Attack{
	Attacker: &attacker,
	NumAttackDice: 1,

	Defender: &defender,
	NumDefenseDice: 3,
	DefenderModifications: []Modification{
	    DefenseDiceSetter{
		desiredResults: []dice.Result{
		    dice.BLANK,
		    dice.EVADE,
		    dice.FOCUS,
		},
	    },
	    Modifications["Defensive Focus"],
	},
    }

    atk.Execute()
    assert.EqualValues(1, atk.DefenseResults.Evades(), "Should still have 1 evade result")
    assert.EqualValues(0, defender.FocusTokens, "Should still have 0 focus tokens")
}

func TestUseDefensiveFocus_withFocuses(t *testing.T) {
    assert := assert.New(t)

    attacker := ship.Ship{
	Name: "Attacker",
    }
    defender := ship.Ship{
	Name: "Defender",
	FocusTokens: 2,
    }
    atk := Attack{
	Attacker: &attacker,
	NumAttackDice: 2,
	AttackerModifications: []Modification{
	    AttackDiceSetter{
		desiredResults: []dice.Result{
		    dice.HIT,
		    dice.CRIT,
		},
	    },
	},

	Defender: &defender,
	NumDefenseDice: 3,
	DefenderModifications: []Modification{
	    DefenseDiceSetter{
		desiredResults: []dice.Result{
		    dice.BLANK,
		    dice.EVADE,
		    dice.FOCUS,
		},
	    },
	    Modifications["Defensive Focus"],
	},
    }

    atk.Execute()
    assert.EqualValues(2, atk.DefenseResults.Evades(), "Should have 2 evade results")
    assert.EqualValues(1, defender.FocusTokens, "Should have 1 focus token")
}

func TestUseDefensiveFocus_withFocusesIfNeeded(t *testing.T) {
    assert := assert.New(t)

    attacker := ship.Ship{
	Name: "Attacker",
    }
    defender := ship.Ship{
	Name: "Defender",
	FocusTokens: 2,
    }
    atk := Attack{
	Attacker: &attacker,
	NumAttackDice: 0,

	Defender: &defender,
	NumDefenseDice: 3,
	DefenderModifications: []Modification{
	    DefenseDiceSetter{
		desiredResults: []dice.Result{
		    dice.BLANK,
		    dice.FOCUS,
		    dice.FOCUS,
		},
	    },
	    Modifications["Defensive Focus"],
	},
    }

    atk.Execute()
    assert.EqualValues(0, atk.DefenseResults.Evades(), "Should have 0 evade results")
    assert.EqualValues(2, defender.FocusTokens, "Should still have 2 focus tokens")
}
