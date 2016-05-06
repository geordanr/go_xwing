/*
Package modification implements modification functions used when executing
the steps of an attack.

A modification has a ModifyState method which takes two arguments, the
game state and the ship to affect.  Modifications also have an actor
which is used to signify what role the ship is playing (usually the
attacker or the defender).

Modifications may access and modify any part of the state, including
attack results, the attack queue, and even what modifications will be
applied to downstream attack steps.
*/
package modification

import (
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/interfaces"
)

// All contains a mapping from human-readable strings to Modification factory functions.
// In general, if a modification's name is the same as an attack step, it is the default handler for that step.
var All = map[string]func() interfaces.Modification{
	"Accuracy Corrector":   func() interfaces.Modification { return new(AccuracyCorrector) },               // accuracycorrector.go
	"C-3PO (guess 0)":      func() interfaces.Modification { return new(C3PO) },                            // c3po.go
	"C-3PO (guess 1)":      func() interfaces.Modification { return &C3PO{guess: 1} },                      // c3po.go
	"C-3PO (guess 2)":      func() interfaces.Modification { return &C3PO{guess: 2} },                      // c3po.go
	"C-3PO (guess 3)":      func() interfaces.Modification { return &C3PO{guess: 3} },                      // c3po.go
	"Cannot Attack Again":  func() interfaces.Modification { return new(CannotAttackAgain) },               // cannotattackagain.go
	"Compare Results":      func() interfaces.Modification { return new(CompareResults) },                  // compareresults.go
	"Crack Shot":           func() interfaces.Modification { return new(CrackShot) },                       // crackshot.go
	"Deal Damage":          func() interfaces.Modification { return new(DealDamage) },                      // dealdamage.go
	"Declare Target":       func() interfaces.Modification { return new(DeclareTarget) },                   // declaretarget.go
	"Spend Evade Token":    func() interfaces.Modification { return new(SpendEvade) },                      // evade.go
	"Spend Focus Token":    func() interfaces.Modification { return new(SpendFocus) },                      // focus.go
	"Gunner":               func() interfaces.Modification { return new(Gunner) },                          // gunner.go
	"Heavy Laser Cannon":   func() interfaces.Modification { return new(HeavyLaserCannon) },                // hlc.go
	"Roll Attack Dice":     func() interfaces.Modification { return &RollDice{actor: constants.ATTACKER} }, // roll.go
	"Roll Defense Dice":    func() interfaces.Modification { return &RollDice{actor: constants.DEFENDER} }, // roll.go
	"Perform Attack Twice": func() interfaces.Modification { return new(PerformAttackTwice) },              // performattacktwice.go
	"Spend Target Lock":    func() interfaces.Modification { return new(SpendTargetLock) },                 // targetlock.go
	"Twin Laser Turret":    func() interfaces.Modification { return new(TwinLaserTurret) },                 // tlt.go
}
