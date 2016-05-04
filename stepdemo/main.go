package main

import (
	"fmt"
	"github.com/geordanr/go_xwing/attack"
	"github.com/geordanr/go_xwing/attack/modification"
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

	nIterations := 100000
	runner := runner.New(step.Steps, nIterations)
	output := make(chan interfaces.GameState, nIterations)
	hitStats := new(stats.Integers)
	critStats := new(stats.Integers)

	go runner.Run(output)

	for i := 0; i < nIterations; i++ {
		runner.InjectState(makeState(hitStats, critStats))
	}

	for i := 0; i < nIterations; i++ {
		<-output
	}

	fmt.Printf("Average hits : %3.2f\n", hitStats.Average())
	fmt.Printf("Stddev hits  : %3.2f\n", hitStats.Stddev())
	fmt.Printf("Average crits: %3.2f\n", critStats.Average())
	fmt.Printf("Stddev crits : %3.2f\n", critStats.Stddev())
}

func makeState(hitStats, critStats *stats.Integers) *gamestate.GameState {
	attacker := ship.New("TIE Fighter", 2, 3, 3, 0)
	defender := ship.New("X-Wing", 3, 2, 3, 2)
	mods := map[string][]interfaces.Modification{
		"Suffer Damage": []interfaces.Modification{
			modification.All["Suffer Damage"],
			&TabulateHits{stats: hitStats},
			&TabulateCrits{stats: critStats},
		},
	}
	state := gamestate.GameState{}
	state.EnqueueAttack(attack.New(attacker, defender, mods))
	// state.EnqueueAttack(attack.New(defender, attacker, mods))
	return &state
}

type TabulateHits struct {
	stats *stats.Integers
	actor constants.ModificationActor
}

func (mod *TabulateHits) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	mod.stats.Update(int(state.HitsLanded()))
}

func (mod TabulateHits) Actor() constants.ModificationActor          { return mod.actor }
func (mod *TabulateHits) SetActor(actor constants.ModificationActor) { mod.actor = actor }
func (mod TabulateHits) String() string                              { return "Tabulate Hits" }

type TabulateCrits struct {
	stats *stats.Integers
	actor constants.ModificationActor
}

func (mod *TabulateCrits) ModifyState(state interfaces.GameState, ship interfaces.Ship) {
	mod.stats.Update(int(state.CritsLanded()))
}

func (mod TabulateCrits) Actor() constants.ModificationActor          { return mod.actor }
func (mod *TabulateCrits) SetActor(actor constants.ModificationActor) { mod.actor = actor }
func (mod TabulateCrits) String() string                              { return "Tabulate Crits" }
