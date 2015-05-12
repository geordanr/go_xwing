package attack

type Modification interface {
    Modify(*Attack) *Attack
    String() string
    ModifiesAttackResults() bool
    ModifiesDefenseResults() bool
}

var Modifications map[string]Modification = map[string]Modification{
    "Offensive Focus": new(offensiveFocus),
    "Target Lock": new(targetLock),
    "Howlrunner": offensiveReroll{name: "Howlrunner", numToReroll: 1},
    "Predator (low PS)": offensiveReroll{name: "Predator (low PS)", numToReroll: 2},
    "Predator (high PS)": offensiveReroll{name: "Predator (high PS)", numToReroll: 1},
    "Marksmanship": new(marksmanship),
    "Chiraneau": new(chiraneau),
    "Han Solo": new(hanSolo),
    "Heavy Laser Cannon": new(heavyLaserCannon),
    "Accuracy Corrector": new(accuracyCorrector),

    "Defensive Focus": new(defensiveFocus),
    "Use Evade Token": new(useEvadeToken),
    "C-3PO (guess 0)": c3po{Guess: 0},
    "C-3PO (guess 1)": c3po{Guess: 1},
    "C-3PO (guess 2)": c3po{Guess: 2},
    "C-3PO (guess 3)": c3po{Guess: 3},
}
