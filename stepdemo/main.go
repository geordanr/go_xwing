package main

import (
	"fmt"
	"github.com/geordanr/go_xwing/attack"
	"github.com/geordanr/go_xwing/attack/runner"
	"github.com/geordanr/go_xwing/attack/step"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/gamestate"
	"github.com/geordanr/go_xwing/interfaces"
	"github.com/geordanr/go_xwing/ship"
	"github.com/geordanr/go_xwing/stats"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	nIterations := 1000
	bufsz := 8
	runner := runner.New(step.All, bufsz)
	output := make(chan interfaces.GameState, bufsz)
	attackerStats := ShipStats{
		Hits:  new(stats.Integers),
		Crits: new(stats.Integers),
	}
	defenderStats := ShipStats{
		Hits:  new(stats.Integers),
		Crits: new(stats.Integers),
	}

	go runner.Run(output)

	for i := 0; i < nIterations; i++ {
		runner.InjectState(makeState(&attackerStats, &defenderStats))
	}

	for i := 0; i < nIterations; i++ {
		<-output
	}

	fmt.Println("Attacker")
	fmt.Println("--------")
	fmt.Printf("Average hits : %3.2f\n", attackerStats.Hits.Average())
	fmt.Printf("Stddev hits  : %3.2f\n", attackerStats.Hits.Stddev())
	fmt.Printf("Average crits: %3.2f\n", attackerStats.Crits.Average())
	fmt.Printf("Stddev crits : %3.2f\n", attackerStats.Crits.Stddev())
	fmt.Println()
	fmt.Println("Defender")
	fmt.Println("--------")
	fmt.Printf("Average hits : %3.2f\n", defenderStats.Hits.Average())
	fmt.Printf("Stddev hits  : %3.2f\n", defenderStats.Hits.Stddev())
	fmt.Printf("Average crits: %3.2f\n", defenderStats.Crits.Average())
	fmt.Printf("Stddev crits : %3.2f\n", defenderStats.Crits.Stddev())
}

func makeState(attackerStats, defenderStats *ShipStats) *gamestate.GameState {
	attacker := ship.New("TIE Fighter", 2, 3, 3, 0)
	defender := ship.New("X-Wing", 3, 2, 3, 2)
	sufferDamageMods := step.All["Suffer Damage"].Mods()
	sufferDamageMods = append(sufferDamageMods, &TabulateStats{
		ship:  attacker,
		stats: attackerStats,
		actor: constants.ATTACKER,
	})
	sufferDamageMods = append(sufferDamageMods, &TabulateStats{
		ship:  defender,
		stats: defenderStats,
		actor: constants.ATTACKER,
	})

	mods := map[string][]interfaces.Modification{
		"Suffer Damage": sufferDamageMods,
	}
	state := gamestate.GameState{}
	state.EnqueueAttack(attack.New(attacker, defender, mods))
	state.EnqueueAttack(attack.New(defender, attacker, mods))
	return &state
}

type ShipStats struct {
	Hits  *stats.Integers
	Crits *stats.Integers
}

type TabulateStats struct {
	ship  *ship.Ship
	stats *ShipStats
	actor constants.ModificationActor
}

func (mod *TabulateStats) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	if ship == mod.ship {
		mod.stats.Hits.Update(int(state.HitsLanded()))
		mod.stats.Crits.Update(int(state.CritsLanded()))
	}
}

func (mod TabulateStats) Actor() constants.ModificationActor          { return mod.actor }
func (mod *TabulateStats) SetActor(actor constants.ModificationActor) { mod.actor = actor }
func (mod TabulateStats) String() string                              { return "Tabulate Stats" }
