package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func main() {
	pid := 80000017
	secret_key := "c249ebf3-eee9-452f-b4d3-9f3e5dcb793c"
	salt := 1574233377811995614
	str := strconv.Itoa(pid) + ":" + secret_key + ":" + strconv.Itoa(salt)

	fmt.Println(strings.ToUpper(GetRandomString(64)))

	w := md5.New()
	io.WriteString(w, str)
	md5Str := fmt.Sprintf("%x", w.Sum(nil)) //w.Sum(nil)将w的hash转成[]byte格式
	fmt.Println(strings.ToUpper(md5Str))
}
