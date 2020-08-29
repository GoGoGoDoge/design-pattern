package memento_test

import (
	"memento"
	"testing"
)

func TestMemento(t *testing.T) {
	mario := &memento.SuperMario{}
	mario.SetState("Init")
	snapshot := mario.GetSnapshot()
	mario.SetState("Fail")
	if mario.GetState() != "Fail" {
		t.Fatal("Incorrect set state")
	}

	mario.RestoreFromSnapshot(snapshot)
	if mario.GetState() != "Init" {
		t.Fatal("Incorrect restore to init state")
	}
}
