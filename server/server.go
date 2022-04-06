package server

import (
	"context"
	"log"
	"net"

	"example/api"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type grpcServer struct {
	api.ExampleServiceServer
}

func (g *grpcServer) example(ctx context.Context, in *api.ExampleRequest) (*api.ExampleResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &api.ExampleResponse{Msg: "Hello " + in.GetName(), Result: 0}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	// 将服务描述(server)及其具体实现(greeterServer)注册到 gRPC 中去.
	// 内部使用的是一个 map 结构存储，类似 HTTP server。
	api.RegisterExampleServiceServer(s, &grpcServer{})
	log.Println("Serving gRPC on 0.0.0.0" + port)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
s