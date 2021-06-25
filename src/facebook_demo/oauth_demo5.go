package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	token:="EAAKfdeoZBLyYBAKa0ZBho1Ug2GKhcwyTw2EC2cOrYmXZBzZCaJ8ZAfO79sUBMEMPvqOZCGS0njX8umvQqqC10ZC48mJk0c56yi3m534oZAoTTYfSE2xQqBmaWtZAEQ4ZCBqnZCR3UNeyhspdOyHTNrinzDiYCwFzyFhRbx6AQRBS8z7N4hxQHGLGjWZCzyBe75xdydHjzaMIIIoQNwZDZD"

	Url, err := url.Parse("https://graph.facebook.com/me")
	if err != nil {
		panic(err.Error())
	}
	params := url.Values{}
	params.Set("fields", "id,name,birthday,email,hometown")
	params.Set("access_token", token)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()

	fmt.Println(urlPath)

	resp, err := http.Get(urlPath)
	if err != nil {
		log.Print(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}