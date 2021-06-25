package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	token:="EAAKfdeoZBLyYBAKm3woWs7OmUJvgXZBcze4MccHicS6AelSRGb47UTsw8aFY2528CoBOX7lmRmRODVVcoQV2TSuWkmcWAUeRB8mSsANPAa2wcan1tjAZAcXuJC7KSf48VileCrgn4RbqLyNvaq8INhyLeOzUgMDCCAcIiE2XcwXu72RUOlkOKams4qcMdG7X68l4hZBRAgZDZD"

	Url, err := url.Parse("https://graph.facebook.com/me")
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