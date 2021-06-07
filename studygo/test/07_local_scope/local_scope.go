package main

import "fmt"

var a = "G"

func main() {
	n()
	m()
	n()
}

func n() {
	// G
	println(a)
	fmt.Printf("global &a=%v \n", &a)
}

func m() {
	a := "O"
	// O
	println(a)
	fmt.Printf("local &a=%v \n", &a)
}
