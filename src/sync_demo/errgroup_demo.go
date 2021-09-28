package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

//eg errgroup.Group
//如果返回错误 — 这一组 Goroutine 最少返回一个错误；
//如果返回空值 — 所有 Goroutine 都成功执行；

func main() {
	var eg errgroup.Group
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.ddddd.com/",
		"http://www.somestupidname.com/",
	}
	for i := range urls {
		url := urls[i]
		eg.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			fmt.Println("http request ", url)
			return err
		})
	}
	if err := eg.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Println("Fail ", err.Error())
	}
}
