package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "go_demo/src/gin_demo/ginmodel"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		h := TestHeader{}

		//if err := c.ShouldBindHeader(&h); err != nil {
		//	c.JSON(200, err)
		//}

		fmt.Printf("%#v\n", h)
		c.JSON(200, gin.H{"Rate": h.Rate, "Domain": h.Domain})
	})

	r.Run()

	// client
	// curl -H "rate:300" -H "domain:music" 127.0.0.1:8080/
	// output
	// {"Domain":"music","Rate":300}
}
