package controller

import (
	"context"

	"github.com/Anhbman/microservice-server-cake/internal/server/cake"
	"github.com/Anhbman/microservice-server-cake/rpc/service"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ServiceServer struct {
	cakeProcessor *cake.Processor
}

var _ service.Service = (*ServiceServer)(nil)

func NewServiceServer(cakeProcessor *cake.Processor) *ServiceServer {
	return &ServiceServer{cakeProcessor: cakeProcessor}
}

func (s *ServiceServer) CreateCake(ctx context.Context, req *service.CreateCakeRequest) (*emptypb.Empty, error) {
	return s.cakeProcessor.Create(ctx, req)
}

func (s *ServiceServer) GetCakeById(ctx context.Context, req *service.GetCakeByIdRequest) (*service.GetCakeByIdResponse, error) {
	return s.cakeProcessor.GetCakeById(ctx, req)
}
