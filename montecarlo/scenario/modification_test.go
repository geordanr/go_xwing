package scenario

import (
    "testing"
    "github.com/geordanr/go_xwing/dice"
)

func TestUseOffensiveFocus_withoutFocuses(t *testing.T) {
    attackResults := dice.RollAttackDice(4)
    defenseResults := dice.RollDefenseDice(1)
    attackResults[0].SetResult(dice.BLANK)
    attackResults[1].SetResult(dice.HIT)
    attackResults[2].SetResult(dice.CRIT)
    attackResults[3].SetResult(dice.FOCUS)
    scenario := Scenario {
	AttackResults: &attackResults,
	DefenseResults: &defenseResults,
    }
    scenario.AttackerModifiesAttackDice = append(scenario.AttackerModifiesAttackDice, Modifications["Offensive Focus"])

    if scenario.AttackResults.Hits() != 1 {
	t.Fatalf("Should have only one hit")
    }

    scenario.Run()

    if scenario.AttackResults.Hits() != 1 {
	t.Errorf("Should still have only one hit")
    }

    if scenario.NumAttackerFocus != 0 {
	t.Errorf("Should still have no focus tokens")
    }
}

func TestUseOffensiveFocus_withFocuses(t *testing.T) {
    attackResults := dice.RollAttackDice(4)
    defenseResults := dice.RollDefenseDice(1)
    attackResults[0].SetResult(dice.BLANK)
    attackResults[1].SetResult(dice.HIT)
    attackResults[2].SetResult(dice.CRIT)
    attackResults[3].SetResult(dice.FOCUS)
    scenario := Scenario {
	AttackResults: &attackResults,
	NumAttackerFocus: 2,
	DefenseResults: &defenseResults,
    }
    scenario.AttackerModifiesAttackDice = append(scenario.AttackerModifiesAttackDice, Modifications["Offensive Focus"])

    if scenario.AttackResults.Hits() != 1 {
	t.Fatalf("Should have only one hit")
    }

    scenario.Run()

    if scenario.AttackResults.Hits() != 2 {
	t.Errorf("Should have two hits")
    }

    if scenario.NumAttackerFocus != 1 {
	t.Errorf("Should have one focus left")
    }
}

func TestTargetLock(t *testing.T) {
    attackResults := dice.RollAttackDice(8)
    defenseResults := dice.RollDefenseDice(8)
    scenario := Scenario {
	AttackResults: &attackResults,
	NumAttackerFocus: 1,
	DefenseResults: &defenseResults,
	NumDefenderEvade: 1,
    }

    t.Logf("Attack before: %s\n", scenario.AttackResults)
    scenario.AttackerModifiesAttackDice = append(scenario.AttackerModifiesAttackDice, Modifications["Target Lock"])

    scenario.Run()
    t.Logf("Attack after: %s\n", scenario.AttackResults)
}

func TestHowlrunner(t *testing.T) {
    attackResults := dice.RollAttackDice(8)
    defenseResults := dice.RollDefenseDice(1)
    scenario := Scenario {
	AttackResults: &attackResults,
	DefenseResults: &defenseResults,
    }
    t.Logf("Attack before: %s\n", scenario.AttackResults)
    scenario.AttackerModifiesAttackDice = append(scenario.AttackerModifiesAttackDice, Modifications["Howlrunner"])

    scenario.Run()
    t.Logf("Attack after: %s\n", scenario.AttackResults)
}

func TestUseDefensiveFocus_withoutFocuses(t *testing.T) {
    attackResults := dice.RollAttackDice(1)
    defenseResults := dice.RollDefenseDice(3)
    defenseResults[0].SetResult(dice.BLANK)
    defenseResults[1].SetResult(dice.EVADE)
    defenseResults[2].SetResult(dice.FOCUS)
    scenario := Scenario {
	AttackResults: &attackResults,
	DefenseResults: &defenseResults,
    }
    scenario.DefenderModifiesDefenseDice = append(scenario.DefenderModifiesDefenseDice, Modifications["Defensive Focus"])

    if scenario.DefenseResults.Evades() != 1 {
	t.Fatalf("Should have only one evade")
    }

    scenario.Run()

    if scenario.DefenseResults.Evades() != 1 {
	t.Errorf("Should still have only one evade")
    }

    if scenario.NumDefenderFocus != 0 {
	t.Errorf("Should still have no focuses")
    }
}


