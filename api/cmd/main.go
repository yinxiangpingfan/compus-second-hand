package main

import (
	logInit "compus-second-hand/api/log"
	"compus-second-hand/global"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	//初始化gin的运行日志
	f, err := os.OpenFile("../log/file/gin.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("gin日志初始化失败: " + err.Error())
	}
	defer f.Close()
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DisableConsoleColor()

	//初始化日志
	logfile, err := os.OpenFile("../log/file/log.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("log日志初始化失败: " + err.Error())
	}
	defer logfile.Close()
	logInit.LogerInit(logfile)

	global.Logger.Error("hello world", "error", "test")

	//启动服务
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "hello world"})
	})
	r.Run(":8000")
}
