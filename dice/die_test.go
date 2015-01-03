package dice

import "testing"

func TestRollable(t *testing.T) {
    var attackDie AttackDie
    var defenseDie DefenseDie
    attackDie.Roll()
    defenseDie.Roll()
}
