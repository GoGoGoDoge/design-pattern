package builder_test

import (
	"builder"
	"testing"
)

func TestBuilder(t *testing.T) {
	sqBuilder := builder.SquareBuilder{}
	_, err := sqBuilder.SetHeight(10).SetWidth(9).BuildSquare()
	if err == nil {
		t.Fatal("Expecting building error but not get it")
	}

	perfectSquare, err := sqBuilder.SetHeight(20).SetWidth(20).BuildSquare()
	if err != nil {
		t.Fatal("Does not expect build error")
	}
	if !perfectSquare.Validate() {
		t.Fatal("Incorrect square is built")
	}
}
