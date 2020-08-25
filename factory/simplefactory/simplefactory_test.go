package simplefactory_test

import (
	"factory/simplefactory"
	"testing"
)

func TestSimpleFactory(t *testing.T) {
	ak47 := simplefactory.NewGun("ak47")
	maverick := simplefactory.NewGun("maverick")

	if ak47.Fire() != "ak47 Fire" {
		t.Fatal("Incorrect ak47")
	}

	if maverick.Fire() != "maverick Fire" {
		t.Fatal("Incorrect maverick")
	}
}
