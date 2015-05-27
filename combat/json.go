package combat

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "github.com/geordanr/go_xwing/attack"
    "github.com/geordanr/go_xwing/ship"
)

type iterationsJSONSchema struct {
    Iterations int
}

type listVsListJSONSchema struct {
    iterationsJSONSchema
    ListOne []modifiedShipJSONSchema
    ListTwo []modifiedShipJSONSchema
}

type modifiedShipJSONSchema struct {
    Name string
    ShipType string
    StatOverrides statOverrideJSONSchema
    AttackerModifications []string
    DefenderModifications []string
    Actions []string
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
    // Just need to parse out iterations first
    var iter iterationsJSONSchema
    if err := json.Unmarshal(b, &iter); err != nil {
	return nil, nil, err
    }
    iterations := iter.Iterations

    atkFactory := func () ([]attack.Attack, error) {
	attacks, _, err := AttacksFromJSON(b)
	if err != nil {
	    return nil, err
	}
	return attacks, nil
    }

    s, r, err := Simulate(atkFactory, iterations)
    return s, r, err
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
    listOne, err := makeList(l.ListOne)
    if err != nil {
	return nil, 0, err
    }
    listTwo, err := makeList(l.ListTwo)
    if err != nil {
	return nil, 0, err
    }
    return ListVersusList(listOne, listTwo), l.Iterations, nil
}

func makeList(l []modifiedShipJSONSchema) ([]ModifiedShip, error) {
    ret := make([]ModifiedShip, len(l))
    for i, s := range(l) {
	ret[i] = ModifiedShip{}
	newship := &ret[i]
	shipfactory, prs := ship.ShipFactory[s.ShipType]
	if !prs {
	    return nil, fmt.Errorf("Unrecognized ship %s", s.ShipType)
	}
	newship.Ship = shipfactory()
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
	    newship.AttackerModifications[j], prs = attack.Modifications[mod]
	    if !prs {
		return nil, fmt.Errorf("Unrecognized attack modifier %s", mod)
	    }
	}

	newship.DefenderModifications = make([]attack.Modification, len(s.DefenderModifications))
	for j, mod := range(s.DefenderModifications) {
	    newship.DefenderModifications[j], prs = attack.Modifications[mod]
	    if !prs {
		return nil, fmt.Errorf("Unrecognized defense modifier %s", mod)
	    }
	}

	newship.Actions = make([]ship.Action, len(s.Actions))
	for j, action := range(s.Actions) {
	    newship.Actions[j], prs = ship.Actions[action]
	    if !prs {
		return nil, fmt.Errorf("Unrecognized action %s", action)
	    }
	}
    }
    return ret, nil
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
