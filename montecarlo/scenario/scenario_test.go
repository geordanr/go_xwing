package scenario

import (
    "testing"
    "github.com/geordanr/xwing/dice"
)

func TestCompareResults_hitsCancelBeforeCrits(t *testing.T) {
    attackResults := dice.RollAttackDice(4)
    attackResults[0].SetResult(dice.BLANK)
    attackResults[1].SetResult(dice.HIT)
    attackResults[2].SetResult(dice.CRIT)
    attackResults[3].SetResult(dice.HIT)

    defenseResults := dice.RollDefenseDice(4)
    defenseResults[0].SetResult(dice.BLANK)
    defenseResults[1].SetResult(dice.FOCUS)
    defenseResults[3].SetResult(dice.BLANK)
    defenseResults[2].SetResult(dice.EVADE)

    scenario := Scenario {
	AttackResults: &attackResults,
	DefenseResults: &defenseResults,
    }

    hits, crits := scenario.CompareResults()

    if hits != 1 {
	// t.Logf("Attack: %s\n", scenario.AttackResults)
	// t.Logf("Defense: %s\n", scenario.DefenseResults)
	// t.Logf("Hits: %d, Crits: %d\n", hits, crits)
	t.Errorf("Evades should cancel hits first")
    }

    if crits != 1 {
	// t.Logf("Attack: %s\n", scenario.AttackResults)
	// t.Logf("Defense: %s\n", scenario.DefenseResults)
	// t.Logf("Hits: %d, Crits: %d\n", hits, crits)
	t.Errorf("Crit should have made it through")
    }
}

func TestCompareResults_allTheEvades(t *testing.T) {
    attackResults := dice.RollAttackDice(4)
    attackResults[0].SetResult(dice.BLANK)
    attackResults[1].SetResult(dice.HIT)
    attackResults[2].SetResult(dice.CRIT)
    attackResults[3].SetResult(dice.HIT)

    defenseResults := dice.RollDefenseDice(4)
    defenseResults[0].SetResult(dice.EVADE)
    defenseResults[1].SetResult(dice.FOCUS)
    defenseResults[3].SetResult(dice.EVADE)
    defenseResults[2].SetResult(dice.EVADE)

    scenario := Scenario {
	AttackResults: &attackResults,
	DefenseResults: &defenseResults,
    }

    hits, crits := scenario.CompareResults()

    if hits != 0 {
	// t.Logf("Attack: %s\n", scenario.AttackResults)
	// t.Logf("Defense: %s\n", scenario.DefenseResults)
	// t.Logf("Hits: %d, Crits: %d\n", hits, crits)
	t.Errorf("Evades should have canceled all hits")
    }

    if crits != 0 {
	// t.Logf("Attack: %s\n", scenario.AttackResults)
	// t.Logf("Defense: %s\n", scenario.DefenseResults)
	// t.Logf("Hits: %d, Crits: %d\n", hits, crits)
	t.Errorf("Evades should have canceled all crits")
    }
}
