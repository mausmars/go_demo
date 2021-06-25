package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	token:="EAAKfdeoZBLyYBAEsVaAgkTJxBhMK9bnCg9slHZAl6rwD0ROmQAIcUTbxZBKYqPzENgPWffZBq7Vw3yKnVg07Olt0GsTMvbCRIuszIPuck2GjFh8BtkjyHaegQf9RgIRV7jeZAbwzMHZBrlygCxdLne15ros2CPuGvNPE4QakcqkFGTEgHbztP7SeLRG9cPUYXv3bFLPHbzegZDZD"

	Url, err := url.Parse("https://graph.facebook.com/me/friends")
	if err != nil {
		panic(err.Error())
	}
	params := url.Values{}
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