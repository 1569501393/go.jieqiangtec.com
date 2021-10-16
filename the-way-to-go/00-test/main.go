package main

import (
	"fmt"
	"time"
)

const LIM = 41

var fibs [LIM]uint64

func main() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)

	/*// 泛型 go run -gcflags=-G=3
	strs := []string{"a", "b", "c"}

	// /d/www/qldev/go/go.jieqiangtec.com
	// $ go run -gcflags=-G=3 ./the-way-to-go/00-test/main.go
	print(strs)*/
}

/*func print[T any](arr []T)  {
	for i, v := range arr {
		fmt.Printf("i=%#v, v=%#v \n", i,v)
	}

}*/

func fibonacci(n int) (res uint64) {
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
