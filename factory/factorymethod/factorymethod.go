package factorymethod

type IGun interface {
	Fire() string
}

// IGunFactory abstracts out Gun creator. 这个是主要区别
type IGunFactory interface {
	NewGun() IGun
}

var cacheFactories map[string]IGunFactory

func init() {
	cacheFactories = make(map[string]IGunFactory)
	cacheFactories["ak47"] = &ak47Factory{}
	cacheFactories["maverick"] = &maverickFactory{}
}

func NewGunFactory(name string) IGunFactory {
	return cacheFactories[name]
}

type Gun struct {
	name string
}

// we don't need NewGun anymore

func (g *Gun) Fire() string {
	return g.name + " Fire"
}

type ak47 struct {
	Gun
}

type ak47Factory struct {
}

func (af *ak47Factory) NewGun() IGun {
	return &ak47{
		Gun: Gun{
			name: "ak47",
		},
	}
}

type maverick struct {
	Gun
}

type maverickFactory struct{}

func (mf *maverickFactory) NewGun() IGun {
	return &maverick{
		Gun: Gun{
			name: "maverick",
		},
	}
}
