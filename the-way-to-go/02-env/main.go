package main

import "C"
import (
	"fmt"
	"runtime"
)

func main() {
	println("Hello", "World")
	// 打印版本
	fmt.Printf("version:%s", runtime.Version())

	// var i int
	// C.unit(i)
	// C.random()
}
