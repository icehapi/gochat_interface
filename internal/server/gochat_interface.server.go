package server

import (
	"context"
	"gochat_interface/internal/service"
	pb "gochat_interface/proto"
	"log"
)

type GochatInterfaceServer struct{}

func (s *GochatInterfaceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	rsp := &pb.LoginResponse{
		Code:      0,
		Message:   "",
		Authtoken: "",
	}
	err := service.Login(ctx, req, rsp)
	log.Println(req, rsp)
	return rsp, err
}

func (s *GochatInterfaceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	rsp := &pb.RegisterResponse{
		Code:      0,
		Message:   "",
		Authtoken: "",
	}
	err := service.Register(ctx, req, rsp)
	log.Println(req, rsp)
	return rsp, err
}
