package serialization

import (
	// "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var shipJson = `{
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

func TestShipsFromJSON(t *testing.T) {
	assert := assert.New(t)

	factoryMap, err := shipsFromJSON([]byte(shipJson))
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

func TestFromJSON(t *testing.T) {
	assert := assert.New(t)

	factoryMap, err := shipsFromJSON([]byte(shipJson))

	paramsJson := ` {
		"iterations": 1,
		"combatants": [
			{
				"name": "Luke Skywalker",
				"ship": "X-Wing",
				"skill": 8,
				"initiative": true,
				"tokens": {
					"focus": 1,
					"targetlock": "Colonel Vessery"
				}
			},
			{
				"name": "Howlrunner",
				"ship": "TIE Fighter",
				"skill": 8,
				"initiative": false,
				"tokens": {
					"focus": 1,
					"evade": 1
				}
			}
		],
		"attack_queue": [
			{
				"attacker": "Luke Skywalker",
				"defender": "Howlrunner",
				"mods": {
					"Modify Attack Dice": [
						["attacker", "Spend Focus Token"]
					],
					"Modify Defense Dice": [
						["defender", "Spend Focus Token"],
						["defender", "Spend Evade Token"]
					]
				}
			},
			{
				"attacker": "Howlrunner",
				"defender": "Luke Skywalker",
				"mods": {
					"Modify Attack Dice": [
						["attacker", "Spend Focus Token"]
					],
					"Modify Defense Dice": [
						["defender", "Spend Focus Token"]
					]
				}
			}
		]
	}`

	output, err := FromJSON([]byte(paramsJson), factoryMap)
	assert.Nil(err)
	for {
		// fmt.Println("reading from output...")
		_, more := <-output
		// fmt.Println("read, more=", more)
		if !more {
			break
		}
	}
}
