package modification

import (
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/interfaces"
)

var attackerSpendFocus = SpendFocus{actor: constants.ATTACKER}
var defenderSpendFocus = SpendFocus{actor: constants.DEFENDER}
var attackerRollDice = RollDice{actor: constants.ATTACKER}
var defenderRollDice = RollDice{actor: constants.DEFENDER}

var All = map[string]interfaces.Modification{
	// compareresults.go
	"Compare Results": &CompareResults{},
	// focus.go
	"Attacker Spend Focus": &attackerSpendFocus,
	"Defender Spend Focus": &defenderSpendFocus,
	// roll.go
	"Roll Attack Dice": &attackerRollDice,
	"Roll Defense Dice": &defenderRollDice,
	// sufferdamage.go
	"Suffer Damage": &SufferDamage{},
}
