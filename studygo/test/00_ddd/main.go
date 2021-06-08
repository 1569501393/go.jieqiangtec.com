package main

import "fmt"

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
}

// 循环
func forRange() {
	x := []string{"a", "b", "c"}

	for idx, v := range x {
		fmt.Printf("idx=%v, v=%v, &v=%v \n", idx, v, &v)
	}
}
