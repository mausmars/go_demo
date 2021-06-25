package main

import (
	"log"
	"github.com/gin-gonic/gin"
	. "go_demo/src/gin_demo/ginmodel"
)

func main() {
	route := gin.Default()
	route.GET("/testing", startPage2)
	route.Run(":8099")
}

func startPage2(c *gin.Context) {
	var person Person2
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
		log.Println(person.CreateTime)
		log.Println(person.UnixTime)
	}

	c.String(200, "Success")
}