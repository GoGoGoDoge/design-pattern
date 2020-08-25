package observer_test

import "observer"

func ExampleProducer() {
	item := &observer.Item{}
	item.Register(&observer.Email{})
	item.Register(&observer.SMS{})
	item.Notify()
	// Output:
	// Email sent
	// Message sent
}
