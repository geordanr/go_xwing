package runner

import (
	// "github.com/geordanr/go_xwing/ship"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShipsFromJSON(t *testing.T) {
	assert := assert.New(t)

	jsonData := `{
		"ships": [
			{
				"name": "X-Wing",
				"attack": 3,
				"agility": 2,
				"hull": 3,
				"shields": 2
			},
			{
				"name": "TIE Fighter",
				"attack": 2,
				"agility": 3,
				"hull": 3,
				"shields": 0
			}
		]
	}`

	factoryMap, err := shipsFromJSON([]byte(jsonData))
	assert.Nil(err)
	assert.Equal(2, len(factoryMap))
	assert.Contains(factoryMap, "X-Wing")
	assert.Contains(factoryMap, "TIE Fighter")

	xwingFactory := factoryMap["X-Wing"]
	xwing := xwingFactory("Wedge Antilles")
	assert.Equal("Wedge Antilles", xwing.Name())
	assert.EqualValues(3, xwing.Attack())
	assert.EqualValues(2, xwing.Agility())
	assert.EqualValues(3, xwing.Hull())
	assert.EqualValues(2, xwing.Shields())

	tieFactory := factoryMap["TIE Fighter"]
	tie := tieFactory("Howlrunner")
	assert.Equal("Howlrunner", tie.Name())
	assert.EqualValues(2, tie.Attack())
	assert.EqualValues(3, tie.Agility())
	assert.EqualValues(3, tie.Hull())
	assert.EqualValues(0, tie.Shields())

}
