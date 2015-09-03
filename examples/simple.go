package main

import (
	"fmt"

	"github.com/choix/cake"
)

func layerA(next cake.HandlerFunc) cake.HandlerFunc {
	return func(env interface{}) {
		fmt.Println("+ A")
		next(env)
		fmt.Println("- A")
	}
}

func layerB(next cake.HandlerFunc) cake.HandlerFunc {
	return func(env interface{}) {
		fmt.Println("+ B")
		next(env)
		fmt.Println("- B")
	}
}

func main() {
	m := cake.New()
	m.Use(layerA)
	m.Use(layerB)
	m.Call(nil)
}
