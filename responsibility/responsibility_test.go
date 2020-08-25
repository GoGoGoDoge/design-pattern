package responsibility_test

import (
	"fmt"
	"responsibility"
)

func ExampleChain_AddHandler() {
	chain := &responsibility.Chain{}
	chain.AddHandler(&responsibility.AliceHandler{})
	chain.AddHandler(&responsibility.BobHandler{})
	chain.Handle(1)
	// output:
	// Intercepted by Alice
}

func ExampleChain_AddHanderFunc() {
	chain := &responsibility.Chain{}
	chain.AddHanderFunc(func(count int) bool {
		fmt.Println("Custom func")
		return true
	})
	chain.HandleByFunc(1)
	// output:
	// Custom func
}
