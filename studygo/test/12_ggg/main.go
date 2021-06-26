package main

import "fmt"

const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

func main() {
	fmt.Println("111")

	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i[0])

	test4()

	a := [5]int{1, 2, 3, 4, 5}
	t := a[3:4:4]
	fmt.Println(t[0])

	// cannot use s2 (variable of type []int) as int value in argument to appendcompiler
	/* s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2)
	fmt.Println(s1)
	*/
}

/* func funcMui(x,y int)(sum int,error){
    return x+y,nil
} */

func hello(num ...int) {
	num[0] = 18
}

func test4() {
	/* a := [2]int{5, 6}
	b := [3]int{5, 6} */
	// cannot compare a == b (mismatched types [2]int and [3]int)compiler
	/* if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	} */
}
