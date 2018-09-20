package main

import (
	"github.com/ChinasoftNobody/go-web-gin/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	util.InitLog()
	engine := gin.Default()
	util.Middleware(engine)
	engine.GET("/ping", func(context *gin.Context) {
		panic("asdasdasd")
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	engine.Run(":8080")

}
