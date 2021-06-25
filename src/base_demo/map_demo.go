package main

import "fmt"

func main() {
	m := make(map[int32]int32)

	m[1] += 10
	fmt.Println(m[1])
}
