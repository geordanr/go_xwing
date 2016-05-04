package stats

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAverage(t *testing.T) {
	assert := assert.New(t)

	s := new(IntStats)

	s.Update(1)
	s.Update(14)
	s.Update(10)
	s.Update(-4)

	assert.EqualValues(5.25, s.Average())
}

func TestStddev(t *testing.T) {
	assert := assert.New(t)

	s := new(IntStats)
	values := []int{2, 4, 4, 4, 5, 5, 7, 9}

	for _, val := range(values) {
		s.Update(val)
	}

	assert.EqualValues(2, s.Stddev())
}
