package memento

import (
	"fmt"
)

type GameState interface {
	SetState(s string)
	GetState() string
	GetSnapshot() Snapshot
	RestoreFromSnapshot(snapshot Snapshot)
}

type Snapshot interface {
	GetState() string
	PersistToDisk()
}

type SuperMario struct {
	state string
}

func (mario *SuperMario) SetState(s string) {
	mario.state = s
}

func (mario *SuperMario) GetState() string {
	return mario.state
}

func (mario *SuperMario) GetSnapshot() Snapshot {
	return &GameSnapshot{
		m: mario.state,
	}
}

func (mario *SuperMario) RestoreFromSnapshot(snapshot Snapshot) {
	mario.state = snapshot.GetState()
}

type GameSnapshot struct {
	m string
}

func (g *GameSnapshot) GetState() string {
	return g.m
}

func (g *GameSnapshot) PersistToDisk() {
	fmt.Println("Saving to disk...")
}