func TestUseDefensiveFocus_withFocuses(t *testing.T) {
    attackResults := dice.RollAttackDice(1)
    defenseResults := dice.RollDefenseDice(3)
    defenseResults[0].SetResult(dice.BLANK)
    defenseResults[1].SetResult(dice.EVADE)
    defenseResults[2].SetResult(dice.FOCUS)
    scenario := Scenario {
	AttackResults: &attackResults,
	NumDefenderFocus: 2,
	DefenseResults: &defenseResults,
    }
    scenario.DefenderModifiesDefenseDice = append(scenario.DefenderModifiesDefenseDice, Modifications["Defensive Focus"])

    if scenario.DefenseResults.Evades() != 1 {
	t.Fatalf("Should have only one evade")
    }

    scenario.Run()

    if scenario.DefenseResults.Evades() != 2 {
	t.Errorf("Should have two evades")
    }

    if scenario.NumDefenderFocus != 1 {
	t.Errorf("Should have one focus left")
    }
}

func TestUseEvadeToken_withoutEvadeToken(t *testing.T) {
    attackResults := dice.RollAttackDice(1)
    defenseResults := dice.RollDefenseDice(1)
    defenseResults[0].SetResult(dice.BLANK)
    scenario := Scenario {
	AttackResults: &attackResults,
	DefenseResults: &defenseResults,
    }
    scenario.DefenderModifiesDefenseDice = append(scenario.DefenderModifiesDefenseDice, Modifications["Use Evade Token"])

    if scenario.DefenseResults.Evades() != 0 {
	t.Fatalf("Should have no evades")
    }

    scenario.Run()

    if scenario.DefenseResults.Evades() != 0 {
	t.Errorf("Should still have no evades")
    }
}

func TestUseEvadeToken_withEvadeTokens(t *testing.T) {
    attackResults := dice.RollAttackDice(1)
    defenseResults := dice.RollDefenseDice(1)
    defenseResults[0].SetResult(dice.BLANK)
    scenario := Scenario {
	AttackResults: &attackResults,
	DefenseResults: &defenseResults,
	NumDefenderEvade: 2,
    }
    scenario.DefenderModifiesDefenseDice = append(scenario.DefenderModifiesDefenseDice, Modifications["Use Evade Token"])

    if scenario.DefenseResults.Evades() != 0 {
	t.Fatalf("Should have no evades")
    }

    scenario.Run()

    if scenario.DefenseResults.Evades() != 1 {
	t.Errorf("Should have one evade")
    }
    if scenario.NumDefenderEvade != 1 {
	t.Errorf("Should have one evade token left")
    }

    if scenario.DefenseResults.Focuses() != 0 {
	t.Errorf("Should have no focuses")
    }
    scenario.DefenseResults.ConvertAll(dice.EVADE, dice.FOCUS)
    if scenario.DefenseResults.Focuses() != 0 {
	t.Errorf("Should still have no focuses")
    }
}

func TestMarksmanship_withFocusTokens(t *testing.T) {
    attackResults := dice.RollAttackDice(4)
    defenseResults := dice.RollDefenseDice(1)
    attackResults[0].SetResult(dice.BLANK)
    attackResults[1].SetResult(dice.HIT)
    attackResults[2].SetResult(dice.FOCUS)
    attackResults[3].SetResult(dice.FOCUS)
    scenario := Scenario {
	AttackResults: &attackResults,
	DefenseResults: &defenseResults,
	NumAttackerFocus: 3,
    }
    scenario.AttackerModifiesAttackDice = append(scenario.AttackerModifiesAttackDice, Modifications["Marksmanship"])

    if scenario.AttackResults.Crits() != 0 {
	t.Fatalf("Should have no crits")
    }

    scenario.Run()

    if scenario.AttackResults.Crits() != 1 {
	t.Fatalf("Should have one crit")
    }

    if scenario.AttackResults.Focuses() != 1 {
	t.Fatalf("Should have one focus result left")
    }

    if scenario.NumAttackerFocus != 2 {
	t.Fatalf("Should have two focus tokens left")
    }
}

