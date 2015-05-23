package combat

import (
    "math"
    "github.com/geordanr/go_xwing/histogram"
)


type stats struct {
    Iterations int

    HitSum uint
    HitSumSquares uint

    CritSum uint
    CritSumSquares uint

    HullSum uint
    HullSumSquares uint

    ShieldSum uint
    ShieldSumSquares uint
}

type simResult struct {
    HitAverage float64
    HitStddev float64
    HitHistogram histogram.IntHistogram

    CritAverage float64
    CritStddev float64
    CritHistogram histogram.IntHistogram

    HullAverage float64
    HullStddev float64
    HullHistogram histogram.IntHistogram

    ShieldAverage float64
    ShieldStddev float64
    ShieldHistogram histogram.IntHistogram
}

// Maps ship name to statistics
type statsByShipName map[string]*stats
type resultsByShipName map[string]*simResult

func (st stats) ComputeStandardDeviations(res *simResult) {
    res.HitAverage = float64(st.HitSum) /  float64(st.Iterations)
    res.HitStddev = math.Sqrt((float64(st.HitSumSquares) / float64(st.Iterations)) - math.Pow(res.HitAverage, 2))

    res.CritAverage = float64(st.CritSum) /  float64(st.Iterations)
    res.CritStddev = math.Sqrt((float64(st.CritSumSquares) / float64(st.Iterations)) - math.Pow(res.CritAverage, 2))

    res.HullAverage = float64(st.HullSum) /  float64(st.Iterations)
    res.HullStddev = math.Sqrt((float64(st.HullSumSquares) / float64(st.Iterations)) - math.Pow(res.HullAverage, 2))

    res.ShieldAverage = float64(st.ShieldSum) /  float64(st.Iterations)
    res.ShieldStddev = math.Sqrt((float64(st.ShieldSumSquares) / float64(st.Iterations)) - math.Pow(res.ShieldAverage, 2))
}

func (this *stats) Add(that stats) *stats {
    this.HitSum += that.HitSum
    this.HitSumSquares += that.HitSumSquares
    this.CritSum += that.CritSum
    this.CritSumSquares += that.CritSumSquares
    this.HullSum += that.HullSum
    this.HullSumSquares += that.HullSumSquares
    this.ShieldSum += that.ShieldSum
    this.ShieldSumSquares += that.ShieldSumSquares
    return this
}

func (this *simResult) AddHistograms(that simResult) *simResult {
    this.HitHistogram.Add(that.HitHistogram)
    this.CritHistogram.Add(that.CritHistogram)
    this.HullHistogram.Add(that.HullHistogram)
    this.ShieldHistogram.Add(that.ShieldHistogram)
    return this
}

func (this *statsByShipName) Add(that statsByShipName) *statsByShipName {
    for name, thatStat := range(that) {
	thisStat, prs := (*this)[name]
	if !prs {
	    (*this)[name] = new(stats)
	    thisStat = (*this)[name]
	}
	thisStat.Add(*thatStat)
    }
    return this
}

func (this *resultsByShipName) AddHistograms(that resultsByShipName) *resultsByShipName {
    for name, thatResult := range(that) {
	thisResult, prs := (*this)[name]
	if !prs {
	    (*this)[name] = new(simResult)
	    thisResult = (*this)[name]
	}
	thisResult.AddHistograms(*thatResult)
    }
    return this
}
