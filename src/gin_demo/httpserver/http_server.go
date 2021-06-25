package main

import (
	"go_demo/src/gin_demo/httpserver/json"
	protomsg "go_demo/src/gin_demo/httpserver/proto"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")

		//if pusher := c.Writer.Pusher(); pusher != nil {
		//	// use pusher.Push() to do server push
		//	if err := pusher.Push("/assets/app.js", nil); err != nil {
		//		log.Printf("Failed to push: %v", err)
		//	} else {
		//		log.Printf("Success to push!!!")
		//	}
		//}
	})
	router.POST("/json", func(c *gin.Context) {
		msg := &jsonmsg.Command{}
		if err := c.ShouldBindJSON(msg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(msg)

		reMsg := &jsonmsg.Command{
			CommandId: 2,
			Body:      "ok",
		}
		data, _ := json.Marshal(reMsg)

		c.Writer.Write(data)
	})
	router.POST("/proto", func(c *gin.Context) {
		//c.Request.Body
		msg := &protomsg.Command{}
		body, _ := ioutil.ReadAll(c.Request.Body)
		proto.Unmarshal(body, msg)

		//fmt.Println(msg)

		if msg.CommandId == 1 {
			reMsg := &protomsg.Command{
				CommandId: int32(1),
				Body:      msg.Body,
			}
			reBody, _ := proto.Marshal(reMsg)
			//
			c.Writer.Write(reBody)
		} else if msg.CommandId == 2 {
			reMsg := &protomsg.Command{
				CommandId: int32(2),
				Body:      []byte("push return"),
			}
			reBody, _ := proto.Marshal(reMsg)
			//
			//println("wait start!!!")
			//time.Sleep(3600*time.Second)
			c.Writer.Write(reBody)
			//println("wait end!!!")

			//go func(c *gin.Context) {
			//
			//}(c)
		}
	})
	router.Run(":8008")
}
