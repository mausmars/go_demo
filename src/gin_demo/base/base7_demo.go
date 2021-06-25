package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) {
			firstname := c.DefaultQuery("firstname", "Guest")
			c.String(http.StatusOK, "Hello %s %s", firstname, firstname)
		})
		v1.POST("/submit", func(c *gin.Context) {})
		v1.POST("/read", func(c *gin.Context) {})
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", func(c *gin.Context) {})
		v2.POST("/submit", func(c *gin.Context) {})
		v2.POST("/read", func(c *gin.Context) {})
	}

	router.Run(":8099")
}