package handler

import "framework/framework"

type HandlerB struct{}

func (h HandlerB) Handle(f *framework.Framework, e framework.Event) {
	println("Handler B is fired")
	e.A()
	e.B()
	f.Next(e)
}
