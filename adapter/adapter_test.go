package adapter_test

import (
	"adapter"
	"testing"
)

func TestAdapter(t *testing.T) {
	adapterObj := adapter.NewAdapter(adapter.NewAdaptee())
	output, err := adapterObj.DesiredFunc()
	if err != nil || output != "customized" {
		t.Fatal("Adapter is implemented incorrectly")
	}
}
