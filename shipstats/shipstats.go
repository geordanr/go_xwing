package shipstats

import (
	"fmt"
	"github.com/geordanr/go_xwing/histogram"
	"github.com/geordanr/go_xwing/ship"
	"github.com/geordanr/go_xwing/stats"
)

type Stats struct {
	hull        *CombinedStat
	shields     *CombinedStat
	focusTokens *CombinedStat
	evadeTokens *CombinedStat
	targetLocks *CombinedStat
}

type CombinedStat struct {
	Stats     *stats.Integers
	Histogram *histogram.Integers
}

func New() *Stats {
	s := Stats{}

	s.hull = newCombinedStat()
	s.shields = newCombinedStat()
	s.focusTokens = newCombinedStat()
	s.evadeTokens = newCombinedStat()
	s.targetLocks = newCombinedStat()

	return &s
}

func (s Stats) Hull() CombinedStat        { return *s.hull }
func (s Stats) Shields() CombinedStat     { return *s.shields }
func (s Stats) FocusTokens() CombinedStat { return *s.focusTokens }
func (s Stats) EvadeTokens() CombinedStat { return *s.evadeTokens }
func (s Stats) TargetLocks() CombinedStat { return *s.targetLocks }

func (s *Stats) Update(ship *ship.Ship) {
	hull := int(ship.Hull())
	s.hull.Stats.Update(hull)
	s.hull.Histogram.Update(hull)

	shields := int(ship.Shields())
	s.shields.Stats.Update(shields)
	s.shields.Histogram.Update(shields)

	focusTokens := int(ship.FocusTokens())
	s.focusTokens.Stats.Update(focusTokens)
	s.focusTokens.Histogram.Update(focusTokens)

	evadeTokens := int(ship.EvadeTokens())
	s.evadeTokens.Stats.Update(evadeTokens)
	s.evadeTokens.Histogram.Update(evadeTokens)

	targetLocks := 0
	if len(ship.TargetLock()) > 0 {
		targetLocks = 1
	}
	s.targetLocks.Stats.Update(targetLocks)
	s.targetLocks.Histogram.Update(targetLocks)
}

func (s Stats) String() string {
	return fmt.Sprintf("Hull\t: average=%2.3f stddev=%2.3f\nShields\t: average=%2.3f stddev=%2.3f\n", s.hull.Stats.Average(), s.hull.Stats.Stddev(), s.shields.Stats.Average(), s.shields.Stats.Stddev())
}

func newCombinedStat() *CombinedStat {
	return &CombinedStat{
		Stats:     stats.New(),
		Histogram: histogram.New(),
	}
}

// StatMap returns a map from human readable name to CombinedStats.
// (because apparently reflection is slow)
func (s Stats) StatMap() map[string]*CombinedStat {
	return map[string]*CombinedStat{
		"hull":        s.hull,
		"shields":     s.shields,
		"focusTokens": s.focusTokens,
		"evadeTokens": s.evadeTokens,
		"targetLocks": s.targetLocks,
	}
}
