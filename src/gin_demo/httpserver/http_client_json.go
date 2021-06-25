package main

import (
	"bytes"
	"crypto/tls"
	jsonmsg "go_demo/src/gin_demo/httpserver/json"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	url := "http://localhost:8008/json"

	msg := &jsonmsg.Command{
		CommandId: 1,
		Body:      "test",
	}
	data, _ := json.Marshal(msg)
	buf := bytes.NewBuffer(data)
	req, _ := http.NewRequest("POST", url, buf)

	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			// 不验证证书
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)

	reMsg := &jsonmsg.Command{}
	json.Unmarshal(body, reMsg)

	fmt.Println(resp.StatusCode)
	fmt.Println(reMsg)
	fmt.Println(resp.Header)
	fmt.Println(resp.Proto)
}
