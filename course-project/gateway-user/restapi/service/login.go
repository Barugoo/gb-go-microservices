package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	pb "github.com/barugoo/gb-go-microservices/course-project/gateway-user/cmd/proto"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/barugoo/gb-go-microservices/course-project/gateway-user/models"
	"github.com/barugoo/gb-go-microservices/course-project/gateway-user/restapi/operations/auth"
)

func (s *GatewayUserService) Login(params auth.LoginUserParams) (response middleware.Responder) {
	resp, err := s.login(context.Background(), params)
	if err != nil {
		log.Println(err)
		return middleware.ResponderFunc(func(rw http.ResponseWriter, pr runtime.Producer) {
			rw.WriteHeader(500)
		})
	}
	bytes, err := json.Marshal(resp)
	if err != nil {
		return nil
	}
	return middleware.ResponderFunc(func(rw http.ResponseWriter, pr runtime.Producer) {
		rw.WriteHeader(200)
		rw.Write(bytes)
	})
}

func (s *GatewayUserService) login(ctx context.Context, params auth.LoginUserParams) (*models.LoginResponse, error) {
	req := &pb.LoginRequest{
		Email:    params.Body.Email,
		Password: params.Body.Password,
	}
	resp, err := s.AuthServiceClient.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	payload := &models.LoginResponse{
		JWT: resp.Jwt,
	}
	return payload, nil
}
