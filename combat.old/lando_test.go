package combat

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/geordanr/go_xwing/attack"
    "github.com/geordanr/go_xwing/histogram"
    "github.com/geordanr/go_xwing/ship"
)

func TestLandoCrew(t *testing.T) {
    assert := assert.New(t)

    var rolledDoubleBlanks, rolledBlankFocus, rolledBlankEvade, rolledFocusEvade, rolledDoubleFocus, rolledDoubleEvade bool

    for i := 0; i < 10000; i++ {
	listOne := []ModifiedShip{
	    ModifiedShip{
		Ship: ship.Ship{
		    Name: "With Lando",
		    FocusTokens: 1,
		    Actions: []ship.Action{
			ship.Actions["Lando (Crew)"],
		    },
		},
	    },
	}

	listTwo := []ModifiedShip{
	    ModifiedShip{
		Ship: ship.Ship{
		    Name: "Without Lando",
		},
	    },
	}

	cbt, _ := New(func () ([]attack.Attack, error) {
	    return ListVersusList(listOne, listTwo), nil
	})

	combatStats := make(statsByShipName)
	combatResults := make(resultsByShipName)
	for name := range(cbt.combatants) {
	    combatStats[name] = new(stats)
	    combatResults[name] = new(simResult)
	    combatResults[name].HitHistogram = make(histogram.IntHistogram)
	    combatResults[name].CritHistogram = make(histogram.IntHistogram)
	    combatResults[name].HullHistogram = make(histogram.IntHistogram)
	    combatResults[name].ShieldHistogram = make(histogram.IntHistogram)
	}
	cbt.Execute(&combatStats, &combatResults)
	atk := cbt.attacks[0]
	attacker := atk.Attacker

	assert.True(attacker.FocusTokens >= 1)
	assert.True(attacker.FocusTokens + attacker.EvadeTokens <= 3)
	if attacker.FocusTokens == 1 && attacker.EvadeTokens == 0 { rolledDoubleBlanks = true }
	if attacker.FocusTokens == 2 && attacker.EvadeTokens == 0 { rolledBlankFocus = true }
	if attacker.FocusTokens == 1 && attacker.EvadeTokens == 1 { rolledBlankEvade = true }
	if attacker.FocusTokens == 2 && attacker.EvadeTokens == 1 { rolledFocusEvade = true }
	if attacker.FocusTokens == 3 && attacker.EvadeTokens == 0 { rolledDoubleFocus = true }
	if attacker.FocusTokens == 1 && attacker.EvadeTokens == 2 { rolledDoubleEvade = true }
    }
    assert.True(rolledDoubleBlanks)
    assert.True(rolledBlankFocus)
    assert.True(rolledBlankEvade)
    assert.True(rolledFocusEvade)
    assert.True(rolledDoubleFocus)
    assert.True(rolledDoubleEvade)
}
