package serialization

// ShipJSONSchema represents ship stats for a ship chassis type.
type ShipJSONSchema struct {
	Name    string `json:"name"`
	Attack  uint   `json:"attack"`
	Agility uint   `json:"agility"`
	Hull    uint   `json:"hull"`
	Shields uint   `json:"shields"`
}

// SimulationJSONSchema represents the specifications for a simulation.
type SimulationJSONSchema struct {
	Iterations  int                   `json:"iterations"`
	Combatants  []CombatantJSONSchema `json:"combatants"`
	AttackQueue []AttackJSONSchema    `json:"attack_queue"`
}

// CombatantJSONSchema represents a combatant in the simulation.
type CombatantJSONSchema struct {
	Name          string `json:"name"`       // unique identifier for a combatant
	ShipType      string `json:"ship"`       // ship chassis name (e.g. "X-Wing")
	Skill         uint   `json:"skill"`      // pilot skill
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
