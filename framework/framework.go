package framework

var (
	cache []Handler
	index int
)

type Framework struct{}

/*
 * You have to implement the handler interface
 */
type Handler interface {
	Handle(f *Framework, e Event)
}

func NewFramework() Framework {
	return Framework{}
}

func (f Framework) AddHandler(h Handler) {
	cache = append(cache, h)
}

func (f Framework) FireEvent() {
	index = 0
	// we ignore for this example that the cache could be empty
	cache[0].Handle(&f, MyEvent{})
}

func (f Framework) Next(e Event) {
	index += 1
	if len(cache) > index {
		cache[index].Handle(&f, e)
	} else {
		println("All Handlers are fired")
	}
}
