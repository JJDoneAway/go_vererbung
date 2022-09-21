package main

import (
	"fmt"
	"framework/framework"
	"framework/handler"
)

func main() {
	fmt.Println("Ramp up new Framework")
	f := framework.NewFramework()
	fmt.Println("--------------")
	fmt.Println("Added Handler A")
	f.AddHandler(handler.HandlerA{})
	fmt.Println("Added Handler B")
	f.AddHandler(handler.HandlerB{})
	fmt.Println("--------------")
	fmt.Println("Fire Event now")
	f.FireEvent()
}
