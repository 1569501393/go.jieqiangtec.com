package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("Hello world \n")

	arrayTest()
	sliceTest()

	// var arr = [5]int{0, 1, 2, 3, 4}
	// sum := sum(arr[:])
	sum := sum([]int{0, 1, 2, 3, 4})
	fmt.Printf("sum is %v \n", sum)
}

func arrayTest() {
	var arr1 [5]int
	for i := 0; i < 5; i++ {
		arr1[i] = i * 2
	}

	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Aray at index %d is %d \n", i, arr1[i])
	}

	a := [...]string{"a", "b", "c"}
	for i, s := range a {
		fmt.Println("Array item", i, "is", s)
	}

	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	for i := 0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %s \n", i, arrKeyValue[i])
	}
}

func sliceTest() {
	var arr1 [6]int
	var slice1 []int = arr1[0:1]

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}
	// arr1 is [0 1 2 3 4 5]
	fmt.Printf("arr1 is %v \n", arr1)

	// Slice at 0 is 2
	// Slice at 1 is 3
	// Slice at 2 is 4
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d \n", i, slice1[i])
	}
	// slice1 is [2 3 4]
	fmt.Printf("slice1 is %v \n", slice1)

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice
	// slice1 = slice1[0:4]
	// for i := 0; i < len(slice1); i++ {
	// 	fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	// }
	// fmt.Printf("The length of slice1 is %d\n", len(slice1))
	// fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	slice2 := []int{1, 2, 3, 4}
	for i, v := range slice2 {
		fmt.Printf("Slice2 at %d is: %d\n", i, v)
	}

	seasons := make([]string, 4)
	seasons[0] = "Spring"
	seasons[1] = "Summer"
	seasons[2] = "Autumn"
	seasons[3] = "Winter"

	for i, season := range seasons {
		fmt.Printf("seasons at %d is: %+v, %#v  \n", i, season, seasons)
	}

	slice2 = make([]int, 0, 10)
	for i := 0; i < cap(slice2); i++ {
		slice2 = slice2[0 : i+1]
		slice2[i] = i
		fmt.Printf("The length of slice2 is %d\n", len(slice2))
	}

	// print the slice:
	for i := 0; i < len(slice2); i++ {
		fmt.Printf("slice2 at %d is %d\n", i, slice2[i])
	}

	slForm := []int{1, 2, 3}
	slTo := make([]int, 10)
	n := copy(slTo, slForm)
	// sl_to is [1 2 3 0 0 0 0 0 0 0]
	fmt.Printf("sl_to is %v \n ", slTo)
	fmt.Printf("Copied %d elements \n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 9, 5, 6)
	fmt.Printf("sl3 is %v \n", sl3)

	s := "hello, 世界"
	for i, c := range s {
		fmt.Printf("%s, %d, %c \n", s, i, c)
	}

	fmt.Printf("s[2:3] is %v \n", s[0:1])
	fmt.Printf("%v,sort.IntsAreSorted(sl3) is %v \n", sl3, sort.IntsAreSorted(sl3))

}

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}
