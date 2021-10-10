package main

import (
	"fmt"
)

type IDuck interface {
	Quack()
	Walk()
}

func DuckDance(duck IDuck) {
	for i := 0; i < 3; i++ {
		duck.Quack()
		duck.Walk()
	}
}

type Bird struct {
}

func (b *Bird) Quack() {
	fmt.Printf("I am quacking! \n")
}

func (b *Bird) Walk() {
	fmt.Printf("I am walking! \n")
}

func main() {
	fmt.Printf("hello, duck dance \n")
	// fmt.Printf("val= %#v, T=%T \n", val, val)

	b := new(Bird)
	DuckDance(b)
}
