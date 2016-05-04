package dice

import (
	"fmt"
)

type Results []Rollable

func (results Results) countResults(result Result) (n uint8) {
	for i := range results {
		if results[i].Result() == result {
			n++
		}
	}
	return
}

func (results Results) Blanks() uint8 {
	return results.countResults(BLANK)
}

func (results Results) Focuses() uint8 {
	return results.countResults(FOCUS)
}

func (results Results) Hits() uint8 {
	return results.countResults(HIT)
}

func (results Results) Crits() uint8 {
	return results.countResults(CRIT)
}

func (results Results) Evades() uint8 {
	return results.countResults(EVADE)
}

func (results Results) String() string {
	return fmt.Sprintf("Blanks: %d, Focuses: %d, Hits: %d, Crits: %d, Evades: %d", results.Blanks(), results.Focuses(), results.Hits(), results.Crits(), results.Evades())
}

func (results *Results) RerollUpto(numToReroll uint, filter func(Result) bool) *Results {
	var rerolled uint = 0
	for i := range *results {
		die := (*results)[i]
		if filter(die.Result()) && die.IsRerollable() && rerolled < numToReroll {
			die.Reroll()
			rerolled++
		}
	}
	return results
}

func (results *Results) Reroll(filter func(Result) bool) *Results {
	return results.RerollUpto(uint(len(*results)), filter)
}

func (results *Results) ConvertUpto(numToConvert uint, from, to Result) *Results {
	var rerolled uint = 0
	for i := range *results {
		die := (*results)[i]
		if !die.Locked() && die.Result() == from && rerolled < numToConvert {
			die.SetResult(to)
			rerolled++
		}
	}
	return results
}

func (results *Results) ConvertAll(from, to Result) *Results {
	return results.ConvertUpto(uint(len(*results)), from, to)
}
