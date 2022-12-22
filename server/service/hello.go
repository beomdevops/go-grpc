package service

import (
	"context"
	"errors"
	"log"

	pb "github.com/beomdevops/go-grpc/proto"
)

type HelloService struct {
	pb.UnimplementedGreeterServer
}

func NewHelloService() *HelloService {
	return &HelloService{}
}

func (s *HelloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())

	if in == nil {
		return nil, errors.New("error")
	}

	return &pb.HelloReply{Message: in.GetName()}, nil
}
