package combat


import (
	"testing"
	"github.com/geordanr/go_xwing/ship"
	"fmt"
	"github.com/geordanr/go_xwing/attack"
	"github.com/stretchr/testify/assert"
)

func TestListVersusList(t *testing.T) {
	// StealthFel
	listOne := make( []ModifiedShip, 1)

	fel := ModifiedShip{}
	fel.Ship = ship.ShipFactory["TIE Interceptor"]()
	fel.Name = "Soontir Fel"
	fel.FocusTokens = 2
	fel.EvadeTokens = 1
	fel.AttackerModifications = []attack.Modification{ attack.Modifications["Offensive Focus"] }
	fel.DefenderModifications = []attack.Modification{
		attack.Modifications["Defensive Focus"],
		attack.Modifications["Use Evade Token"],
		attack.Modifications["Stealth Device"],
		//attack.Modifications["Autothrusters"],
	}
	listOne[0] = fel

	// Accuracy-Corrected B-Wing.
	listTwo := make([]ModifiedShip, 2)
	bwing := ModifiedShip{}
	bwing.Ship = ship.ShipFactory["B-Wing"]()
	bwing.Name = fmt.Sprintf("B-Wing 1")
	bwing.FocusTokens = 1
	bwing.AttackerModifications = []attack.Modification{
		attack.Modifications["Offensive Focus"],
		attack.Modifications["Accuracy Corrector"],
	}

	listTwo[0] = bwing

	bwingTwo := ModifiedShip{}
	bwingTwo.Ship = ship.ShipFactory["B-Wing"]()
	bwingTwo.Name = fmt.Sprintf("B-Wing 2")
	bwingTwo.FocusTokens = 1
	bwingTwo.AttackerModifications = []attack.Modification{
		attack.Modifications["Offensive Focus"],
		attack.Modifications["Accuracy Corrector"],
	}

	attacks := ListVersusList( listOne, listTwo)
	if  len(attacks) != 3  {
		t.Errorf("should have received three attacks back from one on two versus")
	}

	if attacks[0].Attacker.Name != "Soontir Fel" {
		t.Errorf("Soontir shoots first")
	}

	assert.Equal( t, attacks[1].Attacker.Name, "B-Wing 1", "they should be equal")


}
