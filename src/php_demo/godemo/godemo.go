package main

import "fmt"

/*
typedef struct  {
    int A;
    int B;
}TestBean;
*/
import "C"

//export godemoHello
func godemoHello() {
	fmt.Println("hello world")
}

//export godemoAdd
func godemoAdd(a, b int) int {
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
