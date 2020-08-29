package interpreter_test

import (
	"interpreter"
	"testing"
)

func TestINterpreter(t *testing.T) {
	calculator := interpreter.NewExpressionInterpreter()
	result := calculator.Interpret("1 2 3 + -")
	if result != 0 {
		t.Fatal("Wrong interpreter implementation, got wrong output:", result)
	}
}
