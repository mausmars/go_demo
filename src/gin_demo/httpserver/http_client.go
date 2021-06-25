package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	url := "http://localhost:8008/ping"

	req, _ := http.NewRequest("GET", url, nil)

	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			// 不验证证书
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(string(body))
	fmt.Println(resp.Header)
	fmt.Println(resp.Proto)
}
