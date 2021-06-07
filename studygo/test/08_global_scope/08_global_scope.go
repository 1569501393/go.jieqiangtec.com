package main

import "fmt"

var a = "G"

func main() {
	// G
	n()
	// O
	m()
	// O
	n()
}

func n() {
	// G
	fmt.Printf("a=%v, &a=%v \n", a, &a)
}

func m() {
	a = "O"
	// O
	fmt.Printf("a=%v, &a=%v \n", a, &a)
}
