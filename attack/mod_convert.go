package attack

import (
    "github.com/geordanr/go_xwing/dice"
)

type heavyLaserCannon struct {}
func (heavyLaserCannon) Modify(atk *Attack) *Attack {
    atk.AttackResults.ConvertAll(dice.CRIT, dice.HIT)
    return atk
}
func (heavyLaserCannon) String() string { return "Heavy Laser Cannon" }
func (heavyLaserCannon) ModifiesAttackResults() bool { return true }
func (heavyLaserCannon) ModifiesDefenseResults() bool { return false }

type accuracyCorrector struct {}
func (accuracyCorrector) Modify(atk *Attack) *Attack {
    if atk.AttackResults.Hits() + atk.AttackResults.Crits() < 2 {
	results := make(dice.Results, 2)
	for i := range(results) {
	    results[i] = new(dice.AttackDie)
	    results[i].SetResult(dice.HIT)
	    results[i].Lock()
	}
	atk.AttackResults = &results
    }
    return atk
}
func (accuracyCorrector) String() string { return "Accuracy Corrector" }
func (accuracyCorrector) ModifiesAttackResults() bool { return true }
func (accuracyCorrector) ModifiesDefenseResults() bool { return false }


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


type autothrusters struct {}
func (autothrusters) Modify( atk *Attack) *Attack {
    if atk.DefenseResults.Blanks() > 0 {
        atk.DefenseResults.ConvertUpto(1, dice.BLANK, dice.EVADE)
    }
    return atk
}

func (autothrusters) String() string { return "Autothrusters" }
func (autothrusters) ModifiesAttackResults() bool { return false }
func (autothrusters) ModifiesDefenseResults() bool { return true }

