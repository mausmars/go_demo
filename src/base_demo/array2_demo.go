package main

import "fmt"

func main() {
	buf := []byte{1, 2, 3, 4, 5}

	buf2 := buf[4:5]
	fmt.Println(buf2)

	for i, b := range buf {
		buf[i] = b + 1
	}
	for _, b := range buf {
		fmt.Print(b, ",")
	}
}
