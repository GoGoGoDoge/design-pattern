package factorymethod_test

import (
	"factory/factorymethod"
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	ak47Factory := factorymethod.NewGunFactory("ak47")
	maverickFactory := factorymethod.NewGunFactory("maverick")
	ak47 := ak47Factory.NewGun()
	maverick := maverickFactory.NewGun()

	if ak47.Fire() != "ak47 Fire" {
		t.Fatal("Incorrect ak47")
	}

	if maverick.Fire() != "maverick Fire" {
		t.Fatal("Incorrect maverick")
	}
}
