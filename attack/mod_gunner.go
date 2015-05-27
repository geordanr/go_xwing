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

type lukeSkywalkerCrew struct {}
func (lukeSkywalkerCrew) Modify(atk *Attack) *Attack {
    if atk.IsGunnerAttack {
	if atk.AttackResults.Focuses() > 0 {
	    atk.AttackResults.ConvertUpto(1, dice.FOCUS, dice.HIT)
	}
    } else {
	atk.UseGunner = true
    }
    return atk
}
func (lukeSkywalkerCrew) String() string { return "Luke Skywalker (Crew)" }
func (lukeSkywalkerCrew) ModifiesAttackResults() bool { return true }
func (lukeSkywalkerCrew) ModifiesDefenseResults() bool { return false }
