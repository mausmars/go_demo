package main

//https://godoc.org/github.com/hashicorp/golang-lru

import (
	"fmt"
	lru "github.com/hashicorp/golang-lru"
)

func main() {
	cache, _ := lru.New(3)
	cache.Add("bill", 20)
	cache.Add("dable", 19)

	v, ok := cache.Get("bill")
	if ok {
		fmt.Printf("bill's age is %v\n", v)
	}

	cache.Add("cat", "18")
	fmt.Printf("cache length is %d\n", cache.Len())

	v, ok = cache.Get("dable")
	if !ok {
		fmt.Printf("dable was evicted out\n")
	}
}