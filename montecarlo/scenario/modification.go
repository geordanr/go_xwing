package scenario

import (
    "github.com/geordanr/xwing/dice"
    "github.com/geordanr/xwing/dice/filters"
)

type Modification interface {
    Modify(*Scenario) *Scenario
}

// Attacker modifications

type offensiveFocus struct {}
func (offensiveFocus) Modify(scenario *Scenario) *Scenario {
    if scenario.NumAttackerFocus > 0 && scenario.AttackResults.Focuses() > 0 {
	scenario.NumAttackerFocus--
	scenario.AttackResults.ConvertAll(dice.FOCUS, dice.HIT)
    }
    return scenario
}

type targetLock struct {}
func (targetLock) Modify(scenario *Scenario) *Scenario {
    if scenario.NumAttackerFocus > 0 {
	scenario.AttackResults.Reroll(filters.Blanks)
    } else {
	scenario.AttackResults.Reroll(filters.BlanksAndFocuses)
    }
    return scenario
}

type howlrunner struct {}
func (howlrunner) Modify(scenario *Scenario) *Scenario {
    if scenario.NumAttackerFocus > 0 {
	scenario.AttackResults.RerollUpto(1, filters.Blanks)
    } else {
	scenario.AttackResults.RerollUpto(1, filters.BlanksAndFocuses)
    }
    return scenario
}

// Defender modifications

type defensiveFocus struct {}
func (defensiveFocus) Modify(scenario *Scenario) *Scenario {
    if scenario.NumDefenderFocus > 0 && scenario.DefenseResults.Focuses() > 0 {
	scenario.NumDefenderFocus--
	scenario.DefenseResults.ConvertAll(dice.FOCUS, dice.EVADE)
    }
    return scenario
}

type useEvadeToken struct {}
func (useEvadeToken) Modify(scenario *Scenario) *Scenario {
    if scenario.NumDefenderEvade > 0 {
	scenario.NumDefenderEvade--

	evadeDie := new(dice.DefenseDie)
	evadeDie.SetResult(dice.EVADE)
	evadeDie.Lock()
	results := *scenario.DefenseResults
	results = append(results, evadeDie)
	scenario.DefenseResults = &results
    }

    return scenario
}

// Modifications map
var Modifications map[string]Modification = map[string]Modification{
    "Offensive Focus": new(offensiveFocus),
    "Target Lock": new(targetLock),
    "Howlrunner": new(howlrunner),
    "Defensive Focus": new(defensiveFocus),
    "Use Evade Token": new(useEvadeToken),
}
