package service

import (
	pb "github.com/barugoo/gb-go-microservices/course-project/gateway-user/cmd/proto"
)

type GatewayUserService struct {
	AuthServiceClient pb.AuthServiceClient
}
