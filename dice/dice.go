package dice

func RollAttackDice(numDice uint8) Results {
	results := make(Results, numDice)
	for i := range results {
		results[i] = new(AttackDie)
		results[i].Roll()
	}
	return results
}

func RollDefenseDice(numDice uint8) Results {
	results := make(Results, numDice)
	for i := range results {
		results[i] = new(DefenseDie)
		results[i].Roll()
	}
	return results
}
