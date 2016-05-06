package modification

type DamageDealer interface {
	IsDamageDealer() bool
}

type SecondaryWeapon interface {
	IsSecondaryWeapon() bool
}

type Transient interface {
	IsTransient() bool
}
