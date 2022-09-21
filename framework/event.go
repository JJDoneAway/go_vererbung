package framework

import "fmt"

type Event interface {
	A()
	B()
}

type MyEvent struct{}

func (e MyEvent) A() {
	fmt.Println("Method A was called")
}

func (e MyEvent) B() {
	fmt.Println("Method B was Called")
}
