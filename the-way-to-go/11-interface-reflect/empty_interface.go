package main

import "fmt"

var i = 5
var str = "ABC"

type Person struct {
	name string
	age  int
}

type Any interface {
}

func main() {
	var val Any = 5
	var val2 = 5
	// val= 5, T=int
	fmt.Printf("val= %#v, T=%T \n", val, val)

	// val2= 5, T=int
	fmt.Printf("val2= %#v, T=%T \n", val2, val2)

	val = str
	fmt.Printf("val= %#v, T=%T \n", val, val)

	// 	.\empty_interface.go:28:7: cannot use str (type string) as type int in assignment
	// val2 = str

	pers1 := Person{
		name: "Rob Pike",
		age:  55,
	}

	val = pers1
	// val= main.Person{name:"Rob Pike", age:55}, T=main.Person
	fmt.Printf("val= %#v, T=%T \n", val, val)

	// pers2 := new(Person)
	pers2 := pers1
	pers2.name = "jieqiang"
	// val= &main.Person{name:"jieqiang", age:0}, T=*main.Person
	fmt.Printf("Person重置前 pers2= %#v, T=%T \n", pers2, pers2)

	pers3 := new(Person)
	pers3.age = 18
	pers3.name = "kk"
	// val= &main.Person{name:"jieqiang", age:0}, T=*main.Person
	fmt.Printf("Person重置后 pers2= %#v, T=%T \n", pers2, pers2)
	fmt.Printf("Person重置后 pers3= %#v, T=%T \n", pers3, pers3)
	fmt.Printf("Person重置后 val= %#v, T=%T \n", val, val)

	per4 := pers3
	per4.age = 19
	fmt.Printf("per4 := pers3 per4= %#v, T=%T \n", per4, per4)
	fmt.Printf("per4 := pers3 per3= %#v, T=%T \n", pers3, pers3)

	switch t := val.(type) {
	case int:
		fmt.Printf("Type int  %#v, T=%T \n", t, t)
	case string:
		fmt.Printf("Type string  %#v, T=%T \n", t, t)
	case bool:
		fmt.Printf("Type boolean  %#v, T=%T \n", t, t)
	case *Person:
		fmt.Printf("Type pointer to Person  %#v, T=%T \n", t, t)
	case Person:
		fmt.Printf("Type Person  %#v, T=%T \n", t, t)
	default:
		fmt.Printf("Type Unexpected   %#v, T=%T \n", t, t)
	}

}
