package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

func main() {
	var buf bytes.Buffer
	buf.WriteString(strconv.FormatInt(int64(111), 10))
	buf.WriteString("_")
	buf.WriteString(strconv.Itoa(int(2222)))

	fmt.Println(buf.String())
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.LittleEndian.Uint64(buf))
}
