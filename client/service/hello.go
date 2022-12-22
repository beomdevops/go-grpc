package service

import (
	"context"
	"time"

	pb "github.com/beomdevops/go-grpc/proto"
)

type HeleloService struct {
	GreetSrv pb.GreeterClient
}

func NewHelloService(di_client pb.GreeterClient) *HeleloService {
	return &HeleloService{
		GreetSrv: di_client,
	}
}

func (h *HeleloService) GetGrpcMessage(msg string) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	name := &pb.HelloRequest{
		Name: msg,
	}
	res, err := h.GreetSrv.SayHello(ctx, name)

	if err != nil {
		return ""
	}
	return res.GetMessage()
}
