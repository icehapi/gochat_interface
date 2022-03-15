package rpc

import (
	"context"
	pb "gochat_interface/proto"
	"log"

	"google.golang.org/grpc"
)

const (
	address = ":8088"
)

func getRpcClient() pb.GochatDBClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("dial conn err: %v", err)
	}
	defer conn.Close()

	return pb.NewGochatDBClient(conn)
}

func RpcLogin(ctx context.Context, req *pb.GochatDBLoginRequest) (*pb.GochatDBLoginResponse, error) {
	client := getRpcClient()
	rsp, err := client.Login(ctx, req)
	return rsp, err
}

func RpcRegister(ctx context.Context, req *pb.GochatDBRegisterRequest) (*pb.GochatDBRegisterResponse, error) {
	client := getRpcClient()
	rsp, err := client.Register(ctx, req)
	return rsp, err
}