func TestMarksmanship_withoutFocusTokens(t *testing.T) {
    attackResults := dice.RollAttackDice(4)
    defenseResults := dice.RollDefenseDice(1)
    attackResults[0].SetResult(dice.BLANK)
    attackResults[1].SetResult(dice.HIT)
    attackResults[2].SetResult(dice.FOCUS)
    attackResults[3].SetResult(dice.FOCUS)
    scenario := Scenario {
	AttackResults: &attackResults,
	DefenseResults: &defenseResults,
    }
    scenario.AttackerModifiesAttackDice = append(scenario.AttackerModifiesAttackDice, Modifications["Marksmanship"])

    if scenario.AttackResults.Crits() != 0 {
	t.Fatalf("Should have no crits")
    }

    scenario.Run()

    if scenario.AttackResults.Crits() != 0 {
	t.Fatalf("Should still have no crits")
    }

    if scenario.AttackResults.Focuses() != 2 {
	t.Fatalf("Should still have two focus results left")
    }
}

func TestC3PO_guessNoneAndBeWrong(t *testing.T) {
    attackResults := dice.RollAttackDice(0)
    defenseResults := dice.RollDefenseDice(4)
    defenseResults[0].SetResult(dice.EVADE)
    defenseResults[1].SetResult(dice.EVADE)
    defenseResults[2].SetResult(dice.FOCUS)
    defenseResults[3].SetResult(dice.BLANK)
    scenario := Scenario {
	AttackResults: &attackResults,
	DefenseResults: &defenseResults,
    }
    scenario.DefenderModifiesDefenseDice = append(scenario.DefenderModifiesDefenseDice, Modifications["C-3PO (guess 0)"])

    scenario.Run()

    if scenario.DefenseResults.Evades() != 2 {
	t.Errorf("Should still have two evades")
    }

}

func TestC3PO_guessNoneAndBeRight(t *testing.T) {
    attackResults := dice.RollAttackDice(0)
    defenseResults := dice.RollDefenseDice(4)
    defenseResults[0].SetResult(dice.BLANK)
    defenseResults[1].SetResult(dice.FOCUS)
    defenseResults[2].SetResult(dice.FOCUS)
    defenseResults[3].SetResult(dice.BLANK)
    scenario := Scenario {
	AttackResults: &attackResults,
	DefenseResults: &defenseResults,
    }
    scenario.DefenderModifiesDefenseDice = append(scenario.DefenderModifiesDefenseDice, Modifications["C-3PO (guess 0)"])

    scenario.Run()

    if scenario.DefenseResults.Evades() != 1 {
	t.Errorf("Should now have one evade")
    }

}

func TestC3PO_guessOneAndBeWrong(t *testing.T) {
    attackResults := dice.RollAttackDice(0)
    defenseResults := dice.RollDefenseDice(4)
    defenseResults[0].SetResult(dice.EVADE)
    defenseResults[1].SetResult(dice.EVADE)
    defenseResults[2].SetResult(dice.FOCUS)
    defenseResults[3].SetResult(dice.BLANK)
    scenario := Scenario {
	AttackResults: &attackResults,
	DefenseResults: &defenseResults,
    }
    scenario.DefenderModifiesDefenseDice = append(scenario.DefenderModifiesDefenseDice, Modifications["C-3PO (guess 1)"])

    scenario.Run()

    if scenario.DefenseResults.Evades() != 2 {
	t.Errorf("Should still have two evades")
    }

}

func TestC3PO_guessOneAndBeRight(t *testing.T) {
    attackResults := dice.RollAttackDice(0)
    defenseResults := dice.RollDefenseDice(4)
    defenseResults[0].SetResult(dice.BLANK)
    defenseResults[1].SetResult(dice.FOCUS)
    defenseResults[2].SetResult(dice.EVADE)
    defenseResults[3].SetResult(dice.BLANK)
    scenario := Scenario {
	AttackResults: &attackResults,
	DefenseResults: &defenseResults,
    }
    scenario.DefenderModifiesDefenseDice = append(scenario.DefenderModifiesDefenseDice, Modifications["C-3PO (guess 1)"])

    scenario.Run()

    if scenario.DefenseResults.Evades() != 2 {
	t.Errorf("Should now have two evades")
    }

}

