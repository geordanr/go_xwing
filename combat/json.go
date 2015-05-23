package combat

import (
    "encoding/json"
    "io/ioutil"
    "github.com/geordanr/go_xwing/attack"
    "github.com/geordanr/go_xwing/ship"
)

type listVsListJSONSchema struct {
    Iterations int
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

func SimulateFromJSON(b []byte) (*statsByShipName, *resultsByShipName, error) {
    attacks, iterations, err := AttacksFromJSON(b)
    if err != nil {
	return nil, nil, err
    }

    atkFactory := func () []attack.Attack {
	return attacks
    }

    s, r := Simulate(atkFactory, iterations)
    return s, r, nil
}

func SimulateFromJSONPath(path string) (*statsByShipName, *resultsByShipName, error) {
    data, err := ioutil.ReadFile(path)
    if err != nil {
	return nil, nil, err
    }
    return SimulateFromJSON(data)
}

func AttacksFromJSONPath(path string) ([]attack.Attack, int, error) {
    data, err := ioutil.ReadFile(path)
    if err != nil {
	return nil, 0, err
    }
    return AttacksFromJSON(data)
}

func AttacksFromJSON(b []byte) ([]attack.Attack, int, error) {
    var l listVsListJSONSchema
    if err := json.Unmarshal(b, &l); err != nil {
	return nil, 0, err
    }
    listOne := makeList(l.ListOne)
    listTwo := makeList(l.ListTwo)
    return ListVersusList(listOne, listTwo), l.Iterations, nil
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

type combinedSimResults struct {
    Stats *stats
    Results *jsonableSimResult
}

type jsonableSimResult struct {
    HitAverage float64
    HitStddev float64
    HitHistogram map[string]int

    CritAverage float64
    CritStddev float64
    CritHistogram map[string]int

    HullAverage float64
    HullStddev float64
    HullHistogram map[string]int

    ShieldAverage float64
    ShieldStddev float64
    ShieldHistogram map[string]int
}

func SimResultsToJSON(shipStats *statsByShipName, shipResults *resultsByShipName) ([]byte, error) {
    res := make(map[string]combinedSimResults, len(*shipStats))
    for name, s := range(*shipStats) {
	shipResult := (*shipResults)[name]
	j := jsonableSimResult{
	    HitAverage: shipResult.HitAverage,
	    HitStddev: shipResult.HitStddev,
	    HitHistogram: shipResult.HitHistogram.ToStrMap(),

	    CritAverage: shipResult.CritAverage,
	    CritStddev: shipResult.CritStddev,
	    CritHistogram: shipResult.CritHistogram.ToStrMap(),

	    HullAverage: shipResult.HullAverage,
	    HullStddev: shipResult.HullStddev,
	    HullHistogram: shipResult.HullHistogram.ToStrMap(),

	    ShieldAverage: shipResult.ShieldAverage,
	    ShieldStddev: shipResult.ShieldStddev,
	    ShieldHistogram: shipResult.ShieldHistogram.ToStrMap(),
	}
	res[name] = combinedSimResults{
	    Stats: s,
	    Results: &j,
	}
    }
    bytes, err := json.Marshal(res)
    if err != nil {
	return nil, err
    }
    return bytes, nil

}
