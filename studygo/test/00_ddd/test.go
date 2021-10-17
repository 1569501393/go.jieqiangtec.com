package main

import (
	"fmt"
	"reflect"
)

var any interface{}

// init顺序
func init() {
	fmt.Printf("init111=== \n")
}

func init() {
	fmt.Printf("init222=== \n")
}

func init() {
	fmt.Printf("init333=== \n")
}

type People struct {
	// name string
	// age  int
}

func main() {
	name := "Golang"
	fmt.Printf("hello, %s \n", name)
	fmt.Printf("hello v, %v \n", name)
	fmt.Printf("hello +v, %+v \n", name)
	fmt.Printf("hello #v, %#v \n", name)

	// name's type=, string
	fmt.Printf("name's type=, %s \n", reflect.TypeOf(name).Kind())

	updateAny()

	// any切片
	any = []int{1, 2, 3, 4}
	// any= [%!s(int=1) %!s(int=2) %!s(int=3) %!s(int=4)]
	fmt.Printf("any= %s \n", any)
	// any= [1 2 3 4]
	fmt.Printf("any= %v \n", any)
	// any= []int{1, 2, 3, 4}
	fmt.Printf("any= %#v \n", any)

	// any 数字
	any = 222
	// any= %!s(int=222)
	fmt.Printf("any= %s \n", any)
	// any= 222
	fmt.Printf("any= %v \n", any)
	// any= 222
	fmt.Printf("any= %#v \n", any)

	// 数组
	// 一维数组
	var arr [5]int
	arr = [5]int{1, 2, 3, 4, 5}
	// arr= [5]int{1, 2, 3, 4, 5}
	fmt.Printf("arr= %#v \n", arr)

	// 二维数组
	var arr2 [2][3]int
	arr2 = [2][3]int{{1, 1, 1}, {2, 2, 2}}
	// arr2= [2][3]int{[3]int{1, 1, 1}, [3]int{2, 2, 2}}
	fmt.Printf("arr2= %#v \n", arr2)
	// len(arr2) = 2
	fmt.Printf("len(arr2)= %#v \n", len(arr2))
	// cap(arr2)= 2
	fmt.Printf("cap(arr2)= %#v \n", cap(arr2))
	arr2[0][0] = 100
	// arr2= [2][3]int{[3]int{100, 1, 1}, [3]int{2, 2, 2}}
	fmt.Printf("arr2= %#v \n", arr2)

	n := 100
	m := getPtr(&n)
	// ptr value ====
	// 	m0=(*int)(0xc0000aa0e8),
	// 	m1=(*int)(0xc0000aa0f0),
	// 	m2=(*int)(0xc0000aa0f8),
	// 	m3=(*int)(0xc0000aa100),
	// 	m4=(*int)(0xc0000aa108),
	// 	m5=(*int)(0xc0000aa110),
	// 	m= (*int)(0xc0000aa100) *m=0
	fmt.Printf("m= %#v *m=%#v\n", m, *m)

	// map循环每次不一致
	// var nameMap = make(map[string]string)
	// nameMap["zhang"] = "zhangSan"
	// nameMap["li"] = "liSi"
	// nameMap["wang"] = "wangWu"
	nameMap := map[string]string{
		"zhang": "zhangSan",
		"li":    "liSi",
		"wang":  "wangWu",
	}
	// nameMap= map[string]string{"li":"liSi", "wang":"wangWu", "zhang":"zhangSan"},
	// 	&nameMap= &map[string]string{"li":"liSi", "wang":"wangWu", "zhang":"zhangSan"}
	fmt.Printf("nameMap= %#v,\n&nameMap= %#v\n", nameMap, &nameMap)

	// nameMap= 0xc0000c2450,
	// 	&nameMap= 0xc0000ce020
	fmt.Printf("nameMap= %p,\n&nameMap= %p\n", nameMap, &nameMap)

	// for i= 0, namekey="li", nameValue="liSi"
	// for i= 0, namekey="wang", nameValue="wangWu"
	// for i= 0, namekey="zhang", nameValue="zhangSan"
	// for i= 1, namekey="li", nameValue="liSi"
	// for i= 1, namekey="wang", nameValue="wangWu"
	// for i= 1, namekey="zhang", nameValue="zhangSan"
	// for i= 2, namekey="zhang", nameValue="zhangSan"
	// for i= 2, namekey="li", nameValue="liSi"
	// for i= 2, namekey="wang", nameValue="wangWu"
	for i := 0; i < 3; i++ {
		// namekey="wang", nameValue="wangWu"
		// namekey="zhang", nameValue="zhangSan"
		// namekey="li", nameValue="liSi"
		for namekey, nameValue := range nameMap {
			fmt.Printf("for i= %d, namekey=%#v, nameValue=%#v\n", i, namekey, nameValue)
		}
	}

	nameSlice := []string{"zhangSan", "liSi", "wangWu"}
	// for i= 0, index=0, name="zhangSan"
	// for i= 0, index=1, name="liSi"
	// for i= 0, index=2, name="wangWu"
	// for i= 1, index=0, name="zhangSan"
	// for i= 1, index=1, name="liSi"
	// for i= 1, index=2, name="wangWu"
	// for i= 2, index=0, name="zhangSan"
	// for i= 2, index=1, name="liSi"
	// for i= 2, index=2, name="wangWu"
	for i := 0; i < 3; i++ {
		for index, name := range nameSlice {
			fmt.Printf("for i= %d, index=%#v, name=%#v\n", i, index, name)
		}
	}

	myName := "jieqiang"
	// myName= "jieqiang",
	// 	&myName= (*string)(0xc000088330)
	fmt.Printf("myName= %#v,\n&myName= %#v \n", myName, &myName)

	// 空结构体比较
	// struct{}{}= struct {}{},
	// 	struct{}{}= struct {}{},
	// 	struct{}{} == struct{}{}= true
	fmt.Printf("struct{}{}= %#v,\n struct{}{}= %#v,\n struct{}{} == struct{}{}= %#v \n", struct{}{}, struct{}{}, struct{}{} == struct{}{})

	p1 := &People{}
	p2 := &People{}

	// fmt.Printf("%p\n", p1)
	// fmt.Printf("%p\n", p2)

	// true
	fmt.Println(p1 == p2)

	// p1= 0xe9a338,
	// 	p2= 0xe9a338,
	// 	p1 == p2= true
	fmt.Printf("p1= %p,\n p2= %p,\n p1 == p2= %#v \n", p1, p2, p1 == p2)

	// p1= &main.People{name:"", age:0},
	// 	p2= &main.People{name:"", age:0},
	// 	p1 == p2= false

	// p1= &main.People{},
	// 	p2= &main.People{},
	// 	p1 == p2= true
	fmt.Printf("p1= %#v,\n p2= %#v,\n p1 == p2= %#v \n", p1, p2, p1 == p2)

}

// 更改any变量值
func updateAny() {
	any = "updateAny"
	// any in updateAny= "updateAny"
	fmt.Printf("any in updateAny= %#v  \n", any)
}

// go语言局部变量分配在栈还是堆
func getPtr(m0 *int) *int {
	m1 := new(int)
	m2 := new(int)
	m3 := new(int)
	m4 := new(int)
	// m5 := new(int)
	var m5 = new(int)
	fmt.Printf("ptr value ====\n m0=%#v,\n m1=%#v,\n m2=%#v,\n m3=%#v,\n m4=%#v,\n m5=%#v,\n", m0, m1, m2, m3, m4, m5)

	return m3
}
