package combat

import (
    "encoding/json"
    "io/ioutil"
    "github.com/geordanr/go_xwing/attack"
    "github.com/geordanr/go_xwing/ship"
)

type listVsListJSONSchema struct {
    ListOne []modifiedShipJSONSchema
    ListTwo []modifiedShipJSONSchema
}

type modifiedShipJSONSchema struct {
    Name string
    ShipType string
    StatOverrides statOverrideJSONSchema
    AttackerModifications []string
    DefenderModifications []string
}

// Uses pointers so we can tell the difference between an unspecified
// value and a value of 0.
type statOverrideJSONSchema struct {
    Attack *uint
    Agility *uint
    Hull *uint
    Shields *uint
    FocusTokens *uint
    EvadeTokens *uint
}

func AttacksFromJSONPath(path string) ([]attack.Attack, error) {
    data, err := ioutil.ReadFile(path)
    if err != nil {
	return nil, err
    }
    return AttacksFromJSON(data)
}

func AttacksFromJSON(b []byte) ([]attack.Attack, error) {
    var l listVsListJSONSchema
    if err := json.Unmarshal(b, &l); err != nil {
	return nil, err
    }
    listOne := makeList(l.ListOne)
    listTwo := makeList(l.ListTwo)
    return ListVersusList(listOne, listTwo), nil
}

func makeList(l []modifiedShipJSONSchema) []ModifiedShip {
    ret := make([]ModifiedShip, len(l))
    for i, s := range(l) {
	ret[i] = ModifiedShip{}
	newship := &ret[i]
	newship.Ship = ship.ShipFactory[s.ShipType]()
	newship.Name = s.Name

	if s.StatOverrides.Attack != nil {
	    newship.Attack = *s.StatOverrides.Attack
	}

	if s.StatOverrides.Agility != nil {
	    newship.Agility = *s.StatOverrides.Agility
	}

	if s.StatOverrides.Hull != nil {
	    newship.Hull = *s.StatOverrides.Hull
	}

	if s.StatOverrides.Shields != nil {
	    newship.Shields = *s.StatOverrides.Shields
	}

	if s.StatOverrides.FocusTokens != nil {
	    newship.FocusTokens = *s.StatOverrides.FocusTokens
	}

	if s.StatOverrides.EvadeTokens != nil {
	    newship.EvadeTokens = *s.StatOverrides.EvadeTokens
	}

	newship.AttackerModifications = make([]attack.Modification, len(s.AttackerModifications))
	for j, mod := range(s.AttackerModifications) {
	    newship.AttackerModifications[j] = attack.Modifications[mod]
	}

	newship.DefenderModifications = make([]attack.Modification, len(s.DefenderModifications))
	for j, mod := range(s.DefenderModifications) {
	    newship.DefenderModifications[j] = attack.Modifications[mod]
	}
    }
    return ret
}
