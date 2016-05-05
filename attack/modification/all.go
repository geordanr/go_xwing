// Package modification implements modification functions used when
// executing the steps of an attack.
package modification

import (
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/interfaces"
)

// All contains a mapping from human-readable strings to Modification factory functions.
var All = map[string]func() interfaces.Modification{
	"Compare Results":      func() interfaces.Modification { return new(CompareResults) },                              // compareresults.go
	"Spend Evade Token":    func() interfaces.Modification { return new(SpendEvade) },                                  // evade.go
	"Spend Focus Token":    func() interfaces.Modification { return new(SpendFocus) },                                  // focus.go
	"Roll Attack Dice":     func() interfaces.Modification { mod := RollDice{actor: constants.ATTACKER}; return &mod }, // roll.go
	"Roll Defense Dice":    func() interfaces.Modification { mod := RollDice{actor: constants.DEFENDER}; return &mod },
	"Perform Attack Twice": func() interfaces.Modification { return new(PerformAttackTwice) }, // performattacktwice.go
	"Suffer Damage":        func() interfaces.Modification { return new(SufferDamage) },       // sufferdamage.go
	"Spend Target Lock":    func() interfaces.Modification { return new(SpendTargetLock) },    // targetlock.go
	"Twin Laser Turret":    func() interfaces.Modification { return new(TwinLaserTurret) },    // tlt.go
}
