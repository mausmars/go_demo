package main

import "fmt"

/*
typedef struct  {
    int A;
    int B;
}TestBean;
*/
import "C"

//export hello
func hello() {
	fmt.Println("hello world")
}

//export add
func add(a, b int) int {
	return a + b
}

//export createTestBean
func createTestBean(a, b int) C.TestBean {
	pojo := C.TestBean{}
	pojo.A = C.int(a)
	pojo.B = C.int(b)
	return pojo
}

func main() {
}
