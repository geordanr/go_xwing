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
    fel := ship.ShipFactory["TIE Interceptor"]()
    fel.Name = "Soontir Fel"
    fel.Agility++
    fel.FocusTokens = 2
    fel.EvadeTokens = 1

    // Accuracy-Corrected B-Wings
    bwings := make([]ship.Ship, 2)
    for i := range(bwings) {
	bwings[i] = ship.ShipFactory["B-Wing"]()
	bwings[i].Name = fmt.Sprintf("B-Wing %d", i + 1)
	bwings[i].FocusTokens = 1
    }

    // Create attack list
    attacks := make([]attack.Attack, len(bwings) + 1)

    // Soontir shoots first
    attacks[0] = attack.Attack{
	Attacker: &fel,
	NumAttackDice: 4,
	AttackerModifications: []attack.Modification{
	    attack.Modifications["Offensive Focus"],
	},
	Defender: &bwings[0],
	NumDefenseDice: 1,
	DefenderModifications: []attack.Modification{
	    attack.Modifications["Defensive Focus"],
	},
    }

    // Then the B-Wings
    for i := 0; i < len(bwings); i++ {
	attacks[i+1] = attack.Attack{
	    Attacker: &bwings[i],
	    NumAttackDice: 3,
	    AttackerModifications: []attack.Modification{
		attack.Modifications["Offensive Focus"],
		attack.Modifications["Accuracy Corrector"],
	    },
	    Defender: &fel,
	    NumDefenseDice: 1,
	    DefenderModifications: []attack.Modification{
		attack.Modifications["Defensive Focus"],
		attack.Modifications["Use Evade Token"],
	    },
	}
    }

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
