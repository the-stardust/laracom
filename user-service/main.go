package main

import (
	"github.com/micro/go-micro"
	database "github.com/the-stardust/laracom/user-service/db"
	"github.com/the-stardust/laracom/user-service/handler"
	pb "github.com/the-stardust/laracom/user-service/proto/user"
	repository "github.com/the-stardust/laracom/user-service/repo"
	"github.com/the-stardust/laracom/user-service/service"
)

func main() {
	db, err := database.CreateConnection()
	defer db.Close()
	if err != nil {
		panic(err)
	}

	// 和 Laravel 数据库迁移类似
	// 每次启动服务时都会检查，如果数据表不存在则创建，已存在检查是否有修改
	db.AutoMigrate(&pb.User{})
	// 初始化 Repo 实例用于后续数据库操作
	userRepository := &repository.UserRepository{db}
	// 初始化 token service
	token := &service.TokenService{userRepository}
	// 以下是 Micro 创建微服务流程
	srv := micro.NewService(
		micro.Name("laracom.user.service"),
		micro.Version("latest"),
		)
	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(),&handler.UserService{userRepository,token})

	if err := srv.Run();err != nil {
		panic(err)
	}


}
