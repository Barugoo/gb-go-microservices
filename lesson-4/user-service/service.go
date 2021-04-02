package main

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "user-service/api"

	"user-service/pkg/jwt"
)

type UserService struct {
	pb.UnimplementedUserServer
	db *UserStorage
}

func (s *UserService) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	u, err := s.db.GetByEmail(ctx, in.GetEmail())
	if err != nil && err != ErrNotFound {
		log.Println("get user by email", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}
	if err == nil {
		return nil, status.Error(codes.AlreadyExists, "user already exists with email")
	}

	newUser := &User{
		Email: u.Email,
		Pwd:   u.Pwd,
		Name:  u.Name,
	}
	usr, err := s.db.CreateUser(ctx, newUser)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	t, err := jwt.Make(jwt.Payload{u.ID, u.Name, u.IsPaid})
	if err != nil {
		log.Printf("Token error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	resp := &pb.RegisterResponse{
		Id:     usr.ID,
		Name:   usr.Name,
		Email:  usr.Email,
		IsPaid: usr.IsPaid,
		Token:  t,
	}
	return resp, nil
}

func (s *UserService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	u, err := s.db.GetByEmail(ctx, in.GetEmail())
	if err != nil {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	if u.Pwd != in.GetPwd() {
		return nil, status.Error(codes.Unauthenticated, "Wrong creds")
	}

	t, err := jwt.Make(jwt.Payload{u.ID, u.Name, u.IsPaid})

	if err != nil {
		log.Printf("Token error: %v", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.LoginResponse{Token: t}, nil
}
