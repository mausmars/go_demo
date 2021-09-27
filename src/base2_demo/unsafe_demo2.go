package main

import (
	"fmt"
	"unsafe"
)

type Test1 struct {
	a uint8
	b uint8
	c uint8
	d uint32
	e uint64
}

type Test2 struct {
	a uint8
	b uint8
	c uint8
	d uint32
	e uint64
}

func main() {
	data := make([]byte, 32)
	for i := 0; i < len(data); i++ {
		data[i] = 0
	}
	address := (**Test1)(unsafe.Pointer(&data))
	(*address).a = uint8(1)
	(*address).b = uint8(2)
	(*address).c = uint8(3)
	(*address).d = uint32(4)
	(*address).e = uint64(5)
	for i := 0; i < len(data); i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println("--------------------------")
}
