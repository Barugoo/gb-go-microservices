package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"log"

	pb "github.com/barugoo/gb-go-microservices/course-project/auth-service/api"
	jwt "github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	db map[string]*User
}

func NewAuthService() pb.AuthServiceServer {
	return &AuthService{
		db: map[string]*User{
			"barugoo@yandex.ru": &User{
				Email:    "barugoo@yandex.ru",
				Password: "77de38e4b50e618a0ebb95db61e2f42697391659d82c064a5f81b9f48d85ccd5",
			},
		},
	}
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(req.Password))
	password := string(mac.Sum(nil))

	newUser := &User{
		Email:    req.Email,
		Password: password,
	}
	s.db[newUser.Email] = newUser

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": newUser.Email,
	})
	tokenString, _ := token.SignedString(secret)

	return &pb.RegisterResponse{
		Ok:  true,
		Jwt: tokenString,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	u, ok := s.db[req.GetEmail()]
	if !ok {
		return nil, nil
	}

	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(req.Password))
	inputPassword := string(mac.Sum(nil))

	log.Println(inputPassword)

	if inputPassword != u.Password {
		return nil, ErrUnauthorized
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
	})
	tokenString, _ := token.SignedString(secret)

	return &pb.LoginResponse{
		Ok:  true,
		Jwt: tokenString,
	}, nil
}
