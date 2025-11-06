package main

import (
	"compus-second-hand/api/model"
	"compus-second-hand/global"
	"compus-second-hand/rpc/user/handler"
	pb "compus-second-hand/rpc/user/pb"
	"strconv"

	"go-micro.dev/v5"
)

func main() {
	//初始化数据库
	model.DBinit()
	//启动微服务
	svc := micro.NewService(
		micro.Name("user"),
		micro.Address(":"+strconv.Itoa(global.Configs.Server.UserPort)),
	)
	if err := pb.RegisterUserHandler(svc.Server(), new(handler.UserHandlerEr)); err != nil {
		panic("user微服务注册失败")
	}
	if err := svc.Run(); err != nil {
		panic("user微服务启动失败")
	}
}
