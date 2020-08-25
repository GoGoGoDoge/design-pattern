package flyweight

var factory *dressFlyweightFactory

func init() {
	// hungry man init
	factory = &dressFlyweightFactory{
		cached: make(map[string]*DressFlyweight),
	}
}

// DressFlyweightFactory can produce and obtain shared object.
type dressFlyweightFactory struct {
	cached map[string]*DressFlyweight
}

func Get(brand string) *DressFlyweight {
	var (
		dress *DressFlyweight
		exist bool
	)
	if dress, exist = factory.cached[brand]; !exist {
		dress = NewDressFlyweight(brand)
		factory.cached[brand] = dress
	}

	return dress
}

// DressFlyweight defines shared object class.
type DressFlyweight struct {
	brand string // unexported such that not changeable
}

func NewDressFlyweight(brand string) *DressFlyweight {
	return &DressFlyweight{
		brand: brand,
	}
}

// Folowing are usages example

type DressViewer struct {
	dress *DressFlyweight
}

func NewShopViewer(brand string) *DressViewer {
	return &DressViewer{
		dress: Get(brand),
	}
}

func (d *DressViewer) Equal(other *DressViewer) bool {
	return d.dress == other.dress
}
