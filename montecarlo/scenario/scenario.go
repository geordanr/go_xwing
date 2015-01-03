package scenario

import (
    // "fmt"
    "math"
    "github.com/geordanr/go_xwing/dice"
)

type Scenario struct {
    AttackResults *dice.Results
    DefenseResults *dice.Results
    NumAttackerFocus uint
    NumDefenderFocus uint
    NumDefenderEvade uint
    DefenderModifiesAttackDice []Modification
    AttackerModifiesAttackDice []Modification
    AttackerModifiesDefenseDice []Modification
    DefenderModifiesDefenseDice []Modification
}

func (scenario *Scenario) Run() (*Scenario) {
    for i := range(scenario.DefenderModifiesAttackDice) {
	scenario.DefenderModifiesAttackDice[i].Modify(scenario)
    }

    for i := range(scenario.AttackerModifiesAttackDice) {
	scenario.AttackerModifiesAttackDice[i].Modify(scenario)
    }

    for i := range(scenario.AttackerModifiesDefenseDice) {
	scenario.AttackerModifiesDefenseDice[i].Modify(scenario)
    }

    for i := range(scenario.DefenderModifiesDefenseDice) {
	scenario.DefenderModifiesDefenseDice[i].Modify(scenario)
    }

    return scenario
}

// CompareResults returns the number of hits and crits that made it through.
func (scenario Scenario) CompareResults() (hits, crits uint8) {
    hits = scenario.AttackResults.Hits()
    crits = scenario.AttackResults.Crits()
    evades := scenario.DefenseResults.Evades()
    // fmt.Printf("Compare: hits=%d, crits=%d, evades=%d\n", hits, crits, evades)
    // Spend evade results on hits first
    evadesSpentOnHits := uint8(math.Min(float64(hits), float64(evades)))
    // fmt.Printf("Evades spent on hits: %d\n", evadesSpentOnHits)
    hits -= evadesSpentOnHits
    evades -= evadesSpentOnHits
    // fmt.Printf("Hits now %d\nEvades now %d\n", hits, evades)
    // Then crits
    crits = uint8(math.Max(0, float64(crits - evades)))
    // fmt.Printf("Crits now %d\n", crits)
    return hits, crits
}
