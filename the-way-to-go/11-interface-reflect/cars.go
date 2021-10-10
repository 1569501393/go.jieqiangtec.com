package main

import (
	"fmt"
	"reflect"
)

type Any interface {
}

type Car struct {
	Model        string
	Manufacturer string
	BuildYear    int
}

type Cars []*Car

func main() {
	// fmt.Printf("type:= reflect.TypeOf(x): %v, val= %#v, T=%T \n", reflect.TypeOf(x), x, x)
	fmt.Printf("hello, reflect \n")
	// fmt.Printf("val= %#v, T=%T \n", val, val)

	var a Any = 1
	fmt.Printf("a: %#v, T:%T  \n", a, a)

	// make some cars
	// 完整定义并赋值1-指针传递
	ford := &Car{
		Model:        "Fiesta",
		Manufacturer: "Ford",
		BuildYear:    2008,
	}
	// ford type:= reflect.TypeOf(x): *main.Car, val= &main.Car{Model:"Fiesta", Manufacturer:"Ford", BuildYear:2008}, T=*main.Car
	fmt.Printf("ford type:= reflect.TypeOf(x): %v, val= %#v, T=%T \n", reflect.TypeOf(ford), ford, ford)

	// 完整定义并赋值1-值传递
	ford2 := Car{
		Model:        "Fiesta2",
		Manufacturer: "Ford2",
		BuildYear:    20082,
	}
	// ford2 type:= reflect.TypeOf(x): main.Car, val= main.Car{Model:"Fiesta2", Manufacturer:"Ford2", BuildYear:20082}, T=main.Car
	fmt.Printf("ford2 type:= reflect.TypeOf(x): %v, val= %#v, T=%T \n", reflect.TypeOf(ford2), ford2, ford2)

	ford3 := ford2
	ford3.Model = "Fiesta3"
	// ford3 type:= reflect.TypeOf(x): main.Car, val= main.Car{Model:"Fiesta3", Manufacturer:"Ford2", BuildYear:20082}, T=main.Car
	fmt.Printf("ford3 change the model type:= reflect.TypeOf(x): %v, val= %#v, T=%T \n", reflect.TypeOf(ford3), ford3, ford3)
	// ford2 type:= reflect.TypeOf(x): main.Car, val= main.Car{Model:"Fiesta2", Manufacturer:"Ford2", BuildYear:20082}, T=main.Car
	fmt.Printf("ford2 no change the model type:= reflect.TypeOf(x): %v, val= %#v, T=%T \n", reflect.TypeOf(ford2), ford2, ford2)

	fordPtr := ford
	fordPtr.Model = "FiestaPtr"
	// fordPtr type:= reflect.TypeOf(x): *main.Car, val= &main.Car{Model:"FiestaPtr", Manufacturer:"Ford", BuildYear:2008}, T=*main.Car
	fmt.Printf("fordPtr change the model type:= reflect.TypeOf(x): %v, val= %#v, T=%T \n", reflect.TypeOf(fordPtr), fordPtr, fordPtr)
	// ford type:= reflect.TypeOf(x): *main.Car, val= &main.Car{Model:"FiestaPtr", Manufacturer:"Ford", BuildYear:2008}, T=*main.Car
	fmt.Printf("ford change the model type:= reflect.TypeOf(x): %v, val= %#v, T=%T \n", reflect.TypeOf(ford), ford, ford)

	// 简化定义并赋值
	bmw := &Car{"XL 450", "BMW", 2011}
	// bmw type:= reflect.TypeOf(x): *main.Car, val= &main.Car{Model:"XL 450", Manufacturer:"BMW", BuildYear:2011}, T=*main.Car
	fmt.Printf("bmw type:= reflect.TypeOf(x): %v, val= %#v, T=%T \n", reflect.TypeOf(bmw), bmw, bmw)

	bmw2 := &Car{"x 800", "BMW", 2008}
	// bmw2 type:= reflect.TypeOf(x): *main.Car, val= &main.Car{Model:"x 800", Manufacturer:"BMW", BuildYear:2008}, T=*main.Car
	fmt.Printf("bmw2 type:= reflect.TypeOf(x): %v, val= %#v, T=%T \n", reflect.TypeOf(bmw2), bmw2, bmw2)

	// 完整定义并赋值2
	merc := new(Car)
	merc.Model = "D600"
	merc.Manufacturer = "Mercedes"
	merc.BuildYear = 2009
	// ford type:= reflect.TypeOf(x): *main.Car, val= &main.Car{Model:"D600", Manufacturer:"Mercedes", BuildYear:2009}, T=*main.Car
	fmt.Printf("ford type:= reflect.TypeOf(x): %v, val= %#v, T=%T \n", reflect.TypeOf(merc), merc, merc)

	// query
	allCars := Cars{ford, bmw, merc, bmw2}
	// allCars type:= reflect.TypeOf(x): main.Cars, val= main.Cars{(*main.Car)(0xc0000c4390), (*main.Car)(0xc0000c4600), (*main.Car)(0xc0000c46c0), (*main.Car)(0xc0000c4660)}, T=main.Cars
	fmt.Printf("allCars type:= reflect.TypeOf(x): %v, val= %#v, T=%T \n", reflect.TypeOf(allCars), allCars, allCars)

	allNewBMWs := allCars.FindAll(func(car *Car) bool {
		return (car.Manufacturer == "BMW") && (car.BuildYear > 2010)
	})
	// allNewBMWs type:= reflect.TypeOf(x): main.Cars, val= main.Cars{(*main.Car)(0xc0000c4600)}, T=main.Cars
	fmt.Printf("allNewBMWs type:= reflect.TypeOf(x): %v, val= %#v, T=%T \n", reflect.TypeOf(allNewBMWs), allNewBMWs, allNewBMWs)

}
func (cs Cars) Process(f func(car *Car)) {
	for _, c := range cs {
		f(c)
	}
}
func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car, 0)
	cs.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})

	return cars
}
