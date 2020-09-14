package main

import (
	"context"

	"github.com/go-xorm/xorm"

	pb "user-service/api"
)

type Service interface {
	pb.UserServiceServer
	CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error)
	GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error)
}

type service struct {
	pb.UnimplementedUserServiceServer
	db *xorm.Engine
}

func (s *service) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	user := &User{
		Email:    req.Email,
		Password: req.Pwd,
	}
	_, err := s.db.Insert(user)
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:    user.ID,
		Email: user.Email,
		Pwd:   user.Password,
	}, err
}

func (s *service) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	var user User
	_, err := s.db.Where("email = ?", req.Email).Get(&user)
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:    user.ID,
		Email: user.Email,
		Pwd:   user.Password,
	}, err
}

func NewService(db *xorm.Engine) Service {
	return &service{
		db: db,
	}
}
