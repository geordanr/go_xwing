package attack

import (
	"github.com/geordanr/go_xwing/dice"
)

type useEvadeToken struct{}

func (useEvadeToken) Modify(atk *Attack) *Attack {
	if atk.Defender.EvadeTokens > 0 && (atk.AttackResults.Hits()+atk.AttackResults.Crits()) > (atk.DefenseResults.Evades()) {
		atk.Defender.EvadeTokens--
		evadeDie := new(dice.DefenseDie)
		evadeDie.SetResult(dice.EVADE)
		evadeDie.Lock()
		results := *atk.DefenseResults
		results = append(results, evadeDie)
		atk.DefenseResults = &results
	}
	return atk
}
func (useEvadeToken) String() string               { return "Use Evade Token" }
func (useEvadeToken) ModifiesAttackResults() bool  { return false }
func (useEvadeToken) ModifiesDefenseResults() bool { return true }
