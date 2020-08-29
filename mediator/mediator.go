package mediator

type Aircraft interface {
	Connect(mediator *Mediator)
}

type Mediator struct{}

func (m *Mediator) canLand(aircraft Aircraft) bool {
	switch aircraft.(type) {
	case *PlaneA:
		return true
	case *PlaneB:
		return false
	}
	return false
}

type PlaneA struct {
	mediator *Mediator
}

func (p *PlaneA) Connect(mediator *Mediator) {
	p.mediator = mediator
}

func (p *PlaneA) Land() string {
	// need to pass current object to mediator
	if p.mediator.canLand(p) {
		return "Plane A is landing"
	}
	return "Plane A cannot land"
}

type PlaneB struct {
	mediator *Mediator
}

func (p *PlaneB) Connect(mediator *Mediator) {
	p.mediator = mediator
}

func (p *PlaneB) Land() string {
	if p.mediator.canLand(p) {
		return "Plane B is landing"
	}

	return "Plane B cannot land"
}
