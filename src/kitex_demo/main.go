package main

import (
	"go_demo/src/kitex_demo/impl"
	api "go_demo/src/kitex_demo/kitex_gen/api/hello"
	"log"
)

func main() {
	svr := api.NewServer(new(impl.HelloImpl))
	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
