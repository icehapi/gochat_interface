package service

import (
	"context"
	"gochat_interface/internal/rpc"
	pb "gochat_interface/proto"
	"log"
)

func Login(ctx context.Context, req *pb.LoginRequest, rsp *pb.LoginResponse) error {
	// rpc调用完成登录
	rpcReq := &pb.GochatDBLoginRequest{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	}
	rpcRsp, err := rpc.RpcLogin(ctx, rpcReq)
	if err != nil {
		log.Fatal("rpc err:", err)
		return err
	}
	rsp.Authtoken = rpcRsp.Authtoken
	// 存入redis缓存
	return nil
}

func Register(ctx context.Context, req *pb.RegisterRequest, rsp *pb.RegisterResponse) error {
	// rpc调用下级服务完成注册
	rpcReq := &pb.GochatDBRegisterRequest{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	}
	rpcRsp, err := rpc.RpcRegister(ctx, rpcReq)
	if err != nil {
		log.Fatal("rpc err:", err)
		return err
	}
	rsp.Authtoken = rpcRsp.Authtoken
	// 无错误则将 token 存入redis
	return nil
}
