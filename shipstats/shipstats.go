package shipstats

import (
	"fmt"
	"github.com/geordanr/go_xwing/histogram"
	"github.com/geordanr/go_xwing/ship"
	"github.com/geordanr/go_xwing/stats"
)

type Stats struct {
	hull    *CombinedStat
	shields *CombinedStat
}

type CombinedStat struct {
	Stats     *stats.Integers
	Histogram *histogram.Integers
}

func New() *Stats {
	s := Stats{}

	s.hull = newCombinedStat()
	s.shields = newCombinedStat()

	return &s
}

func (s Stats) Hull() CombinedStat    { return *s.hull }
func (s Stats) Shields() CombinedStat { return *s.shields }

func (s *Stats) Update(ship *ship.Ship) {
	s.hull.Stats.Update(int(ship.Hull()))
	s.hull.Histogram.Update(int(ship.Hull()))
	s.shields.Stats.Update(int(ship.Shields()))
	s.shields.Histogram.Update(int(ship.Shields()))
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
