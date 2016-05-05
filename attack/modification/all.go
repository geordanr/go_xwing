// Package modification implements modification functions used when
// executing the steps of an attack.
package modification

import (
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/interfaces"
)

var attackerSpendFocus = SpendFocus{actor: constants.ATTACKER}
var defenderSpendFocus = SpendFocus{actor: constants.DEFENDER}
var attackerRollDice = RollDice{actor: constants.ATTACKER}
var defenderRollDice = RollDice{actor: constants.DEFENDER}

// All contains a mapping from human-readable strings to Modifications.
var All = map[string]interfaces.Modification{
	// compareresults.go
	"Compare Results": &CompareResults{},
	// evade.go
	"Spend Evade": &SpendEvade{},
	// focus.go
	"Attacker Spend Focus": &attackerSpendFocus,
	"Defender Spend Focus": &defenderSpendFocus,
	// roll.go
	"Roll Attack Dice":  &attackerRollDice,
	"Roll Defense Dice": &defenderRollDice,
	// performattacktwice.go
	"Perform Attack Twice": &PerformAttackTwice{},
	// sufferdamage.go
	"Suffer Damage": &SufferDamage{},
	// tlt.go
	"Twin Laser Turret": &TwinLaserTurret{},
}
