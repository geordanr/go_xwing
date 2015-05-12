package attack

import (
    "github.com/geordanr/go_xwing/dice"
)

type offensiveFocus struct {}
func (offensiveFocus) Modify(atk *Attack) *Attack {
    if atk.AttackResults.Focuses() > 0 && atk.Attacker.SpendFocus() {
	atk.AttackResults.ConvertAll(dice.FOCUS, dice.HIT)
    }
    return atk
}
func (offensiveFocus) String() string { return "Offensive Focus" }
func (offensiveFocus) ModifiesAttackResults() bool { return true }
func (offensiveFocus) ModifiesDefenseResults() bool { return false }

type defensiveFocus struct {}
func (defensiveFocus) Modify(atk *Attack) *Attack {
    if atk.DefenseResults.Focuses() > 0 && atk.Defender.SpendFocus() {
	atk.DefenseResults.ConvertAll(dice.FOCUS, dice.EVADE)
    }
    return atk
}
func (defensiveFocus) String() string { return "Defensive Focus" }
func (defensiveFocus) ModifiesAttackResults() bool { return false }
func (defensiveFocus) ModifiesDefenseResults() bool { return true }

type marksmanship struct {}
func (marksmanship) Modify(atk *Attack) *Attack {
    if atk.AttackResults.Focuses() > 0 && atk.Attacker.SpendFocus() {
	atk.AttackResults.ConvertUpto(1, dice.FOCUS, dice.CRIT)
    }
    return atk
}
func (marksmanship) String() string { return "Marksmanship" }
func (marksmanship) ModifiesAttackResults() bool { return true }
func (marksmanship) ModifiesDefenseResults() bool { return false }

type chiraneau struct {}
func (chiraneau) Modify(atk *Attack) *Attack {
    atk.AttackResults.ConvertUpto(1, dice.FOCUS, dice.CRIT)
    return atk
}
func (chiraneau) String() string { return "Chiraneau" }
func (chiraneau) ModifiesAttackResults() bool { return true }
func (chiraneau) ModifiesDefenseResults() bool { return false }
