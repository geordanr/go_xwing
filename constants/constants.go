package constants

type ModificationActor uint8

const (
	IGNORE ModificationActor = iota
	ATTACKER
	DEFENDER
	INITIATIVE
)
