package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type TestBean1 struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Data1 []int   `json:"data1"`
	Data2 [][]int `json:"data2"`
}

type TestBean2 struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Data1 []int   `json:"data1"`
	Data2 [][]int `json:"data2"`
}

func obj2jsonTestBean1() []byte {
	obj1 := TestBean1{}
	obj1.Id = 1
	obj1.Name = "test1"
	obj1.Data1 = []int{1, 2, 3}
	obj1.Data2 = [][]int{{1, 2}, {3, 4}}

	data, err := json.Marshal(obj1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \r\n", data)

	return data
}

func reflectTest1() {
	fmt.Println("--- reflectTest1 ---")
	data := obj2jsonTestBean1()

	clsType := reflect.TypeOf(&TestBean1{})
	iobj := reflect.New(clsType.Elem()).Interface()

	typ := reflect.TypeOf(iobj)
	fmt.Println("typ= ", typ)

	err := json.Unmarshal(data, iobj)
	if err != nil {
		panic(err)
	}
	fmt.Println(iobj)

	obj3 := *iobj.(*TestBean1)
	v := reflect.ValueOf(obj3)
	k := v.FieldByName("Id")
	fmt.Println(k)
}

func reflectTest2() {
	fmt.Println("--- reflectTest2 ---")
	data := obj2jsonTestBean1()

	obj2 := &TestBean1{}
	var iobj interface{} = obj2

	typ := reflect.TypeOf(iobj)
	fmt.Println("typ= ", typ)

	err := json.Unmarshal(data, iobj)
	if err != nil {
		panic(err)
	}
	fmt.Println(iobj)

	obj3 := *iobj.(*TestBean1)
	v := reflect.ValueOf(obj3)
	k := v.FieldByName("Id")
	fmt.Println(k)
}

func reflectTest3() {
	fmt.Println("--- reflectTest3 ---")
	data := obj2jsonTestBean1()

	var obj2 interface{} = TestBean1{}
	var iobj = &obj2

	//先把结构体 转interface 再取地址就有问题
	//需要先转指针 在转interface
	//Unmarshal 方法 应该是无法确定类型 就转成map了

	//** 总结：尽量使用指针，可以保留类型，结构体转interface{} 类型丢失了。。。
	//感觉go 语言应该禁用interface{}类型转指针的操作！不符合c语言的规则。。。
	
	typ := reflect.TypeOf(iobj)
	fmt.Println("typ= ", typ)
	fmt.Println("Elem typ= ", typ.Elem())

	err := json.Unmarshal(data, iobj)
	if err != nil {
		panic(err)
	}
	fmt.Println(*iobj)

	//obj3 := iobj.(TestBean1)
	//v := reflect.ValueOf(obj3)
	//k := v.FieldByName("Id")
	//fmt.Println(k)
}

func main() {
	reflectTest1()
	reflectTest2()
	reflectTest3()
}
