package user

import (
	pb "compus-second-hand/rpc/user/pb"

	"go-micro.dev/v5"
)

var UserMicroClient pb.UserService

func MicroClientInit() {
	service := micro.NewService()
	service.Init()
	UserMicroClient = pb.NewUserService("user", service.Client())
}
