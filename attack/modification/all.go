/*
Package modification implements modification functions used when executing
the steps of an attack.

A modification has a ModifyState method which takes two arguments, the
game state and the ship to affect.  Modifications also have an actor
which is used to signify what role the ship is playing (usually the
attacker or the defender).
*/
package modification

import (
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/interfaces"
)

// All contains a mapping from human-readable strings to Modification factory functions.
var All = map[string]func() interfaces.Modification{
	"C-3PO (guess 0)":      func() interfaces.Modification { return new(C3PO) },                            // c3po.go
	"C-3PO (guess 1)":      func() interfaces.Modification { return &C3PO{guess: 1} },                      // c3po.go
	"C-3PO (guess 2)":      func() interfaces.Modification { return &C3PO{guess: 2} },                      // c3po.go
	"C-3PO (guess 3)":      func() interfaces.Modification { return &C3PO{guess: 3} },                      // c3po.go
	"Cannot Attack Again":  func() interfaces.Modification { return new(CannotAttackAgain) },               // cannotattackagain.go
	"Compare Results":      func() interfaces.Modification { return new(CompareResults) },                  // compareresults.go
	"Declare Target":       func() interfaces.Modification { return new(DeclareTarget) },                   // declaretarget.go
	"Spend Evade Token":    func() interfaces.Modification { return new(SpendEvade) },                      // evade.go
	"Spend Focus Token":    func() interfaces.Modification { return new(SpendFocus) },                      // focus.go
	"Gunner":               func() interfaces.Modification { return new(Gunner) },                          // gunner.go
	"Roll Attack Dice":     func() interfaces.Modification { return &RollDice{actor: constants.ATTACKER} }, // roll.go
	"Roll Defense Dice":    func() interfaces.Modification { return &RollDice{actor: constants.DEFENDER} }, // roll.go
	"Perform Attack Twice": func() interfaces.Modification { return new(PerformAttackTwice) },              // performattacktwice.go
	"Suffer Damage":        func() interfaces.Modification { return new(SufferDamage) },                    // sufferdamage.go
	"Spend Target Lock":    func() interfaces.Modification { return new(SpendTargetLock) },                 // targetlock.go
	"Twin Laser Turret":    func() interfaces.Modification { return new(TwinLaserTurret) },                 // tlt.go
}
