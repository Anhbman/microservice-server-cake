package controller

import (
	"context"

	"github.com/Anhbman/microservice-server-cake/internal/server/cake"
	"github.com/Anhbman/microservice-server-cake/rpc/service"
)

type ControllerServer struct {
	cakeProcessor *cake.Processor
}

var _ service.Service = (*ControllerServer)(nil)

func NewControllerServer(cakeProcessor *cake.Processor) *ControllerServer {
	return &ControllerServer{cakeProcessor: cakeProcessor}
}

func (s *ControllerServer) CreateCake(ctx context.Context, req *service.CreateCakeRequest) (*service.Cake, error) {
	return s.cakeProcessor.Create(ctx, req)
}

func (s *ControllerServer) GetCakeById(ctx context.Context, req *service.GetCakeByIdRequest) (*service.GetCakeByIdResponse, error) {
	return s.cakeProcessor.GetCakeById(ctx, req)
}

func (s *ControllerServer) SearchCake(ctx context.Context, req *service.SearchCakeRequest) (*service.SearchCakeResponse, error) {
	return s.cakeProcessor.SearchCake(ctx, req)
}

func (s *ControllerServer) UpdateCake(ctx context.Context, req *service.Cake) (*service.Cake, error) {
	return s.cakeProcessor.UpdateCake(ctx, req)
}
