package main

import (
	"errors"
	"fmt"
	"os"
)

// var errNotFound error = errors.New("Not found error")
var errNotFound = errors.New("Not found error ")

func main() {
	fmt.Printf("Hello io \n")
	a := 2
	fmt.Printf("a val= %#v, T=%T \n", a, a)

	fmt.Printf("error: %v \n", errNotFound)

	// println("Starting the program")
	// panic("a server error occurred: stopping the program")
	// println("Ending the program")

	var user = os.Getenv("USER")
	fmt.Printf("user val= %#v, T=%T \n", user, user)

}
