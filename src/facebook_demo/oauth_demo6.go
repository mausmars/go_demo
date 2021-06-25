package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	token:="EAAKfdeoZBLyYBADYzo0zK6rHAnmEt9sDgUyIAeWGnM3kHQpCZBYe9Y3xABw78h6Oa19xZApvxPn5qsL02LZAfmjJEdU4LIIDFtAwnPbPWHfCRquQUHerp6dCCvgfSFOm4Is4FudjJEvhXirDV4koP0wcw876srVaU5agQjMubGAwIDnREZCGpZCuZBc4fCZAVxBdxWQ4KE5wKgZDZD"

	Url, err := url.Parse("https://graph.facebook.com/me/photos")
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