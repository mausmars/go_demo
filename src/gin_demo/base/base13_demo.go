package main

import (
	"log"
	"github.com/gin-gonic/gin"
	. "go_demo/src/gin_demo/ginmodel"
)

func main() {
	route := gin.Default()
	route.Any("/testing", startPage)
	route.Run(":8099")
}

func startPage(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}