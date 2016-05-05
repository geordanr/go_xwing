package serialization

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/geordanr/go_xwing/attack"
	"github.com/geordanr/go_xwing/attack/modification"
	"github.com/geordanr/go_xwing/attack/runner"
	"github.com/geordanr/go_xwing/attack/step"
	"github.com/geordanr/go_xwing/constants"
	"github.com/geordanr/go_xwing/gamestate"
	"github.com/geordanr/go_xwing/interfaces"
	"github.com/geordanr/go_xwing/ship"
	"io/ioutil"
	"math"
)

// MAX_ITERATIONS is the maximum number of game states to process.
const MAX_ITERATIONS = 100000

// ShipsFromJSONPath reads a JSON file and returns a map of ship names
// to factory functions.  The factory function accepts a string argument
// which will be the ship's name.
func ShipsFromJSONPath(path string) (map[string]func(string) *ship.Ship, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return shipsFromJSON(bytes)
}

// shipsFromJSON is a helper function (primarily for testing)
func shipsFromJSON(b []byte) (map[string]func(string) *ship.Ship, error) {
	data := map[string][]ShipJSONSchema{}
	err := fromJSON(b, &data)
	if err != nil {
		return nil, err
	}

	factory := map[string]func(string) *ship.Ship{}
	// Expect to find the array of ships in the object property "ships"
	shipList, exists := data["ships"]
	if !exists {
		return nil, errors.New("Expected to find JSON object property 'ships'")
	}

	for _, shipStats := range shipList {
		shipStats := shipStats // silly closure trick
		factory[shipStats.Name] = func(name string) *ship.Ship {
			return ship.New(name, shipStats.Attack, shipStats.Agility, shipStats.Hull, shipStats.Shields)

		}
	}

	return factory, nil
}

func FromJSONPath(path string, shipFactory map[string]func(string) *ship.Ship) (<-chan interfaces.GameState, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return FromJSON(bytes, shipFactory)
}

// FromJSON reads the JSON bytestream, creates a Runner to run the simulation, and returns an output channel to read game states from.
func FromJSON(b []byte, shipFactory map[string]func(string) *ship.Ship) (<-chan interfaces.GameState, error) {
	makeState := func() (interfaces.GameState, error) {
		state := gamestate.GameState{}
		data := SimulationJSONSchema{}
		err := fromJSON(b, &data)
		if err != nil {
			return nil, err
		}

		combatants := map[string]interfaces.Ship{}
		for _, combatant := range data.Combatants {
			shipFunc, exists := shipFactory[combatant.ShipType]
			if !exists {
				return nil, errors.New(fmt.Sprintf("No data for ship type '%s'", combatant.ShipType))
			}
			cbt := shipFunc(combatant.Name)
			cbt.SetFocusTokens(combatant.Tokens.FocusTokens)
			cbt.SetEvadeTokens(combatant.Tokens.EvadeTokens)
			// TODO target locks
			combatants[combatant.Name] = cbt
		}
		state.SetCombatants(combatants)

		// eventually we need to reverse the list
		tmp := []*attack.Attack{}
		for _, atkParams := range data.AttackQueue {
			attacker, exists := combatants[atkParams.Attacker]
			if !exists {
				return nil, errors.New(fmt.Sprintf("Attacker '%s' not found", atkParams.Attacker))
			}

			defender, exists := combatants[atkParams.Defender]
			if !exists {
				return nil, errors.New(fmt.Sprintf("Defender '%s' not found", atkParams.Defender))
			}

			mods := map[string][]interfaces.Modification{}
			for stepName, stepMods := range atkParams.Modifications {
				mods[stepName] = []interfaces.Modification{}
				for _, modParams := range stepMods {
					var a constants.ModificationActor
					actor := modParams[0]
					modName := modParams[1]

					modFactory, exists := modification.All[modName]
					if !exists {
						return nil, errors.New(fmt.Sprintf("No modification '%s'", modName))
					}
					mod := modFactory()
					switch actor {
					case "attacker":
						a = constants.ATTACKER
					case "defender":
						a = constants.DEFENDER
					case "initiative":
						a = constants.INITIATIVE
					default:
						a = constants.IGNORE
					}
					mod.SetActor(a)

					mods[stepName] = append(mods[stepName], mod)
				}
			}

			tmp = append(tmp, attack.New(attacker, defender, mods))
		}
		// Finally reverse the attack list to put it in the correct queue order
		for i := len(tmp) - 1; i > -1; i-- {
			state.EnqueueAttack(tmp[i])
		}

		return &state, nil
	}

	// need to unpack number of iterations
	data := SimulationJSONSchema{}
	err := fromJSON(b, &data)
	if err != nil {
		// fmt.Println("Returning error:", err)
		return nil, err
	}
	nStates := int(math.Min(float64(MAX_ITERATIONS), float64(data.Iterations)))
	// fmt.Println("Running", nStates, "iterations")
	runner := runner.New(step.All, nStates)
	runnerOut := make(chan interfaces.GameState, nStates)
	output := make(chan interfaces.GameState, nStates)
	go runner.Run(runnerOut)

	for i := 0; i < nStates; i++ {
		state, err := makeState()
		if err != nil {
			// fmt.Println("makestate error", err)
			return nil, err
		}
		runner.InjectState(state)
	}

	go func() {
		for i := 0; i < nStates; i++ {
			output <- <-runnerOut
		}
		close(output)
	}()

	return output, nil
}

// fromJSON reads the JSON bytestream b and unmarshals it into the structure data.
func fromJSON(b []byte, data interface{}) error {
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}
	return nil
}

// fromJSONPath opens the given path and unmarshals the JSON inside into the structure data.
func fromJSONPath(path string, data interface{}) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return fromJSON(bytes, data)
}
