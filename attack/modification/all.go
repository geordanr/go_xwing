package modification

import (
	"github.com/geordanr/go_xwing/interfaces"
)

var All = map[string]interfaces.Modification{
	// compareresults.go
	"Compare Results": &CompareResults{},
	// focus.go
	"Spend Focus": &SpendFocus{},
	// roll.go
	"Roll Dice": &RollDice{},
	// sufferdamage.go
	"Suffer Damage": &SufferDamage{},
}
