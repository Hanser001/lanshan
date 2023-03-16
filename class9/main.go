package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	hub := newHub()
	go hub.run()
	r.GET("/websocket", func(c *gin.Context) {
		websocketFc(c, hub)
	})
	r.Run(":8080")
}
