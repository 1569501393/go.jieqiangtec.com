package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	scoreTest := make(map[string]int)

	for i := 0; i < 3; i++ {
		scoreTest[string(i)] = i
	}

	fmt.Printf("scoreMap:%v\nscoreMap0:%v\n", scoreTest, scoreTest["0"])

	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100

	fmt.Printf("scoreMap:%v \n, type of scoreMap:%v \n", scoreMap, scoreMap)

	// 判断某个键是否存在
	i := scoreMap["张三2"]
	fmt.Printf("i:%v \n, type of i:%v \n", i, i)

	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

}
