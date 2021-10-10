package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello io \n")
	a := 2
	fmt.Printf("a val= %#v, T=%T \n", a, a)

	for i := 0; i < 5; i++ {
		fmt.Printf("i v:%v \n", i)
	}

}
