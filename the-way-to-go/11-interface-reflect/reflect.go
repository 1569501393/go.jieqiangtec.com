package main

import (
	"fmt"
	"reflect"
)

type Any interface {
}

func main() {
	fmt.Printf("hello, reflect \n")
	// fmt.Printf("val= %#v, T=%T \n", val, val)

	var a Any = 1
	fmt.Printf("a: %#v, T:%T  \n", a, a)

	var x float64 = 3.4
	// type:= reflect.TypeOf(x): float64, val= 3.4, T=float64
	fmt.Printf("type:= reflect.TypeOf(x): %v, val= %#v, T=%T \n", reflect.TypeOf(x), x, x)

	v := reflect.ValueOf(x)
	// v := reflect.ValueOf(x): 3.4, val= 3.4, T=reflect.Value
	fmt.Printf("v := reflect.ValueOf(x): %v, val= %#v, T=%T \n", v, v, v)

	// type: float64, val= &reflect.rtype{size:0x8, ptrdata:0x0, hash:0x2ea27ffb, tflag:0x7, align:0x8, fieldAlign:0x8, kind:0xe, equal:(func(unsafe.Pointer, unsafe.Pointer) bool)(0x2c3400), gcdata:(*uint8)(0x3b1177), str:5921, ptrToThis:27040}, T=*reflect.rtype
	fmt.Printf("type: %v, val= %#v, T=%T \n", v.Type(), v.Type(), v.Type())

	// kind: float64, val= 0xe, T=reflect.Kind
	fmt.Printf("kind: %v, val= %#v, T=%T \n", v.Kind(), v.Kind(), v.Kind())

	// value: 3.4, val= 3.4, T=float64
	fmt.Printf("value: %v, val= %#v, T=%T \n", v.Float(), v.Float(), v.Float())

	// v.interface: 3.4, val= 3.4, T=float64
	fmt.Printf("v.interface: %v, val= %#v, T=%T \n", v.Interface(), v.Interface(), v.Interface())

	y := v.Interface().(float64)
	// v.interface: 3.4, val= 3.4, T=float64
	fmt.Printf("v.interface: %v, val= %#v, T=%T \n", y, y, y)

	// 通过反射修改 (设置) 值
	// settability of v: false
	fmt.Println("settability of v:", v.CanSet())
	v = reflect.ValueOf(&x) // Note: take the address of x.
	// settability of v: false
	fmt.Println("settability of v:", v.CanSet())

	v = v.Elem()
	// v.Elem(): 3.4, val= 3.4, T=reflect.Value
	fmt.Printf("v.Elem(): %v, val= %#v, T=%T \n", v, v, v)

	// settability of v: true
	fmt.Println("settability of v:", v.CanSet())

	v.SetFloat(666.555)
	// v.SetFloat: 666.555, val= 666.555, T=reflect.Value
	fmt.Printf("v.SetFloat: %v, val= %#v, T=%T \n", v, v, v)

}
