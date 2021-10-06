package main

import (
	"fmt"
	pack1 "gin111/the-way-to-go/09-package/09-02-package"
	"math"
	"math/big"
	"regexp"
	"strconv"
)

func main() {
	fmt.Printf("Hello world \n")
	regexpTest()
	mathTest()
	str := pack1.ReturnStr()
	fmt.Printf("str in package1 is `%v`  \n", str)
	fmt.Printf("Integer in package1 is `%v`  \n", pack1.PackInt)
}

func regexpTest() {
	// searchIn := "hello,world 666"
	//
	// pat := "[0-9]+.[0-9]+"
	//
	// f := func(s string) string {
	// 	v, _ := strconv.ParseFloat(s, 32)
	// 	return strconv.FormatFloat(v*2, 'f', 2, 32)
	// }
	//
	// if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
	// 	fmt.Println("match found!")
	// }
	//
	// re, _ := regexp.Compile(pat)
	//
	// str := re.ReplaceAllString(searchIn, "##.#")
	// fmt.Println(str)
	//
	// str2 := re.ReplaceAllFunc(searchIn, f)
	// fmt.Println(str2)

	// 目标字符串
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+" // 正则

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match Found!")
	}

	re, _ := regexp.Compile(pat)
	// 将匹配到的部分替换为"##.#"
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)
	// 参数为函数时
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
}

func mathTest() {
	// Here are some calculations with bigInts:
	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(1)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	// Big int:43492122561469640008497075573153004
	fmt.Printf("Big int:%v \n", ip)

	// Here are some calculations with bigInts:
	rm := big.NewRat(math.MaxInt64, 1956)
	rn := big.NewRat(-1956, math.MaxInt64)
	ro := big.NewRat(19, 56)
	rp := big.NewRat(1111, 2222)
	rq := big.NewRat(1, 1)
	rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
	// Big Rat:-37/112
	fmt.Printf("Big Rat:%v \n", rq)

}
