package main

import (
	"context"
	"log"

	"github.com/codepandaid/tutorial-go/proto"
	grpc "google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4321", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	client := proto.NewAddserviceClient(conn)

	req := &proto.Request{A: 12, B: 1}
	res, err := client.Add(ctx, req)
	if err != nil {
		log.Println(err)
	}
	log.Println(res)
}
