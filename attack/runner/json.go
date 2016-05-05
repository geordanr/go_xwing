package runner

import (
	"encoding/json"
	"errors"
	// "fmt"
	"github.com/geordanr/go_xwing/ship"
	"io/ioutil"
)

// ShipJSONSchema represents ship stats for a ship chassis type.
type ShipJSONSchema struct {
	Name    string `json:"name"`
	Attack  uint   `json:"attack"`
	Agility uint   `json:"agility"`
	Hull    uint   `json:"hull"`
	Shields uint   `json:"shields"`
}

// CombatantJSONSchema represents a combatant in the simulation.
type CombatantJSONSchema struct {
	Name          string `json:"name"`       // unique identifier for a combatant
	ShipType      string `json:"ship"`       // ship chassis name (e.g. "X-Wing")
	Skill         uint   `json:"skill"`      // pilot skill (not used?)
	HasInitiative bool   `json:"initiative"` // whether this combatant has initiative in the event of tied skill (not used?)
	Tokens        TokenJSONSchema
}

// TokenJSONSchema represents the state of tokens for a combatant at the start of the combat phase.
type TokenJSONSchema struct {
	FocusTokens uint   `json:"focus"`
	EvadeTokens uint   `json:"evade"`
	TargetLock  string `json:"targetlock"`
}

// AttackJSONSchema represents the parameters for a single attack in the combat phase.
type AttackJSONSchema struct {
	Attacker      string                `json:"attacker"` // identifier for attacking combatant
	Defender      string                `json:"defender"` // identifier for defending combatant
	Modifications map[string][][]string `json:"mods"`     // maps attack step to list of {actor, modificationName}
}

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
