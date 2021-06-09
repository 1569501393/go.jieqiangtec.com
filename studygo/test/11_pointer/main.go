package main

import "fmt"

func main() {
	var i1 = 5
	// An integer: 5, its location in memory: 0xc0000160a0
	fmt.Printf("An integer: %d, its location in memory: %p \n", i1, &i1)

	var intP *int

	intP = &i1

	// the value at memory location 0xc0000160a0 is 5
	fmt.Printf("the value at memory location %p is %d \n", intP, *intP)

	// 打印指针
	printStringPoint()
}

func printStringPoint() {
	s := "good bye"
	var p *string = &s
	*p = "ciao"

	fmt.Printf("Here is the pointer p: %p, \n", p)
	fmt.Printf("Here is the string *p: %s, \n", *p)
	fmt.Printf("Here is the string s: %s \n", s)

	// 获取常量和文字的地址异常
	// const i = 5
	// .\main.go:32:9: cannot take the address of i
	// ptr := &i

	// .\main.go:34:10: cannot take the address of 10
	// ptr2 := &10

	// 一个空指针的反向引用是不合法的，并且会使程序崩溃
	// .\main.go:32:9: cannot take the address of i
	// [signal 0xc0000005 code=0x1 addr=0x0 pc=0x7b738c]

	// var p0 *int = nil
	// *p0 = 0
}
