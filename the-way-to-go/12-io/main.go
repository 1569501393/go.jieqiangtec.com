package main

import (
	"fmt"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12/ 5212/ Go"
	format                 = "%f/ %d/ %s"
	a                      = 1
)

func main() {
	fmt.Printf("Hello io \n")
	// a := 2
	fmt.Printf("a val= %#v, T=%T \n", a, a)

	// println("Please enter your full name:")
	fmt.Println("Please enter your full name:")

	fmt.Scanln(&firstName, &lastName)

	// fmt.Printf("Hi %s %s ! \n", firstName, lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName)
	fmt.Sscanf(input, format, &f, &i, &s)
	// println("From the string we read: ", f, i, s)
	fmt.Println("From the string we read: ", f, i, s)

}
