package main

import (
	"fmt"
	"reflect"
	"strings"
)

type struct1 struct {
	i1  int
	f1  float32
	str string
}

type Person struct {
	firstName string
	lastName  string
}

type TagType struct {
	F1 bool   `json:"An important answer"`
	F2 string `json:"The name of the ting"`
	F3 int    `json:"How much there are"`
}

type TwoInts struct {
	a int
	b int
}

func main() {
	fmt.Printf("Hello world \n")
	structTest()

	var p1 Person
	p1.firstName = "Chris"
	p1.lastName = "Woodward"
	upPerson(&p1)
	fmt.Printf("The name of the person is %s %s, %#v \n", p1.firstName, p1.lastName, p1)

	tt := TagType{
		F1: true,
		F2: "Barak Obama",
		F3: 1,
	}

	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}

	two1 := TwoInts{
		a: 12,
		b: 10,
	}

	fmt.Printf("The sum is: %d \n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(10))
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)

	fmt.Printf("%v \n", ixField.Tag)

}

func structTest() {
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.6
	ms.str = "Chris"

	fmt.Printf("The int is: %d, v:%v \n", ms.i1, ms)
	fmt.Printf("The float  is: %f, +v:%+v \n", ms.f1, ms)
	fmt.Printf("The string  is: %s, #v:%#v \n", ms.str, ms)
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToLower(p.lastName)
}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}
