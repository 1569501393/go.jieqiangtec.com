package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello world \n")
	mapTest()
}

func mapTest() {
	var user map[string]int
	user = map[string]int{"zhangSan": 18, "liSi": 19, "wangWu": 20}
	fmt.Printf("user = %v \n", user)

	var user2 = make(map[string]int)
	user2["zhangSan"] = 18
	user2["liSi"] = 19
	user2["wangWu"] = 20
	fmt.Printf("user2 = %v \n", user2)
	fmt.Printf("user2[\"zhangSan\"] = %v \n", user2["zhangSan"])
	aaa := user2["aaa"]

	if _, ok := user2["aaa"]; ok {
		fmt.Printf("user2[\"aaa\"] = %v \n", user2["aaa"])
	}

	fmt.Printf("user2[\"aaa\"] = %v \n", aaa)

}
