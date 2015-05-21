package attack

import "github.com/geordanr/go_xwing/dice"

type Modification interface {
    Modify(*Attack) *Attack
    String() string
    ModifiesAttackResults() bool
    ModifiesDefenseResults() bool
}

var Modifications map[string]Modification = map[string]Modification{
    "Offensive Focus": new(offensiveFocus),
    "Target Lock": new(targetLock),
    "Howlrunner": offensiveReroll{name: "Howlrunner", numToReroll: 1},
    "Predator (low PS)": offensiveReroll{name: "Predator (low PS)", numToReroll: 2},
    "Predator (high PS)": offensiveReroll{name: "Predator (high PS)", numToReroll: 1},
    "Marksmanship": new(marksmanship),
    "Chiraneau": new(chiraneau),
    "Han Solo": new(hanSolo),
    "Heavy Laser Cannon": new(heavyLaserCannon),
    "Accuracy Corrector": new(accuracyCorrector),

    "Defensive Focus": new(defensiveFocus),
    "Use Evade Token": new(useEvadeToken),
    "C-3PO (guess 0)": c3po{Guess: 0},
    "C-3PO (guess 1)": c3po{Guess: 1},
    "C-3PO (guess 2)": c3po{Guess: 2},
    "C-3PO (guess 3)": c3po{Guess: 3},
}

// For testing

// AttackDiceSetter is used to force specific attack dice results.  Should be the first attack modification.
type AttackDiceSetter struct {
    desiredResults []dice.Result
}
func (mod AttackDiceSetter) Modify(atk *Attack) *Attack {
    results := make(dice.Results, len(mod.desiredResults))
    for i, result := range(mod.desiredResults) {
	results[i] = new(dice.AttackDie)
	results[i].SetResult(result)
    }
    atk.AttackResults = &results
    return atk
}
func (AttackDiceSetter) String() string { return "Attack Die Setter" }
func (AttackDiceSetter) ModifiesAttackResults() bool { return true }
func (AttackDiceSetter) ModifiesDefenseResults() bool { return false }

// DefenseDiceSetter is used to force specific attack dice results.  Should be the first defense modification.
type DefenseDiceSetter struct {
    desiredResults []dice.Result
}
func (mod DefenseDiceSetter) Modify(atk *Attack) *Attack {
    results := make(dice.Results, len(mod.desiredResults))
    for i, result := range(mod.desiredResults) {
	results[i] = new(dice.DefenseDie)
	results[i].SetResult(result)
    }
    atk.DefenseResults = &results
    return atk
}
func (DefenseDiceSetter) String() string { return "Defense Die Setter" }
func (DefenseDiceSetter) ModifiesAttackResults() bool { return false }
func (DefenseDiceSetter) ModifiesDefenseResults() bool { return true }