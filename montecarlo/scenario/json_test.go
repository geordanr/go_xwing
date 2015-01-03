package scenario

import (
    "io/ioutil"
    "path/filepath"
    "runtime"
    "testing"
)

func TestScenarioFromJSON(t *testing.T) {
    _, thisfile, _, _ := runtime.Caller(0)
    thisdir := filepath.Dir(thisfile)
    data, err := ioutil.ReadFile(filepath.Join(thisdir, "example.json"))
    if err != nil {
	t.Fatal(err)
    }

    s, err := ScenarioFromJSON(data)
    if err != nil {
	t.Fatal(err)
    }

    if s == nil {
	t.Fatalf("Scenario is nil")
    }

    var expected, got uint

    expected = 4
    got = uint(len(*s.AttackResults))
    if got != expected {
	t.Errorf("Expected %d attack dice, got %d\n", expected, got)
    }

    expected = 3
    got = uint(len(*s.DefenseResults))
    if got != expected {
	t.Errorf("Expected %d defense dice, got %d\n", expected, got)
    }

    expected = 1
    got = s.NumAttackerFocus
    if got != expected {
	t.Errorf("Expected %d attacker focus, got %d\n", expected, got)
    }

    expected = 2
    got = s.NumDefenderFocus
    if got != expected {
	t.Errorf("Expected %d defender focus, got %d\n", expected, got)
    }

    expected = 0
    got = s.NumDefenderEvade
    if got != expected {
	t.Errorf("Expected %d evade, got %d\n", expected, got)
    }

    if len(s.DefenderModifiesAttackDice) != 0 {
	t.Errorf("Unexpected %d entries in defender modifies attack dice step", len(s.DefenderModifiesAttackDice))
    }

    if len(s.AttackerModifiesDefenseDice) != 0 {
	t.Errorf("Unexpected %d entries in attacker modifies defense dice step", len(s.AttackerModifiesDefenseDice))
    }

    if len(s.AttackerModifiesAttackDice) != 2 {
	t.Errorf("Unexpected %d entries in attacker modifies attack dice step", len(s.AttackerModifiesAttackDice))
    }

    if s.AttackerModifiesAttackDice[0] != Modifications["Howlrunner"] {
	t.Errorf("Expected Howlrunner as first attacker modification")
    }

    if s.AttackerModifiesAttackDice[1] != Modifications["Offensive Focus"] {
	t.Errorf("Expected Offensive Focus as second attacker modification")
    }

    if len(s.DefenderModifiesDefenseDice) != 1 {
	t.Errorf("Unexpected %d entries in defender modifies defense dice step", len(s.DefenderModifiesDefenseDice))
    }

    if s.DefenderModifiesDefenseDice[0] != Modifications["Defensive Focus"] {
	t.Errorf("Expected Defensive Focus as defender modification")
    }

}
