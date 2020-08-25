package observer

import "fmt"

type Producer interface {
	Register(Consumer)
	Notify()
}

type Item struct {
	methods []Consumer
}

func (item *Item) Register(c Consumer) {
	item.methods = append(item.methods, c)
}

func (item *Item) Notify() {
	// item has changed status

	// notify
	for _, c := range item.methods {
		c.Action()
	}
}

// Consumer defines observer behavior
type Consumer interface {
	Action()
}

type Email struct {
}

func (e *Email) Action() {
	fmt.Println("Email sent")
}

type SMS struct {
}

func (sms *SMS) Action() {
	fmt.Println("Message sent")
}
