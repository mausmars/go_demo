package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"hash/crc32"
)

// 生成md5
func MD5(str string) string {
	c := md5.New()
	c.Write([]byte(str))
	return hex.EncodeToString(c.Sum(nil))
}

//生成sha1
func SHA1(str string) string{
	c:=sha1.New()
	c.Write([]byte(str))
	return hex.EncodeToString(c.Sum(nil))
}

func CRC32(str string) uint32{
	return crc32.ChecksumIEEE([]byte(str))
}

func main() {
	fmt.Println(CRC32("123456"))
	fmt.Println(MD5("123456"))
	fmt.Println(SHA1("123456"))
}