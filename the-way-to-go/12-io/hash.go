package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

func main() {
	fmt.Printf("Hello io \n")
	a := 111

	// a val= 111, T=int
	fmt.Printf("a val= %#v, T=%T \n", a, a)

	hasher := sha1.New()
	io.WriteString(hasher, "test")
	// b := []byte{}
	var b []byte

	// b val= []byte{}, T=[]uint8
	fmt.Printf("b val= %#v, T=%T \n", b, b)
	// hasher.Sum(b) val= []byte{0xa9, 0x4a, 0x8f, 0xe5, 0xcc, 0xb1, 0x9b, 0xa6, 0x1c, 0x4c, 0x8, 0x73, 0xd3, 0x91, 0xe9, 0x87, 0x98, 0x2f, 0xbb, 0xd3}, T=[]uint8
	fmt.Printf("hasher.Sum(b) val= %#v, T=%T \n", hasher.Sum(b), hasher.Sum(b))

	hasher.Reset()
	data := []byte("We shall overcome!")
	// data := []byte("")
	n, err := hasher.Write(data)
	if n != len(data) || err != nil {
		log.Printf("Hash write error: %v /%v", n, err)
	}

	checkSum := hasher.Sum(b)
	// Result: e2222bfc59850bbb00a722e764a555603bb59b2a
	fmt.Printf("Result: %x \n", checkSum)
}
