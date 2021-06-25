package main

//https://www.cnblogs.com/bradmiller/archive/2010/11/25/1887945.html

import (
	"fmt"
	"math"
	"strconv"
)

const (
	Accuracy_Float  = 1e-8
	Accuracy_Double = 1e-11
)

func main() {
	var n1 float64
	var n2 float64
	var n0 float64
	n0 = 0.0000000000000001
	n1 = 10.222222222222229
	n2 = 10.222222222222225
	if n1 == n2 {
		fmt.Println("==")
	} else {
		fmt.Println("!=")
	}
	fmt.Println("=====================")
	fmt.Println(math.Dim(n1, n2))
	isEqual := math.Dim(n1, n2) < n0
	if isEqual {
		fmt.Println("==")
	} else {
		fmt.Println("!=")
	}
	// f：要转换的浮点数
	// fmt：格式标记（b、e、E、f、g、G）
	// prec：精度（数字部分的长度，不包括指数部分）
	// bitSize：指定浮点类型（32:float32、64:float64）
	string1 := strconv.FormatFloat(n1, 'E', -1, 64)
	fmt.Println(string1)
	string2 := strconv.FormatFloat(n2, 'E', -1, 64)
	fmt.Println(string2)

	fmt.Println(math.Pow(2, 23))

	d := 96e+15
	fmt.Println(d)
	f64_a := math.Pow(2, 52)
	fmt.Println(f64_a)
	f64_b := f64_a + d
	fmt.Println(f64_b)
	fmt.Println(f64_a - f64_b)
}
