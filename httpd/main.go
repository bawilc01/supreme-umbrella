package main

import (
	"fmt"
	
	"newsfeeder/handler"

	"newsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", handler.PingGet())

	r.Run()

}
