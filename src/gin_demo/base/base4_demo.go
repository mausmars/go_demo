package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//Multipart/Urlencoded Form
func main() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		fmt.Println("message=",message)


		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8099")
}