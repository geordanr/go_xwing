package attack

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/geordanr/go_xwing/dice"
    "github.com/geordanr/go_xwing/ship"
)

func TestGunner_firstHitLands(t *testing.T) {
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
	    Modifications["Gunner"],
	    AttackDiceSetter{
		desiredResults: []dice.Result{
		    dice.CRIT,
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
	},
    }

    hits, crits := atk.Execute()
    assert.False(atk.IsGunnerAttack)
    assert.EqualValues(0, hits)
    assert.EqualValues(1, crits)
}

func TestGunner_firstHitMisses(t *testing.T) {
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
	    Modifications["Gunner"],
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
		    dice.EVADE,
		},
	    },
	},
    }

    atk.Execute()
    assert.True(atk.IsGunnerAttack)
}

func TestGunner_trackTokenSpend(t *testing.T) {
    assert := assert.New(t)

    attacker := ship.Ship{
	Name: "Attacker",
	FocusTokens: 2,
	Hull: 1,
    }
    defender := ship.Ship{
	Name: "Defender",
	Hull: 1,
    }
    atk := Attack{
	Attacker: &attacker,
	NumAttackDice: 2,
	AttackerModifications: []Modification{
	    Modifications["Gunner"],
	    AttackDiceSetter{
		desiredResults: []dice.Result{
		    dice.FOCUS,
		    dice.BLANK,
		},
	    },
	    GunnerAttackDiceSetter{
		desiredResults: []dice.Result{
		    dice.FOCUS,
		    dice.CRIT,
		},
	    },
	    Modifications["Offensive Focus"],
	},

	Defender: &defender,
	NumDefenseDice: 1,
	DefenderModifications: []Modification{
	    DefenseDiceSetter{
		desiredResults: []dice.Result{
		    dice.EVADE,
		},
	    },
	},
    }

    hits, crits := atk.Execute()
    assert.True(atk.IsGunnerAttack)
    assert.EqualValues(0, hits)
    assert.EqualValues(1, crits)
    assert.EqualValues(0, attacker.FocusTokens)
}


func TestLukeSkywalker_firstHitLands(t *testing.T) {
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
	    Modifications["Luke Skywalker"],
	    AttackDiceSetter{
		desiredResults: []dice.Result{
		    dice.CRIT,
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
	},
    }

    hits, crits := atk.Execute()
    assert.False(atk.IsGunnerAttack)
    assert.EqualValues(0, hits)
    assert.EqualValues(1, crits)

}

func TestLukeSkywalker_firstHitMisses(t *testing.T) {
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
	NumAttackDice: 3,
	AttackerModifications: []Modification{
	    AttackDiceSetter{
		desiredResults: []dice.Result{
		    dice.BLANK,
		    dice.BLANK,
		    dice.BLANK,
		},
	    },
	    GunnerAttackDiceSetter{
		desiredResults: []dice.Result{
		    dice.BLANK,
		    dice.FOCUS,
		    dice.CRIT,
		},
	    },
	    Modifications["Luke Skywalker"],
	},

	Defender: &defender,
	NumDefenseDice: 1,
	DefenderModifications: []Modification{
	    DefenseDiceSetter{
		desiredResults: []dice.Result{
		    dice.BLANK,
		},
	    },
	},
    }

    hits, crits := atk.Execute()
    assert.True(atk.IsGunnerAttack)
    assert.EqualValues(1, hits)
    assert.EqualValues(1, crits)

}
