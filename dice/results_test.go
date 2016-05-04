package dice

import "testing"

func testRollAttackDice(t *testing.T) {
	results := RollAttackDice(100)
	if results.Evades() > 0 {
		t.Errorf("Attack die rolled evade")
	}
}

func testRollDefenseDice(t *testing.T) {
	results := RollDefenseDice(100)
	if results.Hits() > 0 {
		t.Errorf("Defense die rolled hits")
	}
	if results.Crits() > 0 {
		t.Errorf("Defense die rolled crits")
	}
}

func testConvertAll(t *testing.T) {
	results := RollDefenseDice(100)
	expectedEvades := results.Evades() + results.Blanks()
	results.ConvertAll(BLANK, EVADE)
	if results.Blanks() > 0 {
		t.Errorf("Did not convert all blanks")
	}
	if results.Evades() != expectedEvades {
		t.Errorf("Did not get expected evades")
	}
}

func testConvertUpto(t *testing.T) {
	results := RollAttackDice(5)
	results[0].SetResult(BLANK)
	results[1].SetResult(HIT)
	results[2].SetResult(HIT)
	results[3].SetResult(BLANK)
	results[4].SetResult(BLANK)
	results.ConvertUpto(2, BLANK, CRIT)
	if results.Blanks() != 1 {
		t.Errorf("Converted too many blanks")
	}
	if results.Crits() != 2 {
		t.Errorf("Did not get expected crits")
	}
}
