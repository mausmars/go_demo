package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"hash/crc32"
)

func main() {
	m := make(map[uint32]string)
	count:=0

	for i:=0;i<10000000;i++{
		u1 := uuid.NewV4()
		num := crc32.ChecksumIEEE(u1.Bytes())
		if _, ok := m[num]; !ok {
			m[num] = ""
		}else{
			count++
		}
	}
	fmt.Println(count)
}