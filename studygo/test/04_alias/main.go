package main

// Package implementing formatted I/O
import fm "fmt"

func main() {
	fm.Println("hello, world")
	fm.Println("hello, zhangsan")
	fm.Println("hello, lisi")
	fm.Print("hello, zhaoliu \n\n")
	fm.Print("hello, " + "jieqiang")
}

func Sum(a, b int) int {
	return a + b
}
