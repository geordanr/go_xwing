package attack

import (
    "fmt"
    "github.com/geordanr/go_xwing/dice"
)

type c3po struct {
    Guess uint8
}

func (threepio c3po) Modify(atk *Attack) *Attack {
    if atk.DefenseResults.Evades() == threepio.Guess {
	evadeDie := new(dice.DefenseDie)
	evadeDie.SetResult(dice.EVADE)
	evadeDie.Lock()
	results := *atk.DefenseResults
	results = append(results, evadeDie)
	atk.DefenseResults = &results
    }
    return atk
}
func (threepio c3po) String() string {
    return fmt.Sprintf("C-3PO (guess %d)", threepio.Guess)
}
func (c3po) ModifiesAttackResults() bool { return false }
func (c3po) ModifiesDefenseResults() bool { return true }
