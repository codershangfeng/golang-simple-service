package main

import (
	"github.com/codershangfeng/grpcservice/incomingadapter"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.POST("/register", func(c *gin.Context) {
		incomingadapter.Register(c)
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run()
}
