package attack

import (
    "github.com/geordanr/go_xwing/dice"
)

type landoCrew struct {}
func (landoCrew) Modify(atk *Attack) *Attack {
    results := *atk.DefenseResults
    landoResults := dice.RollDefenseDice(2)

    var i uint8
    for i = 0; i < landoResults.Evades(); i++ {
	evadeDie := new(dice.DefenseDie)
	evadeDie.SetResult(dice.EVADE)
	evadeDie.Lock()
	results = append(results, evadeDie)
    }

    for i = 0; i < landoResults.Focuses(); i++ {
	evadeDie := new(dice.DefenseDie)
	evadeDie.SetResult(dice.FOCUS)
	evadeDie.Lock()
	results = append(results, evadeDie)
    }

    return atk
}
func (landoCrew) String() string { return "Lando (Crew)" }
func (landoCrew) ModifiesAttackResults() bool { return false }
func (landoCrew) ModifiesDefenseResults() bool { return true }
