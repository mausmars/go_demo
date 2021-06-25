package main

//https://www.coder.work/article/199596

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"reflect"
)

type Enum int

const (
	Zero Enum = 0
)

type IPerson interface {
	name() string
}

type Person struct {
	a string
}

func (p *Person) name() string {
	return "abc"
}

func showType(obj interface{}) {
	oType := reflect.TypeOf(obj)
	fmt.Println("type name= ", oType.String())
}

func main() {
	p := &Person{a: "a"}
	pType := reflect.TypeOf(p)

	m := map[reflect.Type]string{}
	m[reflect.TypeOf(p)] = pType.String()

	fmt.Println("type name= ", pType.String())
	fmt.Println("type name= ", pType.Align())
	fmt.Println("type name= ", pType.PkgPath())
	fmt.Println("type name= ", pType.Kind())

	pType3:= reflect.TypeOf(&Person{})
	newP := reflect.New(pType3).Elem().Interface()
	np:=newP.(*Person)
	fmt.Println(np)
	fmt.Println("---------------------------------")
	real := new(Person)
	reflected := reflect.New(reflect.TypeOf(real).Elem()).Interface()
	fmt.Println(real)
	fmt.Println(reflected)
	fmt.Println("---------------------------------")
	fmt.Println("type name= ", pType3.Elem().Kind())
	fmt.Println("type name= ", pType3.Elem().Name())
	fmt.Println("---------------------------------")
	p2 := &Person{}
	fmt.Println(m[reflect.TypeOf(p2)])

	fmt.Println("---------------------------------")
	showType(p)

	fmt.Println("---------------------------------")
	pType2 := reflect2.TypeOf(p)
	fmt.Println("type name= ", pType2.String())
	fmt.Println("---------------------------------")
	// 获取Zero常量的反射类型对象
	typeOfA := reflect.TypeOf(Zero)
	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
	fmt.Println("---------------------------------")




}
