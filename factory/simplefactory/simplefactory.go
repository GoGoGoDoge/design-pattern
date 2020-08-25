package simplefactory

type IGun interface {
	Fire() string
}

type Gun struct {
	name string
}

func NewGun(name string) IGun {
	if name == "ak47" {
		return newAk47()
	}

	if name == "maverick" {
		return newMaverick()
	}

	return nil
}

func (g *Gun) Fire() string {
	return g.name + " Fire"
}

type ak47 struct {
	Gun
}

func newAk47() IGun {
	return &ak47{
		Gun: Gun{
			name: "ak47",
		},
	}
}

type maverick struct {
	Gun
}

func newMaverick() IGun {
	return &maverick{
		Gun: Gun{
			name: "maverick",
		},
	}
}
