package adapter

type Target interface {
	DesiredFunc() (string, error)
}

// Adaptee wants to convert Adaptee to our target
type Adapter struct {
	Adaptee
}

func NewAdapter(a Adaptee) *Adapter {
	return &Adapter{
		Adaptee: a,
	}
}

func (a *Adapter) DesiredFunc() (string, error) {
	// we can access CustomizedFunc because of anoymous composite.
	return a.CustomizedFunc(), nil
}

// Adaptee is from external source
type Adaptee interface {
	CustomizedFunc() string
}

func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

type adapteeImpl struct{}

func (a *adapteeImpl) CustomizedFunc() string {
	return "customized"
}
