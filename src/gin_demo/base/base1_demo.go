package main

import "github.com/gin-gonic/gin"
//Using GET, POST, PUT, PATCH, DELETE and OPTIONS
func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/someGet", func(c *gin.Context) {})
	router.POST("/somePost",  func(c *gin.Context) {})
	router.PUT("/somePut",  func(c *gin.Context) {})
	router.DELETE("/someDelete",  func(c *gin.Context) {})
	router.PATCH("/somePatch",  func(c *gin.Context) {})
	router.HEAD("/someHead",  func(c *gin.Context) {})
	router.OPTIONS("/someOptions",  func(c *gin.Context) {})

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run(":8090")
	// router.Run(":3000") for a hard coded port
}
