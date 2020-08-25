package decorator

type Weapon interface {
	Damage() int
}

type Gun struct {
	base int
}

func NewGun(base int) *Gun {
	return &Gun{base: base}
}
func (g *Gun) Damage() int {
	return g.base
}

type Explosive struct {
	Weapon     // anonymous import
	additional int
}

func AddExplosive(w Weapon, additional int) *Explosive {
	return &Explosive{
		Weapon:     w,
		additional: additional,
	}
}

func (e *Explosive) Damage() int {
	return e.Weapon.Damage() + e.additional
}
