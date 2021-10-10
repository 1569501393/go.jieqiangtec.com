package main

import (
	"bytes"
	"fmt"
)

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Printf("Value of asset is %f, %#v \n", asset.getValue(), asset)
}

type ReadWrite interface {
	Read(b bytes.Buffer) bool
	Write(b bytes.Buffer) bool
}

type Lock interface {
	Lock()
	Unlock()
}

type Stringer interface {
	String() string
}

type File interface {
	ReadWrite
	Lock
	Close()
}

func main() {
	fmt.Printf("Hello world \n")

	sq1 := new(Square)
	sq1.side = 5

	var areaIntf Shaper
	areaIntf = sq1

	fmt.Printf("The square has area %f, %T \n", areaIntf.Area(), areaIntf.Area())
	fmt.Printf("The square has area %#v, %T \n", areaIntf.Area(), areaIntf.Area())
	fmt.Printf("The square has area %+v, %T \n", areaIntf.Area(), areaIntf.Area())

	r := Rectangle{
		length: 5,
		width:  3,
	}

	q := &Square{side: 5}
	// q := Square{side: 5}

	shapes := []Shaper{r, q}
	fmt.Printf("Looping through shapes for area... \n")

	for n, shape := range shapes {
		fmt.Printf("Shape details:%#v, %#v, %#v \n", shapes[n], shape, n)
		fmt.Printf("Area of this shape is:%#v \n", shapes[n].Area())
	}

	/*o := stockPosition{
		ticker:     "GOOG",
		sharePrice: 577.2,
		count:      4,
	}*/

	/*var o valuable = stockPosition{
		ticker:     "GOOG",
		sharePrice: 577.2,
		count:      4,
	}*/

	/*var o = stockPosition{
		ticker:     "GOOG",
		sharePrice: 577.2,
		count:      4,
	}*/

	var o valuable
	o = stockPosition{
		ticker:     "GOOG",
		sharePrice: 577.2,
		count:      4,
	}

	showValue(o)

	/*o2 := car{
		make:  "BMW",
		model: "M3",
		price: 66500,
	}*/
	o = car{
		make:  "BMW",
		model: "M3",
		price: 66500,
	}
	showValue(o)

	// 测试一个值是否实现了某个接口
	var v interface{}
	if sv, ok := v.(Stringer); ok {
		fmt.Printf("v implements String(): %#v \n", sv.String())
	} else {
		fmt.Printf("v not implements String(): %#v \n", sv)
	}
}
