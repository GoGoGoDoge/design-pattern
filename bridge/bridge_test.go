package bridge_test

import (
	"bridge"
	"testing"
)

type DriverImpl struct {
}

func (d *DriverImpl) Run() string {
	return "Driver Impl"
}

func init() {
	// this is the driver trick
	bridge.Register(&DriverImpl{})
}

func TestBridge(t *testing.T) {
	result := bridge.Action()
	if result != "Mac:Driver Impl" {
		t.Fatal("Bridge is not correctly implemented:" + result)
	}
}
