package main

import (
	"context"
	"net"

	proto "github.com/codepandaid/tutorial-go/proto"
	grpc "google.golang.org/grpc"
)

type srv struct {
}

func main() {
	listener, err := net.Listen("tcp", ":4321")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	proto.RegisterAddserviceServer(server, &srv{})
	// reflection.Register(server)

	if e := server.Serve(listener); e != nil {
		panic(err)
	}
}

func (s *srv) Add(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	a, b := req.GetA(), req.GetB()

	result := a + b
	return &proto.Response{Result: result}, nil
}

func (s *srv) Multiply(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	a, b := req.GetA(), req.GetB()

	result := a * b
	return &proto.Response{Result: result}, nil
}