func TestC3PO_guessTwoAndBeWrong(t *testing.T) {
    attackResults := dice.RollAttackDice(0)
    defenseResults := dice.RollDefenseDice(4)
    defenseResults[0].SetResult(dice.EVADE)
    defenseResults[1].SetResult(dice.BLANK)
    defenseResults[2].SetResult(dice.FOCUS)
    defenseResults[3].SetResult(dice.BLANK)
    scenario := Scenario {
	AttackResults: &attackResults,
	DefenseResults: &defenseResults,
    }
    scenario.DefenderModifiesDefenseDice = append(scenario.DefenderModifiesDefenseDice, Modifications["C-3PO (guess 2)"])

    scenario.Run()

    if scenario.DefenseResults.Evades() != 1 {
	t.Errorf("Should still have one evade")
    }

}

func TestC3PO_guessTwoAndBeRight(t *testing.T) {
    attackResults := dice.RollAttackDice(0)
    defenseResults := dice.RollDefenseDice(4)
    defenseResults[0].SetResult(dice.BLANK)
    defenseResults[1].SetResult(dice.EVADE)
    defenseResults[2].SetResult(dice.FOCUS)
    defenseResults[3].SetResult(dice.EVADE)
    scenario := Scenario {
	AttackResults: &attackResults,
	DefenseResults: &defenseResults,
    }
    scenario.DefenderModifiesDefenseDice = append(scenario.DefenderModifiesDefenseDice, Modifications["C-3PO (guess 2)"])

    scenario.Run()

    if scenario.DefenseResults.Evades() != 3 {
	t.Errorf("Should now have three evades")
    }

}

func TestChiraneau_withFocusResult (t *testing.T) {
    attackResults := dice.RollAttackDice(4)
    defenseResults := dice.RollDefenseDice(1)
    attackResults[0].SetResult(dice.BLANK)
    attackResults[1].SetResult(dice.HIT)
    attackResults[2].SetResult(dice.FOCUS)
    attackResults[3].SetResult(dice.FOCUS)
    scenario := Scenario {
	AttackResults: &attackResults,
	DefenseResults: &defenseResults,
    }
    scenario.AttackerModifiesAttackDice = append(scenario.AttackerModifiesAttackDice, Modifications["Chiraneau"])

    scenario.Run()

    if scenario.AttackResults.Crits() != 1 {
	t.Errorf("Should now have one crit")
    }

    if scenario.AttackResults.Focuses() != 1 {
	t.Errorf("Should now have one focus result left")
    }
}

func TestChiraneau_withoutFocusResult (t *testing.T) {
    attackResults := dice.RollAttackDice(4)
    defenseResults := dice.RollDefenseDice(1)
    attackResults[0].SetResult(dice.BLANK)
    attackResults[1].SetResult(dice.HIT)
    attackResults[2].SetResult(dice.BLANK)
    attackResults[3].SetResult(dice.HIT)
    scenario := Scenario {
	AttackResults: &attackResults,
	DefenseResults: &defenseResults,
    }
    scenario.AttackerModifiesAttackDice = append(scenario.AttackerModifiesAttackDice, Modifications["Chiraneau"])

    scenario.Run()

    if scenario.AttackResults.Crits() != 0 {
	t.Errorf("Should still have no crits")
    }

    if scenario.AttackResults.Focuses() != 0 {
	t.Errorf("Should still have no focus results")
    }
}

func TestHeavyLaserCannon (t *testing.T) {
    attackResults := dice.RollAttackDice(4)
    defenseResults := dice.RollDefenseDice(1)
    attackResults[0].SetResult(dice.BLANK)
    attackResults[1].SetResult(dice.HIT)
    attackResults[2].SetResult(dice.CRIT)
    attackResults[3].SetResult(dice.HIT)
    scenario := Scenario {
	AttackResults: &attackResults,
	DefenseResults: &defenseResults,
    }
    scenario.AttackerModifiesAttackDice = append(scenario.AttackerModifiesAttackDice, Modifications["Heavy Laser Cannon"])

    scenario.Run()

    if scenario.AttackResults.Crits() != 0 {
	t.Errorf("Should have no crits")
    }

    if scenario.AttackResults.Hits() != 3 {
	t.Errorf("Should have three hits")
    }
}
