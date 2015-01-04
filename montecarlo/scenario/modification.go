package scenario

import (
    "github.com/geordanr/go_xwing/dice"
    "github.com/geordanr/go_xwing/dice/filters"
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

type marksmanship struct {}
func (marksmanship) Modify(scenario *Scenario) *Scenario {
    if scenario.NumAttackerFocus > 0 {
	scenario.NumAttackerFocus--
	scenario.AttackResults.ConvertUpto(1, dice.FOCUS, dice.CRIT)
    }
    return scenario
}

type chiraneau struct {}
func (chiraneau) Modify(scenario *Scenario) *Scenario {
    scenario.AttackResults.ConvertUpto(1, dice.FOCUS, dice.CRIT)
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

type c3po struct {
    Guess uint8
}
func (threepio c3po) Modify(scenario *Scenario) *Scenario {
    if scenario.DefenseResults.Evades() == threepio.Guess {
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
    "Marksmanship": new(marksmanship),
    "C-3PO (guess 0)": c3po{Guess: 0},
    "C-3PO (guess 1)": c3po{Guess: 1},
    "C-3PO (guess 2)": c3po{Guess: 2},
    "C-3PO (guess 3)": c3po{Guess: 3},
    "Chiraneau": new(chiraneau),
}
