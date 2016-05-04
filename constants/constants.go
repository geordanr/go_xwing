package constants

type ModificationActor uint8

const (
	ATTACKER ModificationActor = iota
	DEFENDER
	INITIATIVE
)
