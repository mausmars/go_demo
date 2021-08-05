package main

import "C"
import "fmt"

//export HelloFromGo
func HelloFromGo() {
	fmt.Printf("Hello from Go!\n")
}
