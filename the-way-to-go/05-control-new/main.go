package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Printf("Abs(-5) is %v \n", Abs(-5))
	IfTest()
	SwitchTest()
	forTest()
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
func IfTest() {
	bool := true
	if bool {
		fmt.Printf("the value is true \n")
	} else {
		fmt.Printf("the value is false \n")
	}

	var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { //Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}

	fmt.Printf("prompt is %v \n", prompt)

	filePath := "./aaa2.txt"
	if _, err := os.Open(filePath); err != nil {
		fmt.Printf("os.Open(%s):%s \n", filePath, err)
	}
}

func SwitchTest() {
	age := 18
	switch age {
	case 18, 19, 20:
		fmt.Printf("age is %v, in 18, 19, 29 \n", age)
	case 28:
		fmt.Printf("age is %v,equal 28 \n", age)
	case 38:
		fmt.Printf("age is %v,equal 38 \n", age)
	default:
		fmt.Printf("age is %v,equal other \n", age)
	}

	switch {
	case age < 18:
		fmt.Printf("age is %v,age < 18 \n", age)
	case age >= 18:
		fmt.Printf("age is %v,age >= 18 \n", age)
	}

	k := 6

	switch k {
	case 4:
		fmt.Println(k, "was <= 4")
		fallthrough
	case 5:
		fmt.Println(k, "was <= 5")
		fallthrough
	case 6:
		fmt.Println(k, "was <= 6")
		fallthrough
	case 7:
		fmt.Println(k, "was <= 7")
		fallthrough
	case 8:
		fmt.Println(k, "was <= 8")
		fallthrough
	default:
		fmt.Println(k, "default case")
	}
}

func forTest() {
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration \n", i)
	}

	str := "Go is a beautiful language"
	fmt.Printf("'%v' :The length of str is: %v  \n", str, len(str))
	for i := 0; i < len(str); i++ {
		fmt.Printf("Character on position %d is: %c \n", i, str[i])
	}

	str2 := "Chinese: 日本语"
	fmt.Printf("The length of str2 is: %d \n", len(str2))
	for pos, char := range str2 {
		fmt.Printf("'%v' :Character on position %d is: %c \n", str2, pos, char)
	}

	for pos, char := range str {
		fmt.Printf("'%v' :Character on position %d is: %c \n", str, pos, char)
	}

	println("index \t int(rune) \t rune \t char \t bytes")
	for index, rune := range str2 {
		fmt.Printf("%-2d \t %d \t %U \t '%c' \t %X \n", index, rune, rune, rune, []byte(string(rune)))

	}

	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}

	// for i := 0; ; i++ {
	// 	fmt.Println("Value of i is now:", i)
	// }

	// for i := 0; i < 3; {
	// 	fmt.Println("Value of i:", i)
	// }

	var i int = 5
	for {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
		if i < 0 {
			// break
			continue
		}
		fmt.Printf("The variable i is now2: %d\n", i)
	}
}
