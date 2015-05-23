package attack

import (
    "github.com/geordanr/go_xwing/dice"
)

type gunner struct {}
func (gunner) Modify(atk *Attack) *Attack {
    atk.UseGunner = true
    return atk
}
func (gunner) String() string { return "Gunner" }
func (gunner) ModifiesAttackResults() bool { return true }
func (gunner) ModifiesDefenseResults() bool { return false }

type lukeSkywalker struct {}
func (lukeSkywalker) Modify(atk *Attack) *Attack {
    if atk.IsGunnerAttack {
	if atk.AttackResults.Focuses() > 0 {
	    atk.AttackResults.ConvertUpto(1, dice.FOCUS, dice.HIT)
	}
    } else {
	atk.UseGunner = true
    }
    return atk
}
func (lukeSkywalker) String() string { return "Luke Skywalker" }
func (lukeSkywalker) ModifiesAttackResults() bool { return true }
func (lukeSkywalker) ModifiesDefenseResults() bool { return false }
