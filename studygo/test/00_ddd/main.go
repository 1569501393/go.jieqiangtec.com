package main

import "fmt"
import "rsc.io/quote"

func main() {

	  fmt.Println("hello, world")

	x := 1

	fmt.Println(x)

	{
		 fmt.Println(x)
		x := 2
		fmt.Println(x)
	}

	fmt.Println(x)

	forRange()

	printVar()

	printGo()
}

// 循环
func forRange() {
	x := []string{"a", "b", "c"}

	for idx, v := range x {
		fmt.Printf("idx=%v, v=%v, &v=%v \n", idx, v, &v)
	}
}

// 打印变量
func printVar() {
	a := 1
	b := 2
	c := a + b
	fmt.Printf("a+b=%v+%v=%v \n", a, b, c)
}

func printGo() {
	fmt.Println(quote.Go())

}
