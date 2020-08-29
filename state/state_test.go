package state_test

import (
	"state"
	"testing"
)

func TestState(t *testing.T) {
	dayContext := state.NewDayContext()
	if dayContext.Today() != "Sunday" {
		t.Fatal("Init error")
	}

	dayContext.Next()
	if dayContext.Today() != "Monday" {
		t.Fatal("State transfer error, expecting Monday")
	}

	dayContext.Next()
	if dayContext.Today() != "Sunday" {
		t.Fatal("State transfer error, expecting Sunday")
	}
}
