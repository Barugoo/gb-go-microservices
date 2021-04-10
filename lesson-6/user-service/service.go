package main

import (
	"context"
	"log"

	pb "user-service/api"

	"github.com/azomio/courses/lesson4/pkg/jwt"
)

type UserService struct {
}

func (s *UserService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	u := UU.GetByEmail(in.GetEmail())

	if u == nil {
		return &pb.LoginResponse{Error: "Пользователь не найден"}, nil
	}

	if u.Pwd != in.GetPwd() {
		logger.Warnf(ctx, "wrong password for user %s", in.Email)
		return &pb.LoginResponse{Error: "Неправильный email или пароль"}, nil
	}
	logger.Infof(ctx, "correct password %s ", in.Email)

	t, err := jwt.Make(jwt.Payload{u.ID, u.Name, u.IsPaid})

	if err != nil {
		log.Printf("Token error: %v", err)
		return &pb.LoginResponse{Error: "Внутренняя ошибка сервиса"}, nil
	}

	return &pb.LoginResponse{Jwt: t}, nil
}
