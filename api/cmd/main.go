package main

import (
	userHandler "compus-second-hand/api/handler/user"
	logInit "compus-second-hand/api/log"
	"compus-second-hand/api/model"
	"compus-second-hand/global"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	//变更gin的默认日志打印到文件
	logFileGin := initGin()
	defer logFileGin.Close()

	//初始化日志
	logFileLog := logInit.InitLog("../log/file/log.log")
	defer logFileLog.Close()

	//初始化数据库
	model.DBinit()

	//启动服务
	global.Engine = gin.Default()

	//注册路由
	initRouter()

	global.Engine.Run(":8000")
	fmt.Println("网关启动成功")
}

func initGin() *os.File {
	f, err := os.OpenFile("../log/file/gin.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("gin日志初始化失败: " + err.Error())
	}
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DisableConsoleColor()
	return f
}

func initRouter() {
	{
		v1 := global.Engine.Group("/v1")
		{
			user := v1.Group("/user")
			user.POST("/register", userHandler.Register())
			user.POST("/login", userHandler.Login())
		}
	}
}
