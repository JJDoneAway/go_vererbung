package handler

import (
	"fmt"
	"framework/framework"
)

type HandlerA struct{}

type funnyEvent struct {
	framework.Event
	counter int
}

func (f *funnyEvent) A() {
	fmt.Println("Funny Method A was called.")
	f.counter += 1
}

func (h HandlerA) Handle(f *framework.Framework, e framework.Event) {
	println("Handler A is fired")
	e.A()
	e.B()
	fun := funnyEvent{e, 0}
	// magic is in here as we just hand over the pointer to the struct,
	// but in the struct there is the real implementation
	f.Next(&fun)
	fmt.Printf("Funny Event Wrapper was called %d times\n‚", fun.counter)
}
