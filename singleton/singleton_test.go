package singleton_test

import (
	"testing"

	"singleton"
)

func TestSingleton(t *testing.T) {
	ins1 := singleton.GetInstance()
	ins2 := singleton.GetInstance()

	if ins1 != ins2 {
		t.Fatal("Instance is not equal")
	}
}

func TestMutipleLazyInit(t *testing.T) {
	for i := 0; i < 100; i++ {
		go singleton.GetLazyInstance()
	}
	if singleton.LazyCount != 1 {
		t.Fatal("Somewhere is messed up with lazy singleton init")
	}
}

func TestLazySingleton(t *testing.T) {
	ins1 := singleton.GetLazyInstance()
	ins2 := singleton.GetLazyInstance()
	if ins1 != ins2 {
		t.Fatal("Lazy Instance is not equal")
	}
}
