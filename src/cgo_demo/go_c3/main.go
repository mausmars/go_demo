package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")

	p := NewPoint(0.0, 0.0)
	q := NewPoint(3.0, 4.0)
	dist := Distance(p, q)

	fmt.Printf("Wrong distance %f \r\n", dist)

	p.Delete()
	q.Delete()
}
