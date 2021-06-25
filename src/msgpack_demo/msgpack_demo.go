package main

import (
	"fmt"
	"github.com/shamaton/msgpack"
)

func main() {
	type TestBean struct {
		S string
		A int32
	}

	v := &TestBean{
		S: "msgpack",
		A:1,
	}

	d, err := msgpack.Encode(v)
	if err != nil {
		panic(err)
	}
	fmt.Println(d)

	r := &TestBean{}
	err = msgpack.Decode(d, r)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}