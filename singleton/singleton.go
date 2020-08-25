package singleton

import (
	"fmt"
	"sync"
)

// singleton must be lowercase to restrict access
type singleton struct{}

var instances map[int]*singleton
var limitedNumber int = 3

func init() {
	instances = make(map[int]*singleton)
	for i := 0; i < limitedNumber; i++ {
		instances[i] = &singleton{}
	}
}

func GetInstance() *singleton {
	return GetInstanceById(0)
}

func GetInstanceById(id int) *singleton {
	if id >= limitedNumber {
		return nil
	}
	return instances[id]
}

func (s *singleton) DoSomething() {
	fmt.Println("do something...")
}

var lock = &sync.Mutex{}

type lazysingleton struct{}

func (s *lazysingleton) DoSomething() {
	fmt.Println("do something lazy...")
}

var lazySingletonInstance *lazysingleton
var LazyCount int

func GetLazyInstance() *lazysingleton {
	if lazySingletonInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		// double check here
		if lazySingletonInstance == nil {
			LazyCount++
			// fmt.Println("This shall only be printed once")
			lazySingletonInstance = &lazysingleton{}
		} else {
			// fmt.Println("Lazy singleton is already created - 1")
		}
	} else {
		// fmt.Println("Lazy singleton is already created - 2")
	}
	return lazySingletonInstance
}
