package combat

import (
    "path/filepath"
    "runtime"
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestAttacksFromJSONPath(t *testing.T) {
    assert := assert.New(t)

    _, thisfile, _, _ := runtime.Caller(0)
    thisdir := filepath.Dir(thisfile)
    attacks, iterations, err := AttacksFromJSONPath(filepath.Join(thisdir, "sample.json"))
    if err != nil {
	t.Fatal(err)
    }

    assert.EqualValues(10000, iterations)

    assert.EqualValues(3, len(attacks))

    // Attack 0: Soontir vs. Blue Squad
    assert.EqualValues("Soontir Fel", attacks[0].Attacker.Name)
    assert.EqualValues(2, attacks[0].Attacker.FocusTokens)
    assert.EqualValues(1, attacks[0].Attacker.EvadeTokens)
    assert.EqualValues(3, attacks[0].NumAttackDice)
    assert.EqualValues(1, attacks[0].NumDefenseDice)
    assert.EqualValues("Offensive Focus", attacks[0].AttackerModifications[0].String())
    assert.EqualValues("Defensive Focus", attacks[0].DefenderModifications[0].String())

    // Attack 1: Blue Squad vs. Soontir
    assert.EqualValues("Blue Squad", attacks[1].Attacker.Name)
    assert.EqualValues(1, attacks[1].Attacker.FocusTokens)
    assert.EqualValues(0, attacks[1].Attacker.EvadeTokens)
    assert.EqualValues(3, attacks[1].NumAttackDice)
    assert.EqualValues(4, attacks[1].NumDefenseDice)
    assert.EqualValues(0, len(attacks[1].AttackerModifications))
    assert.EqualValues("Use Evade Token", attacks[1].DefenderModifications[0].String())
    assert.EqualValues("Defensive Focus", attacks[1].DefenderModifications[1].String())

    // Attack 2: Green Squad vs. Soontir
    assert.EqualValues("Green Squad", attacks[2].Attacker.Name)
    assert.EqualValues(0, attacks[2].Attacker.FocusTokens)
    assert.EqualValues(1, attacks[2].Attacker.EvadeTokens)
    assert.EqualValues(2, attacks[2].Attacker.Hull)
    assert.EqualValues(2, attacks[2].NumAttackDice)
    assert.EqualValues(4, attacks[2].NumDefenseDice)
    assert.EqualValues("Target Lock", attacks[2].AttackerModifications[0].String())
    assert.EqualValues("Use Evade Token", attacks[2].DefenderModifications[0].String())
    assert.EqualValues("Defensive Focus", attacks[2].DefenderModifications[1].String())
}

func TestSimulateFromJSONPath(t *testing.T) {
    var prs bool

    assert := assert.New(t)

    _, thisfile, _, _ := runtime.Caller(0)
    thisdir := filepath.Dir(thisfile)
    shipStats, shipResults, err := SimulateFromJSONPath(filepath.Join(thisdir, "sample.json"))
    if err != nil {
	t.Fatal(err)
    }

    _, prs = (*shipStats)["Soontir Fel"]
    assert.True(prs)
    _, prs = (*shipResults)["Soontir Fel"]
    assert.True(prs)

    _, prs = (*shipStats)["Blue Squad"]
    assert.True(prs)
    _, prs = (*shipResults)["Blue Squad"]
    assert.True(prs)

    _, prs = (*shipStats)["Green Squad"]
    assert.True(prs)
    _, prs = (*shipResults)["Green Squad"]
    assert.True(prs)

}
