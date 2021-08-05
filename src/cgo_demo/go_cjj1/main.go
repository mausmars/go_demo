package main

import (
	"fmt"
	"go_demo/src/cgo_demo/go_cjj1/point"
)

func main() {
	fmt.Println("start")
	p := point.NewPoint(0.0, 0.0)
	q := point.NewPoint(3.0, 4.0)
	dist := point.Distance(p, q)
	fmt.Printf("Wrong distance %f \r\n", dist)
	p.Delete()
	q.Delete()
}
