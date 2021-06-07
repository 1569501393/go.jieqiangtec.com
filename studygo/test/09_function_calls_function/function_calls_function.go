package main

import "fmt"

var a string

func main() {
	a = "G"

	// a=G, &a=0x17bc50
	fmt.Printf("a=%v, &a=%v \n", a, &a)

	// O G
	f1()
}

func f1() {
	a := "O"
	// a=O, &a=0xc000050200
	fmt.Printf("a=%v, &a=%v \n", a, &a)
	// a=G, &a=0x17bc50
	f2()
}

func f2() {
	// a=G, &a=0x17bc50
	fmt.Printf("a=%v, &a=%v \n", a, &a)
}
