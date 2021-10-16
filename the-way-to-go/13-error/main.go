package main

import (
	"errors"
	"fmt"
	calc "gin111/the-way-to-go/14-unit-test"
	"os"
)

// var errNotFound error = errors.New("Not found error")
var errNotFound = errors.New("Not found error ")

func main() {
	// 调用calc包中的Add方法
	sum := calc.Add(1, 2)
	fmt.Printf("sum=%#v \n", sum)
	fmt.Printf("Hello io \n")
	a := 2
	fmt.Printf("a val= %#v, T=%T \n", a, a)

	fmt.Printf("error: %v \n", errNotFound)

	// println("Starting the program")
	// panic("a server error occurred: stopping the program")
	// println("Ending the program")

	var user = os.Getenv("USER")
	fmt.Printf("user val= %#v, T=%T \n", user, user)

	// slice
	s := make([]int, 3, 10)
	b := s[0]
	fmt.Printf("s=%#v, b = %#v \n", s, b)

	myMap := make(map[string]string)
	myMap["zhangSan"] = "1"
	myMap["liSi"] = "2"
	myMap["wangWu"] = "3"
	fmt.Printf("myMap=%#v \n", myMap)

	// reflect
	var aa string
	aa = "abcd"

	var allType interface{}
	allType = aa

	str, _ := allType.(string)
	fmt.Printf("str = %#v \n", str)
	fmt.Println(str)
}

func add(num1, num2 int) int {
	return num1 + num2
}
