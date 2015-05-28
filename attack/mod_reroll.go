package attack

import (
    "fmt"
    "github.com/geordanr/go_xwing/dice/filters"
    "github.com/geordanr/go_xwing/dice"
)

type targetLock struct {}
func (targetLock) Modify(atk *Attack) *Attack {
    if atk.Attacker.FocusTokens > 0 {
	atk.AttackResults.Reroll(filters.Blanks)
    } else {
	atk.AttackResults.Reroll(filters.BlanksAndFocuses)
    }
    return atk
}
func (targetLock) String() string { return "Target Lock" }
func (targetLock) ModifiesAttackResults() bool { return true }
func (targetLock) ModifiesDefenseResults() bool { return false }

type offensiveReroll struct {
    numToReroll uint
    name string
}
func (o offensiveReroll) Modify(atk *Attack) *Attack {
    if atk.Attacker.FocusTokens > 0 {
	atk.AttackResults.RerollUpto(o.numToReroll, filters.Blanks)
    } else {
	atk.AttackResults.RerollUpto(o.numToReroll, filters.BlanksAndFocuses)
    }
    return atk
}
func (o offensiveReroll) String() string {
    return fmt.Sprintf("%s", o.name)
}
func (offensiveReroll) ModifiesAttackResults() bool { return true }
func (offensiveReroll) ModifiesDefenseResults() bool { return false }

type lukeSkywalkerPilot struct {}

func (lukeSkywalkerPilot) Modify( atk *Attack) *Attack {
    if atk.DefenseResults.Focuses() > 0 {
        atk.DefenseResults.ConvertUpto(1, dice.FOCUS, dice.EVADE)
    }
    return atk
}

func (lukeSkywalkerPilot) String() string { return "Luke Skywalker" }
func (lukeSkywalkerPilot) ModifiesAttackResults() bool { return false }
func (lukeSkywalkerPilot) ModifiesDefenseResults() bool { return true }



type hanSolo struct {}
func (hanSolo) Modify(atk *Attack) *Attack {
    if atk.AttackResults.Hits() + atk.AttackResults.Crits() < uint8(len(*atk.AttackResults)) {
	atk.AttackResults.Reroll(filters.Everything)
    }
    return atk
}
func (hanSolo) String() string { return "Han Solo" }
func (hanSolo) ModifiesAttackResults() bool { return true }
func (hanSolo) ModifiesDefenseResults() bool { return false }

type defensiveReroll struct {
    numToReroll uint
    name string
}
func (o defensiveReroll) Modify(atk *Attack) *Attack {
    if atk.Defender.FocusTokens > 0 {
	atk.DefenseResults.RerollUpto(o.numToReroll, filters.Blanks)
    } else {
	atk.DefenseResults.RerollUpto(o.numToReroll, filters.BlanksAndFocuses)
    }
    return atk
}
func (o defensiveReroll) String() string {
    return fmt.Sprintf("%s", o.name)
}
func (defensiveReroll) ModifiesAttackResults() bool { return false }
func (defensiveReroll) ModifiesDefenseResults() bool { return true }
