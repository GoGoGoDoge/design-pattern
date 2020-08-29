package state

type Week interface {
	Today() string
	Next(*DayContext)
}

type DayContext struct {
	today Week
}

func NewDayContext() *DayContext {
	return &DayContext{
		today: &Sunday{},
	}
}

func (ctx *DayContext) Today() string {
	return ctx.today.Today()
}

func (ctx *DayContext) Next() {
	ctx.today.Next(ctx)
}

type Sunday struct {
	ctx *DayContext
}

func (day *Sunday) Today() string {
	return "Sunday"
}

func (day *Sunday) Next(ctx *DayContext) {
	// This can be further optimized to use singleton to avoid new object
	ctx.today = &Monday{}
}

type Monday struct {
	ctx *DayContext
}

func (day *Monday) Today() string {
	return "Monday"
}

func (day *Monday) Next(ctx *DayContext) {
	ctx.today = &Sunday{}
}
