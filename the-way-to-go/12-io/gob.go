package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type pStruct struct {
	X, Y, Z int
	Name    string
}
type qStruct struct {
	X, Y *int32
	Name string
}

func main() {
	fmt.Printf("Hello io \n")
	a := 111

	// a val= 111, T=int
	fmt.Printf("a val= %#v, T=%T \n", a, a)

	p := pStruct{1, 2, 3, "Jieqiang"}
	fmt.Printf("p val= %#v, T=%T \n", p, p)

	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)

	err := enc.Encode(pStruct{
		X:    3,
		Y:    4,
		Z:    6,
		Name: "pythagoras",
	})

	if err != nil {
		log.Fatal("encode error:", err)
	}
	var q qStruct
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	// "pythagoras":{3, 4}
	fmt.Printf("%q:{%d, %d} \n", q.Name, *q.X, *q.Y)

	getPStruct()
}

func getPStruct() {
	p := pStruct{}
	// p val= main.pStruct{X:0, Y:0, Z:0, Name:""}, T=main.pStruct
	fmt.Printf("p val= %#v, T=%T \n", p, p)
}
