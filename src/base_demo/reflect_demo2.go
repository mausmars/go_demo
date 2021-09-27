package main

import (
	"fmt"
	"reflect"
)

func main() {
	//接口到反射对象
	author := "draven"
	fmt.Println("TypeOf author:", reflect.TypeOf(author))
	fmt.Println("ValueOf author:", reflect.ValueOf(author))

	//反射对象到接口
	v := reflect.ValueOf(1)
	a := v.Interface().(int)
	fmt.Println("a:", a)

	i := 1
	b := reflect.ValueOf(&i)
	b.Elem().SetInt(10)
	fmt.Println(i)
	fmt.Println("b:", b)
}
