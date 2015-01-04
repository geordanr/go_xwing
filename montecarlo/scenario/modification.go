package scenario

import (
    "fmt"
    "github.com/geordanr/go_xwing/dice"
    "github.com/geordanr/go_xwing/dice/filters"
)

type Modification interface {
    Modify(*Scenario) *Scenario
    String() string
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
func (offensiveFocus) String() string { return "Offensive Focus" }

type targetLock struct {}
func (targetLock) Modify(scenario *Scenario) *Scenario {
    if scenario.NumAttackerFocus > 0 {
	scenario.AttackResults.Reroll(filters.Blanks)
    } else {
	scenario.AttackResults.Reroll(filters.BlanksAndFocuses)
    }
    return scenario
}
func (targetLock) String() string { return "Target Lock" }

type offensiveReroll struct {
    numToReroll uint
    name string
}
func (o offensiveReroll) Modify(scenario *Scenario) *Scenario {
    if scenario.NumAttackerFocus > 0 {
	scenario.AttackResults.RerollUpto(o.numToReroll, filters.Blanks)
    } else {
	scenario.AttackResults.RerollUpto(o.numToReroll, filters.BlanksAndFocuses)
    }
    return scenario
}
func (o offensiveReroll) String() string {
    return fmt.Sprintf("%s", o.name)
}

type marksmanship struct {}
func (marksmanship) Modify(scenario *Scenario) *Scenario {
    if scenario.NumAttackerFocus > 0 {
	scenario.NumAttackerFocus--
	scenario.AttackResults.ConvertUpto(1, dice.FOCUS, dice.CRIT)
    }
    return scenario
}
func (marksmanship) String() string { return "Marksmanship" }

type chiraneau struct {}
func (chiraneau) Modify(scenario *Scenario) *Scenario {
    scenario.AttackResults.ConvertUpto(1, dice.FOCUS, dice.CRIT)
    return scenario
}
func (chiraneau) String() string { return "Chiraneau" }

type heavyLaserCannon struct {}
func (heavyLaserCannon) Modify(scenario *Scenario) *Scenario {
    scenario.AttackResults.ConvertAll(dice.CRIT, dice.HIT)
    return scenario
}
func (heavyLaserCannon) String() string { return "Heavy Laser Cannon" }

// Defender modifications

type defensiveFocus struct {}
func (defensiveFocus) Modify(scenario *Scenario) *Scenario {
    if scenario.NumDefenderFocus > 0 && scenario.DefenseResults.Focuses() > 0 {
	scenario.NumDefenderFocus--
	scenario.DefenseResults.ConvertAll(dice.FOCUS, dice.EVADE)
    }
    return scenario
}
func (defensiveFocus) String() string { return "Defensive Focus" }

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
func (useEvadeToken) String() string { return "Use Evade Token" }

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
func (threepio c3po) String() string {
    return fmt.Sprintf("C-3PO (guess %d)", threepio.Guess)
}

type defensiveReroll struct {
    numToReroll uint
    name string
}
func (o defensiveReroll) Modify(scenario *Scenario) *Scenario {
    if scenario.NumDefenderFocus > 0 {
	scenario.DefenseResults.RerollUpto(o.numToReroll, filters.Blanks)
    } else {
	scenario.DefenseResults.RerollUpto(o.numToReroll, filters.BlanksAndFocuses)
    }
    return scenario
}
func (o defensiveReroll) String() string {
    return fmt.Sprintf("%s", o.name)
}

// Modifications map
var Modifications map[string]Modification = map[string]Modification{
    "Offensive Focus": new(offensiveFocus),
    "Target Lock": new(targetLock),
    "Howlrunner": offensiveReroll{name: "Howlrunner", numToReroll: 1},
    "Predator (low PS)": offensiveReroll{name: "Predator (low PS)", numToReroll: 2},
    "Predator (high PS)": offensiveReroll{name: "Predator (high PS)", numToReroll: 1},
    "Defensive Focus": new(defensiveFocus),
    "Use Evade Token": new(useEvadeToken),
    "Marksmanship": new(marksmanship),
    "C-3PO (guess 0)": c3po{Guess: 0},
    "C-3PO (guess 1)": c3po{Guess: 1},
    "C-3PO (guess 2)": c3po{Guess: 2},
    "C-3PO (guess 3)": c3po{Guess: 3},
    "Chiraneau": new(chiraneau),
    "Heavy Laser Cannon": new(heavyLaserCannon),
}
