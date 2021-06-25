package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"hash/crc64"
)

func main() {
	m := make(map[uint64]string)
	count:=0

	for i:=0;i<100000000;i++{
		u1 := uuid.NewV4()

		table := crc64.MakeTable(crc64.ECMA)
		num := crc64.Checksum(u1.Bytes(), table)
		if _, ok := m[num]; !ok {
			m[num] = ""
		}else{
			count++
		}
	}
	fmt.Println(count)
}