package bridge

type Driver interface {
	Run() string
}

type Computer interface {
	Do() string
}

type Macbook struct {
	driver Driver
	name   string
}

func (m *Macbook) Do() string {
	return m.name + ":" + m.driver.Run()
}

// we can wrap it like a driver use case
var macbookInstance *Macbook

func Register(driver Driver) {
	macbookInstance = &Macbook{
		name:   "Mac",
		driver: driver,
	}
}

func Action() string {
	return macbookInstance.Do()
}
