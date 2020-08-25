package flyweight_test

import (
	"flyweight"
	"testing"
)

func TestFlyweight(t *testing.T) {
	viewer1 := flyweight.NewShopViewer("GoGoDoge")
	viewer2 := flyweight.NewShopViewer("GoGoDoge")
	if !viewer1.Equal(viewer2) {
		t.Fatal("Not the same object!")
	}
}
