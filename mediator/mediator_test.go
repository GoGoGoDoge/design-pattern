package mediator_test

import (
	"mediator"
	"testing"
)

func TestMediator(t *testing.T) {
	m := &mediator.Mediator{}
	planeA := &mediator.PlaneA{}
	planeB := &mediator.PlaneB{}
	planeA.Connect(m)
	planeB.Connect(m)
	if planeA.Land() != "Plane A is landing" {
		t.Fatal("Mediator does not handle Plane A correctly")
	}
	if planeB.Land() != "Plane B cannot land" {
		t.Fatal("Mediator does not handle Plane B correctly")
	}
}
