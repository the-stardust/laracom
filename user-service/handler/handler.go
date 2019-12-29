package handler

import (
	"context"
	pb "github.com/the-stardust/laracom/user-service/proto/user"
	"github.com/the-stardust/laracom/user-service/repo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo repo.UserRepository
}

func (srv *UserService) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := srv.Repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (srv *UserService) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := srv.Repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (srv *UserService) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	// 密码加密
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hash)
	if err := srv.Repo.Create(req); err != nil {
		return err
	}
	res.User = req
	return nil
}
