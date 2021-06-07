package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	var a int
	var b int32

	a = 15

	// 编译错误 cannot use a + a (type int) as type int32 in assignment
	// b = a + a

	// 因为 5 是常量，所以可以通过编译
	b = b + 5

	var n int16 = 34

	// 应该变量声明和下一行的赋值合并
	// should merge variable declaration with assignment on next line (S1021)
	/* var m int32

	// compiler error: cannot use n (type int16) as type int32 in assignment
	//m = n
	m = int32(n)
	*/

	var m int32 = int32(n)

	fmt.Printf("32 bit int is: %d \n", m)
	fmt.Printf("16 bit int is: %d \n", n)
	fmt.Printf("a: %d, b:%d \n", a, b)

	var c1 complex64 = 5 + 10i
	fmt.Printf("The value is: %v \n", c1)

	re := 3.333
	im := 5.555

	c := complex(re, im)

	fmt.Printf("c: %v \n", c)

	// 位运算
	fmt.Printf("1 & 1 = %v \n", 1&1)

	// 位左移 <<
	fmt.Printf("1 << 10 = %v \n", 1<<10)
	fmt.Printf("1 << 10 = %v \n", 1<<20)

	// 位左移常见实现存储单位的用例
	type ByteSize float64
	const (
		// 通过赋值给空白标识符来忽略值
		_           = iota
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)

	// 在通讯中使用位左移表示标识的用例
	type BitFlag int
	const (
		// 1 << 0 == 1
		Active BitFlag = 1 << iota
		// 1 << 1 == 2
		Send
		// 1 << 2 == 4
		Receive
	)

	flag := Active | Send
	fmt.Printf("Active | Send %v \n", flag)

	/* aa, bb := 10, 0
	// panic: runtime error: integer divide by zero
	cc := aa / bb
	fmt.Printf("aa / bb = %v \n", cc) */

	// 随机数
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("a=%v, &a=%v \n", a, &a)
	}

	for i := 0; i < 3; i++ {
		a := rand.Intn(8)
		fmt.Printf("a=%v, &a=%v \n", a, &a)
	}

	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)

	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / \n", 100*rand.Float32())
	}

	/* // 运算符与优先级
	有些运算符拥有较高的优先级，二元运算符的运算方向均是从左至右。下表列出了所有运算符以及它们的优先级，由上至下代表优先级由高到低：
	优先级     运算符
	7      ^ !
	6      * / % << >> & &^
	5      + - | ^
	4      == != < <= >= >
	3      <-
	2      &&
	1      ||
	*/

}

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x)

		if fraction >= 0.5 {
			whole++
		}

		return int(whole)
	}

	panic(fmt.Sprintf("%g is out of the int32 range", x))
}
