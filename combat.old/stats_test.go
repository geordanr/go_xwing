package combat

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestStatsAdd(t *testing.T) {
    assert := assert.New(t)

    one := stats{
	Iterations: 1,
	HitSum: 2,
	HitSumSquares: 3,
	CritSum: 4,
	CritSumSquares: 5,
	HullSum: 6,
	HullSumSquares: 7,
	ShieldSum: 8,
	ShieldSumSquares: 9,
    }

    two := stats{
	Iterations: 10,
	HitSum: 20,
	HitSumSquares: 30,
	CritSum: 40,
	CritSumSquares: 50,
	HullSum: 60,
	HullSumSquares: 70,
	ShieldSum: 80,
	ShieldSumSquares: 90,
    }

    one.Add(two)

    assert.EqualValues(11, one.Iterations)
    assert.EqualValues(22, one.HitSum)
    assert.EqualValues(33, one.HitSumSquares)
    assert.EqualValues(44, one.CritSum)
    assert.EqualValues(55, one.CritSumSquares)
    assert.EqualValues(66, one.HullSum)
    assert.EqualValues(77, one.HullSumSquares)
    assert.EqualValues(88, one.ShieldSum)
    assert.EqualValues(99, one.ShieldSumSquares)
}
