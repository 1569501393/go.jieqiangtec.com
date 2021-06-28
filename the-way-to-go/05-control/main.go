package main

import (
	"fmt"
	"runtime"
	"strconv"
)

var prompt = "Enter a digit, e.g. 3 OR %s to quit."

// 初始入口
func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else {
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}

	fmt.Printf("the os is %s, prompt eq %v \n", runtime.GOOS, prompt)
}

// 主函数
func main() {
	// if判断bool值
	testBooleans()

	/* var i int
	i = testReturn() */

	// i := testReturn()
	var i = testReturn()
	fmt.Printf("i=%d \n", i)

	// switch
	testSwitch()

	testFor()

	testFor2()
}

// 判断bool值
func testBooleans() {
	bool1 := true
	if bool1 {
		fmt.Printf("the value is true=%v \n", bool1)
	} else {
		fmt.Printf("the value is false=%v \n", bool1)
	}
}

// 测试返回值
func testReturn() int {
	bool1 := true
	if bool1 {
		return 1
	}
	return 2
}

// 测试绝对值
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 判断更大值
func isGreater(x, y int) bool {
	/* if x > y {
		return true
	}
	return false */
	return x > y
}

// ifelse模式
func ifelse() {
	var first int = 10
	var cond int
	if first <= 0 {
		fmt.Printf("first is less than or equal to 0 \n")
	} else if first > 0 && first < 5 {
		fmt.Printf("first is between 0 and 5 \n")
	} else {
		fmt.Printf("first is 5 OR greater \n")
	}

	if cond = 5; cond > 10 {
		fmt.Printf("cond is greater than 10 \n")
	} else {
		fmt.Printf("cond is not greater than 10 \n")
	}
}

// 字符串转化
func stringConversion() {
	var orig string = "ABC"
	var newS string

	fmt.Printf("The size of ints is: %d \n", strconv.IntSize)

	an, err := strconv.Atoi(orig)

	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error \n", orig)
		return
	}

	fmt.Printf("the integer is %d \n", an)

	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("the new string is %s \n", newS)

}

// switch语法
func testSwitch() {
	var num1 int = 100

	switch num1 {
	case 98, 99:
		fmt.Println("It's equal to 98")
	case 100:
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's not equal to 98 or 100")
	}

	// was <= 6
	// was <= 7
	// was <= 8
	// default case
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}

}

// This is the 0 iteration
// This is the 1 iteration
// This is the 2 iteration
// This is the 3 iteration
// This is the 4 iteration
func testFor() {
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration \n", i)
	}
}

func testFor2() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("for i=%d, j=%d \n", i, j)
		}
	}

	fmt.Println("*********")

}
