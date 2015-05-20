package main

import (
    "fmt"
    "math/rand"
    "strings"
    "time"
    "github.com/geordanr/go_xwing/attack"
    "github.com/geordanr/go_xwing/combat"
    "github.com/geordanr/go_xwing/ship"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    _, results := combat.Simulate(twoAccBvsFel, 1000)

    for shipName, res := range(*results) {
	fmt.Println(shipName)
	fmt.Println(strings.Repeat("-", len(shipName)))
	fmt.Printf("Hits Landed: %-.3f (stddev=%-.3f)\n", res.HitAverage, res.HitStddev)
	fmt.Println(res.HitHistogram)
	fmt.Printf("Crits Landed: %-.3f (stddev=%-.3f)\n", res.CritAverage, res.CritStddev)
	fmt.Println(res.CritHistogram)
	fmt.Printf("Hull Remaining: %-.3f (stddev=%-.3f)\n", res.HullAverage, res.HullStddev)
	fmt.Println(res.HullHistogram)
	fmt.Printf("Shields Remaining: %-.3f (stddev=%-.3f)\n", res.ShieldAverage, res.ShieldStddev)
	fmt.Println(res.ShieldHistogram)
    }
}

// Tanked up StealthFel vs. two Accuracy Corrector B-Wings
func twoAccBvsFel() []attack.Attack{
    // StealthFel
	fel := combat.ModifiedShip{}
    fel.Ship = ship.ShipFactory["TIE Interceptor"]()
    fel.Name = "Soontir Fel"
    fel.Agility=4 //stealth device
    fel.FocusTokens = 2 //turtled
    fel.EvadeTokens = 1
	fel.AttackerModifications = []attack.Modification{
		attack.Modifications["Offensive Focus"],
	}
	fel.DefenderModifications =  []attack.Modification{
		attack.Modifications["Defensive Focus"],
		attack.Modifications["Use Evade Token"],
	}


	soontirList := make([]combat.ModifiedShip,1)
	soontirList[0] = fel

    // Accuracy-Corrected B-Wings
    bwings := make([]combat.ModifiedShip, 2)
    for i := range(bwings) {
		bwings[i].Ship = ship.ShipFactory["B-Wing"]()
		bwings[i].Name = fmt.Sprintf("B-Wing %d", i + 1)
		bwings[i].FocusTokens = 1
		bwings[i].AttackerModifications = []attack.Modification{
			attack.Modifications["Offensive Focus"],
			attack.Modifications["Accuracy Corrector"],
		}
    }

	attacks := combat.ListVersusList( soontirList, bwings)
    return attacks
}



func gunner() (hits, crits uint) {
    atk := attack.Attack{
	Attacker: &ship.Ship{FocusTokens: 0},
	NumAttackDice: 3,
	AttackerModifications: []attack.Modification{
	    // attack.Modifications["Predator (high PS)"],
	    // attack.Modifications["Chiraneau"],
	    attack.Modifications["Accuracy Corrector"],
	    // attack.Modifications["Offensive Focus"],
	},

	Defender: &ship.Ship{FocusTokens: 1},
	NumDefenseDice: 3,
	DefenderModifications: []attack.Modification{
	    attack.Modifications["Defensive Focus"],
	},
    }

    hits, crits = atk.Execute()

    if hits + crits == 0 {
	gunnerAtk := atk.Copy()

	h, c := gunnerAtk.Execute()
	hits += h
	crits += c
    }

    return hits, crits
}
