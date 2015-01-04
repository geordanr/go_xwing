package filters

import (
    "testing"
    "github.com/geordanr/go_xwing/dice"
)

func filterHelper(t *testing.T, filter func(dice.Result) bool, name string, good, bad []dice.Result) {
    for i := range(good) {
	if !filter(good[i]) {
	    t.Errorf("Filter %s should have matched %s", name, good[i])
	}
    }
    for i := range(bad) {
	if filter(bad[i]) {
	    t.Errorf("Filter %s should not have matched %s", name, bad[i])
	}
    }
}

func TestBlanks(t *testing.T) {
    good := []dice.Result{dice.BLANK}
    bad := []dice.Result{dice.FOCUS, dice.HIT, dice.CRIT, dice.EVADE}
    filterHelper(t, Blanks, "Blanks", good, bad)
}

func TestFocuses(t *testing.T) {
    good := []dice.Result{dice.FOCUS}
    bad := []dice.Result{dice.BLANK, dice.HIT, dice.CRIT, dice.EVADE}
    filterHelper(t, Focuses, "Focuses", good, bad)
}

func TestBlanksAndFocuses(t *testing.T) {
    good := []dice.Result{dice.BLANK, dice.FOCUS}
    bad := []dice.Result{dice.HIT, dice.CRIT, dice.EVADE}
    filterHelper(t, BlanksAndFocuses, "BlanksAndFocuses", good, bad)
}

func TestHits(t *testing.T) {
    good := []dice.Result{dice.HIT}
    bad := []dice.Result{dice.FOCUS, dice.BLANK, dice.CRIT, dice.EVADE}
    filterHelper(t, Hits, "Hits", good, bad)
}

func TestCrits(t *testing.T) {
    good := []dice.Result{dice.CRIT}
    bad := []dice.Result{dice.FOCUS, dice.HIT, dice.BLANK, dice.EVADE}
    filterHelper(t, Crits, "Crits", good, bad)
}

func TestEvades(t *testing.T) {
    good := []dice.Result{dice.EVADE}
    bad := []dice.Result{dice.FOCUS, dice.HIT, dice.CRIT, dice.BLANK}
    filterHelper(t, Evades, "Evades", good, bad)
}

func TestEverything(t *testing.T) {
    good := []dice.Result{dice.EVADE, dice.FOCUS, dice.HIT, dice.CRIT, dice.BLANK}
    bad := []dice.Result{}
    filterHelper(t, Everything, "Everything", good, bad)
}
