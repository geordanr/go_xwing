package step

import (
	"github.com/geordanr/go_xwing/attack/modification"
	"github.com/geordanr/go_xwing/interfaces"
)

var Steps = map[string]interfaces.Step{
	"__START__":                 &Step{name: "__START__", next: "Declare Target"},
	"Declare Target":            &Step{name: "Declare Target", next: "Roll Attack Dice"},
	"Roll Attack Dice":          &Step{name: "Roll Attack Dice", mods: []interfaces.Modification{modification.All["Roll Attack Dice"]}, next: "Modify Attack Dice"},
	"Modify Attack Dice":        &Step{name: "Modify Attack Dice", next: "Roll Defense Dice"},
	"Roll Defense Dice":         &Step{name: "Roll Defense Dice", mods: []interfaces.Modification{modification.All["Roll Defense Dice"]}, next: "Modify Defense Dice"},
	"Modify Defense Dice":       &Step{name: "Modify Defense Dice", next: "Compare Results"},
	"Compare Results":           &Step{name: "Compare Results", mods: []interfaces.Modification{modification.All["Compare Results"]}, next: "Suffer Damage"},
	"Suffer Damage":             &Step{name: "Suffer Damage", mods: []interfaces.Modification{modification.All["Suffer Damage"]}, next: "After Attacking/Defending"},
	"After Attacking/Defending": &Step{name: "After Attacking/Defending", next: "Perform Attack Twice"},
	"Perform Attack Twice":      &Step{name: "Perform Attack Twice", next: "Perform Additional Attack"},
	"Perform Additional Attack": &Step{name: "Perform Additional Attack"},
}
