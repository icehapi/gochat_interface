`proto` 编译：
```bash
protoc -I. -I$GOPATH/src -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --go_out=plugins=grpc:. gochat_interface.proto

protoc -I. -I$GOPATH/src -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --grpc-gateway_out=logtostderr=true:. gochat_interface.proto
```