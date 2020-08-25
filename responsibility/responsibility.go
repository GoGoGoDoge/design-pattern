package responsibility

import "fmt"

type IHandler interface {
	handle(int) bool
}

type Chain struct {
	handlers     []IHandler
	handlerFuncs []func(int) bool
}

func (chain *Chain) AddHandler(handler IHandler) {
	chain.handlers = append(chain.handlers, handler)
}
func (chain *Chain) AddHanderFunc(f func(int) bool) {
	chain.handlerFuncs = append(chain.handlerFuncs, f)
}
func (chain *Chain) HandleByFunc(count int) {
	for _, f := range chain.handlerFuncs {
		if f(count) {
			break
		}
	}
}
func (chain *Chain) Handle(count int) {
	for _, handler := range chain.handlers {
		if handler.handle(count) {
			break
		}
	}
}

type AliceHandler struct{}

func (hanlder *AliceHandler) handle(count int) bool {
	if count < 10 {
		fmt.Println("Intercepted by Alice")
		return true
	}
	return false
}

type BobHandler struct {
}

func (handler *BobHandler) handle(count int) bool {
	if count >= 10 {
		fmt.Println("Intercepted by Bob")
		return true
	}
	return false
}
