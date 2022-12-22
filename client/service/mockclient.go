package service

import (
	"context"
	"errors"

	pb "github.com/beomdevops/go-grpc/proto"
	grpc "google.golang.org/grpc"
)

type MockClient struct {
}

func NewMock() pb.GreeterClient {
	return &MockClient{}
}

func (m *MockClient) SayHello(ctx context.Context, in *pb.HelloRequest, opts ...grpc.CallOption) (*pb.HelloReply, error) {

	if in == nil {
		return nil, errors.New("nil")
	}

	res := &pb.HelloReply{
		Message: in.GetName(),
	}

	return res, nil
}
