package main

import (
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

const LIM = 41

var fibs [LIM]int
var fibs2 [LIM]int

func main() {

	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}

	println("In main before calling greeting")
	greeting()
	where()

	println("In main after calling greeting")
	sayHello("aaa")
	MultiPly3Nums(2, 3, 5)
	greetingTowho("hello, ", "zs", "ls")

	x := min(1, 3, 2, 0)
	fmt.Printf("The min is: %d \n", x)

	slice := []int{7, 1, 2, 3, 4, 5, 9}
	x = min(slice...)
	fmt.Printf("The min in the slice is: %d, %T \n", x, slice)

	// In function1 at the top
	// In function1 at the bottom
	// function2: Deferred until the end of the calling function
	function1()

	deferTest()
	doDbOperations()
	deferLog("Go")

	result := 0
	// result2 := 0
	start := time.Now()

	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)

		// result2 = fibonacciOptimize(i)
		// fmt.Printf("fibonacciOptimize(%d) is: %d\n", i, result2)
	}

	anonymousTest()

	p2 := Add2()
	fmt.Printf("Call Add2 for 3 givs: %v \n", p2(3))

	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))

	where()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)

	// var result int = 0
	start2 := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci2(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end2 := time.Now()
	delta2 := end2.Sub(start2)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta2)
}

func greeting() {
	println("In greeting:Hi!!!")
}

func sayHello(object string) {
	fmt.Printf("%v say Hello \n", object)
}

func MultiPly3Nums(a, b, c int) {
	fmt.Printf("a*b*c = %d*%d*%d=%v \n", a, b, c, a*b*c)
}

func greetingTowho(prefix string, who ...string) {
	fmt.Printf("%v, %v \n", prefix, who)
}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}

	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}

	return min
}

func function1() {
	fmt.Printf("In function1 at the top \n")
	defer function2()
	fmt.Printf("In function1 at the bottom \n")
}

func function2() {
	fmt.Printf("function2: Deferred until the end of the calling function \n")
}

// 关键字 defer 允许我们进行一些函数执行完成后的收尾工作
func deferTest() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("i: %v \n", i)
	}
}

func connectToDb() {
	println("ok, connected to db")
}

func disconnectFromDb() {
	println("ok, disconnect from db")
}

func doDbOperations() {
	defer disconnectFromDb()
	connectToDb()
	fmt.Println("Deferring the database disconnect.")
	fmt.Println("Doing some Db operations")
	fmt.Println("Oops! some crash or network error ...")
	fmt.Println("Returning form function here!")
}

func deferLog(s string) (n int, err error) {
	defer func() {
		log.Printf("deferLog(%q) = %d, %v \n", s, n, err)
	}()

	return 7, io.EOF
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}

	return
}

func fibonacciOptimize(n int) (res int) {
	// memoization: check if fibonacci(n) is already known in array:
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}

	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}

	fibs[n] = res

	return
}

func fibonacci2(n int) (res int) {
	// memoization: check if fibonacci(n) is already known in array:
	if fibs2[n] != 0 {
		res = fibs2[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci2(n-1) + fibonacci2(n-2)
	}
	fibs2[n] = res
	return
}

func anonymousTest() {
	for i := 0; i < 5; i++ {
		g := func(i int) { fmt.Printf("%d", i) }
		g(i)
		fmt.Printf(" - g is of type %T and has value %v \n", g, g)
	}
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		fmt.Printf("Adder a=%v, b=%v \n", a, b)
		return a + b
	}
}
