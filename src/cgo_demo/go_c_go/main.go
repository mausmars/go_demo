package main

// go run .

/*
extern int helloFromC();
*/
import "C"

func main() {
	//call c function
	C.helloFromC()
}
