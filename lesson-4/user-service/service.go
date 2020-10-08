package main

import (
	"context"
	"log"

	"github.com/Barugoo/gb-go-microservices/lesson-4/jwt"
	pb "github.com/Barugoo/gb-go-microservices/lesson-4/user-service/api"
)

type UserService struct {
}

func (s *UserService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	u := UU.GetByEmail(in.GetEmail())

	if u == nil {
		return &pb.LoginResponse{Error: "Пользователь не найден"}, nil
	}

	if u.Pwd != in.GetPwd() {
		return &pb.LoginResponse{Error: "Неправильный email или пароль"}, nil
	}

	t, err := jwt.Make(jwt.Payload{u.ID, u.Name, u.IsPaid})

	if err != nil {
		log.Printf("Token error: %v", err)
		return &pb.LoginResponse{Error: "Внутренняя ошибка сервиса"}, nil
	}

	return &pb.LoginResponse{Jwt: t}, nil
}
