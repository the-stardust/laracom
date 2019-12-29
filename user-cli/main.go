package main

import (
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	pb "github.com/the-stardust/laracom/user-service/proto/user"
)

func main() {
	// 初始化客户端服务，定义命令行参数标识
	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "Your Name",
			},
			cli.StringFlag{
				Name:  "email",
				Usage: "Your Email",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "Your Password",
			},
		),
	)
	// 远程服务客户端调用句柄
	client := pb.NewUserServiceClient()
}
