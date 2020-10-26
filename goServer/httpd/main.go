package main

import (
	"mulTic/httpd/handler"
	"mulTic/platform/mulTic"

	"github.com/gin-gonic/gin"
)

func main() {
	status := mulTic.New()
	r := gin.Default()

	r.GET("/board", handler.BoardGet(status))
	r.POST("/board", handler.BoardPost(status))
	r.POST("/boardNew", handler.BoardNew(status))

	r.Run()
}
