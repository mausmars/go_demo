package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

func main(){
	projectId:="100027"
	timestamp:=time.Now().Second()
	key:="8a2037cc-1195-11ea-8848-060d5b3e7372"

	str:=projectId+":"+strconv.Itoa(timestamp)+":"+key

	c := md5.New()
	c.Write([]byte(str))
	v:=hex.EncodeToString(c.Sum(nil))
	fmt.Println(v)
}
