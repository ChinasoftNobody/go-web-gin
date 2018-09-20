package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func InitLog() {
	infoWriter, err := os.Create("logs\\gin.log")
	errWriter, err := os.Create("logs\\gin_err.log")
	if err != nil {
		panic("日至初始化失败")
	}
	gin.DefaultWriter = io.MultiWriter(infoWriter, os.Stdout)
	gin.DefaultErrorWriter = io.MultiWriter(errWriter, os.Stderr)
}

func LogInfo(args ...interface{}) {
	fmt.Fprintln(gin.DefaultWriter, args)
}

func LogError(args ...interface{}) {
	fmt.Fprintln(gin.DefaultErrorWriter, args)
}
