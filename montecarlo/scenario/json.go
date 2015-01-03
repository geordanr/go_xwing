package scenario

import (
    "encoding/json"
    "github.com/geordanr/xwing/dice"
)

type scenarioJSON struct {
    NumAttackDice uint
    NumAttackerFocus uint
    NumDefenseDice uint
    NumDefenderFocus uint
    NumDefenderEvade uint
    DefenderModifiesAttackDice []string
    AttackerModifiesAttackDice []string
    AttackerModifiesDefenseDice []string
    DefenderModifiesDefenseDice []string
}

func ScenarioFromJSON(b []byte) (*Scenario, error) {
    var s scenarioJSON
    err := json.Unmarshal(b, &s)

    if err != nil {
	return nil, err
    } else {
	attackResults := dice.RollAttackDice(uint8(s.NumAttackDice))
	defenseResults := dice.RollDefenseDice(uint8(s.NumDefenseDice))

	defenderModifiesAttackDice := make([]Modification, len(s.DefenderModifiesAttackDice))
	for i := range(s.DefenderModifiesAttackDice) {
	    defenderModifiesAttackDice[i] = Modifications[s.DefenderModifiesAttackDice[i]]
	}

	attackerModifiesAttackDice := make([]Modification, len(s.AttackerModifiesAttackDice))
	for i := range(s.AttackerModifiesAttackDice) {
	    attackerModifiesAttackDice[i] = Modifications[s.AttackerModifiesAttackDice[i]]
	}

	attackerModifiesDefenseDice := make([]Modification, len(s.AttackerModifiesDefenseDice))
	for i := range(s.AttackerModifiesDefenseDice) {
	    attackerModifiesDefenseDice[i] = Modifications[s.AttackerModifiesDefenseDice[i]]
	}

	defenderModifiesDefenseDice := make([]Modification, len(s.DefenderModifiesDefenseDice))
	for i := range(s.DefenderModifiesDefenseDice) {
	    defenderModifiesDefenseDice[i] = Modifications[s.DefenderModifiesDefenseDice[i]]
	}

	scenario := Scenario {
	    AttackResults: &attackResults,
	    DefenseResults: &defenseResults,
	    NumAttackerFocus: uint(s.NumAttackerFocus),
	    NumDefenderFocus: uint(s.NumDefenderFocus),
	    NumDefenderEvade: uint(s.NumDefenderEvade),
	    DefenderModifiesAttackDice: defenderModifiesAttackDice,
	    AttackerModifiesAttackDice: attackerModifiesAttackDice,
	    AttackerModifiesDefenseDice: attackerModifiesDefenseDice,
	    DefenderModifiesDefenseDice: defenderModifiesDefenseDice,
	}
	return &scenario, nil
    }
}
