// Package stats provides a container to collect statistics.
package stats

import "math"

type Integers struct {
    n int
    sum int
    sumSquares int
}

func (s *Integers) Update(val int) {
    s.n++
    s.sum += val
    s.sumSquares += val * val
}

func (s *Integers) Average() float64 {
    return float64(s.sum) / float64(s.n)
}

func (s *Integers) Stddev() float64 {
    return math.Sqrt((float64(s.sumSquares) / float64(s.n)) - math.Pow(s.Average(), 2))
}
