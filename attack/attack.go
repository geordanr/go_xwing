package attack

import (
    // "fmt"
    "math"
    "github.com/geordanr/go_xwing/dice"
    "github.com/geordanr/go_xwing/ship"
)

type Attack struct {
    Attacker *ship.Ship
    NumAttackDice uint8
    AttackerModifications []Modification
    AttackResults *dice.Results

    Defender *ship.Ship
    NumDefenseDice uint8
    DefenderModifications []Modification
    DefenseResults *dice.Results
}

// Copy creates a copy of the attack, with nil results.
func (atk Attack) Copy() *Attack {
    newatk := Attack{
	Attacker: atk.Attacker,
	NumAttackDice: atk.NumAttackDice,

	Defender: atk.Defender,
	NumDefenseDice: atk.NumDefenseDice,
    }
    newatk.AttackerModifications = make([]Modification, len(atk.AttackerModifications))
    copy(newatk.AttackerModifications, atk.AttackerModifications)
    newatk.DefenderModifications = make([]Modification, len(atk.DefenderModifications))
    copy(newatk.DefenderModifications, atk.DefenderModifications)

    return &newatk
}

func (atk *Attack) compareResults() (hits, crits uint) {
    hits = uint(atk.AttackResults.Hits())
    crits = uint(atk.AttackResults.Crits())
    evades := uint(atk.DefenseResults.Evades())
    // fmt.Printf("Compare: hits=%d, crits=%d, evades=%d\n", hits, crits, evades)
    // Spend evade results on hits first
    evadesSpentOnHits := uint(math.Min(float64(hits), float64(evades)))
    // fmt.Printf("Evades spent on hits: %d\n", evadesSpentOnHits)
    hits -= evadesSpentOnHits
    evades -= evadesSpentOnHits
    // fmt.Printf("Hits now %d\nEvades now %d\n", hits, evades)
    // Then crits
    crits = uint(math.Max(0, float64(crits - evades)))
    // fmt.Printf("Crits now %d\n", crits)
    return hits, crits
}

// Execute rolls and modifies dice using specified strategies, and assigns damage.
func (atk *Attack) Execute() (uint, uint) {
    // Attacker rolls attack dice

    atkResults := dice.RollAttackDice(atk.NumAttackDice)
    atk.AttackResults = &atkResults

    // Defender modifies attack dice
    for _, mod := range(atk.DefenderModifications) {
	if mod.ModifiesAttackResults() {
	    mod.Modify(atk)
	}
    }
    // Attacker modifies attack dice
    for _, mod := range(atk.AttackerModifications) {
	if mod.ModifiesAttackResults() {
	    mod.Modify(atk)
	}
    }

    // Defender rolls defense dice
    defResults := dice.RollDefenseDice(atk.NumDefenseDice)
    atk.DefenseResults = &defResults

    // Attacker modifies defense dice
    for _, mod := range(atk.AttackerModifications) {
	if mod.ModifiesDefenseResults() {
	    mod.Modify(atk)
	}
    }
    // Defender modifies defense dice
    for _, mod := range(atk.DefenderModifications) {
	if mod.ModifiesDefenseResults() {
	    mod.Modify(atk)
	}
    }

    hits, crits := atk.compareResults()
    atk.Defender.SufferDamage(hits + crits) // treating both as 1 damage for now

    return hits, crits
}
