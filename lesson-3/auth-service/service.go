package main

import (
	"context"
	"log"

	"github.com/azomio/courses/lesson4/pkg/grpc/user"
	jwt "github.com/azomio/courses/lesson4/pkg/jwt"
)

type User struct {
	ID     int
	Name   string
	Email  string
	IsPaid bool
	Pwd    string
}

type AuthService struct {
	db map[string]*User
}

func (s *AuthService) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	u, ok := s.db[req.GetEmail()]
	if !ok {
		return &user.LoginResponse{Error: "Пользователь не найден"}, nil
	}

	if u.Pwd != req.GetPwd() {
		return &user.LoginResponse{Error: "Неправильный email или пароль"}, nil
	}

	t, err := jwt.Make(jwt.Payload{u.ID, u.Name, u.IsPaid})

	if err != nil {
		log.Printf("Token error: %v", err)
		return &user.LoginResponse{Error: "Внутренняя ошибка сервиса"}, nil
	}

	return &user.LoginResponse{Jwt: t}, nil
}
