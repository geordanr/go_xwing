package filters

import "github.com/geordanr/xwing/dice"

func Blanks(result dice.Result) bool {
    return result == dice.BLANK
}

func Focuses(result dice.Result) bool {
    return result == dice.FOCUS
}

func BlanksAndFocuses(result dice.Result) bool {
    return result == dice.BLANK || result == dice.FOCUS
}

func Hits(result dice.Result) bool {
    return result == dice.HIT
}

func Crits(result dice.Result) bool {
    return result == dice.CRIT
}

func Evades(result dice.Result) bool {
    return result == dice.EVADE
}
