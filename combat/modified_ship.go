package combat

import (
	"github.com/geordanr/go_xwing/attack"
	"github.com/geordanr/go_xwing/ship"
)

type ModifiedShip struct {
	ship.Ship
	DefenderModifications []attack.Modification
	AttackerModifications []attack.Modification
}

//TODO: add list sorting (initiative) and target selection (focused fire) by some user defined order
func ListVersusList( listOne []ModifiedShip, listTwo []ModifiedShip) []attack.Attack {
	var countOfAttackSets int
	countOfAttackSets = len(listOne) + len(listTwo)

	attacks := make([]attack.Attack, countOfAttackSets)

	for i := 0; i < len(listOne); i++ {
		attacks[i] = attack.Attack{
			Attacker: &listOne[i].Ship,
			NumAttackDice: uint8(listOne[i].Attack),
			AttackerModifications: listOne[i].AttackerModifications,
			Defender: &listTwo[0].Ship,
			NumDefenseDice: uint8(listTwo[0].Agility),
			DefenderModifications: listTwo[0].DefenderModifications,
		}
	}

	var lenListOne int
	lenListOne = len(listOne)

	for j := 0; j < len(listTwo); j++ {
		attacks[j + lenListOne] = attack.Attack{
			Attacker: &listTwo[j].Ship,
			NumAttackDice: uint8(listTwo[j].Attack),
			AttackerModifications: listTwo[j].AttackerModifications,
			Defender: &listOne[0].Ship,
			NumDefenseDice: uint8(listOne[0].Agility),
			DefenderModifications: listOne[0].DefenderModifications,
		}
	}

	return attacks
}

