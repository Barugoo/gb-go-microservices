package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	pb "github.com/barugoo/gb-go-microservices/course-project/gateway-user/cmd/proto"
	"github.com/barugoo/gb-go-microservices/course-project/gateway-user/models"
	"github.com/barugoo/gb-go-microservices/course-project/gateway-user/restapi/operations/auth"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

func (s *GatewayUserService) Register(params auth.RegisterUserParams) (response middleware.Responder) {
	resp, err := s.register(context.Background(), params)
	if err != nil {
		log.Println(err)
		return middleware.ResponderFunc(func(rw http.ResponseWriter, pr runtime.Producer) {
			rw.WriteHeader(500)
		})
	}
	bytes, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return middleware.ResponderFunc(func(rw http.ResponseWriter, pr runtime.Producer) {
			rw.WriteHeader(500)
		})
	}
	return middleware.ResponderFunc(func(rw http.ResponseWriter, pr runtime.Producer) {
		rw.WriteHeader(200)
		rw.Write(bytes)
	})
}

func (s *GatewayUserService) register(ctx context.Context, params auth.RegisterUserParams) (*models.RegisterResponse, error) {
	req := &pb.RegisterRequest{
		Email:    params.Body.Email,
		Password: params.Body.Password,
	}
	resp, err := s.AuthServiceClient.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	payload := &models.RegisterResponse{
		JWT: resp.Jwt,
	}
	return payload, nil
}
