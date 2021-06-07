package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
)

// 常量
const C = "C"
const PI = 3.141592654
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

const (
	a = iota
	b = iota
	c = iota
)

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Staturday
)

// 变量
var v int = 5

// 简短型只能在局部定义
// abc:="abc"

var abc string = "abc"

type T struct {
}

var Pi float64

func init() {
	// init() function computes Pi
	Pi = 4 * math.Atan(1)

	// init 函数也经常被用在当一个程序开始之前调用后台执行的 goroutine，如下面这个例子当中的 backend()：
	// go backend()
}

func main() {
	fmt.Println("template")
	// var a int
	Func1()
	fmt.Println(a)
	fmt.Println(a + v)

	fmt.Printf("%d+%d=%d \n", a, v, a+v)
	fmt.Printf("%d+%d+%d=%d \n", a, b, c, a+b+c)
	fmt.Println("aaaaa")
	// fmt.Println(a + v)

	fmt.Printf("PI=%v\n", PI)
	fmt.Printf("a=%v, Monday=%v, Tuesday=%v\n", a, Monday, Tuesday)

	// 调用函数
	variables()

	// 注意大小写
	fmt.Printf("global var Pi=%v \n", Pi)
	fmt.Printf("global const PI=%v \n", PI)

}

func (t T) Method1() {
}

func Func1() {
}

func variables() {
	// 指针
	var a, b *int
	fmt.Printf("a,b position is %v,%v\n", &a, &b)
	/* a = 1
	b = 2
	fmt.Printf("a,b position is %v,%v", &a, &b) */

	HOME := os.Getenv("HOME")
	USER := os.Getenv("USER")
	GOROOT := os.Getenv("GOROOT")
	fmt.Printf("HOME=%s, USER=%s, GOROOT=%s \n", HOME, USER, GOROOT)

	/* // 获取环境变量
	for k, v := range os.Environ() {
		fmt.Printf("k=%v, v=%v \n", k, v)
	} */

	var goos string = runtime.GOOS
	fmt.Printf("the os is:%s, &goos=%v \n", goos, &goos)

	goos = "abc"
	fmt.Printf("the os is:%s, &goos=%v \n", goos, &goos)

	/* path := os.Getenv("PATH")
	fmt.Printf("Path is %s \n", path) */

	fmt.Print("Hello:", 23, "\n")
	fmt.Print(0+23, "\n")

	fmt.Printf("global abc=%v \n", abc)

	// .\main.go:113:6: aaa declared but not used
	// var aaa string = "aaa"

}
