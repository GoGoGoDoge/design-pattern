package decorator_test

import (
	"decorator"
	"testing"
)

func TestDecorator(t *testing.T) {
	gun := decorator.NewGun(20)
	if decorator.AddExplosive(gun, 5).Damage() != 25 {
		t.Fatal("Decorator is not implemented correctly")
	}
}
