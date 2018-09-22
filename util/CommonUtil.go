package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
//asdasdasda
//asdasdasd
func Middleware(engine *gin.Engine) {
	engine.Use(errorHandle())
}

func errorHandle() func(c *gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				LogError(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": "Fail",
					"reason": err,
				})
			}
		}()
		LogInfo("调用接口")
		c.Next()
		LogInfo("调用接口结束")
	}
}
