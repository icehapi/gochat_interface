package main

import (
	"context"
	"gochat_interface/internal/server"
	pb "gochat_interface/proto"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"google.golang.org/grpc"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

func run() {
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard,
			&runtime.JSONPb{OrigName: true, EmitDefaults: true}),
	)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	log.Println("start server...")
	err := pb.RegisterGochatInterfaceHandlerFromEndpoint(context.Background(), mux, "localhost:7700", opts)
	if err != nil {
		log.Fatal(err)
	}

	rpcServer := grpc.NewServer()
	pb.RegisterGochatInterfaceServer(rpcServer, &server.GochatInterfaceServer{})
	http.ListenAndServe(
		":7700",
		grpcHandlerFunc(rpcServer, mux),
	)
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func main() {
	run()
}
