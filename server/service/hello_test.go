package service

import (
	"context"
	"testing"
	"time"

	pb "github.com/beomdevops/go-grpc/proto"
	"github.com/stretchr/testify/assert"
)

func Test_HelloService_s(t *testing.T) {
	t.Run("성공", func(t *testing.T) {
		helloSrv := NewHelloService()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		name := "park"
		req := &pb.HelloRequest{
			Name: name,
		}
		res, err := helloSrv.SayHello(ctx, req)
		defer cancel()
		assert.Equal(t, res.GetMessage(), name)
		assert.NoError(t, err)
	})
	t.Run("실패", func(t *testing.T) {
		helloSrv := NewHelloService()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		res, err := helloSrv.SayHello(ctx, nil)
		defer cancel()
		assert.Nil(t, res)
		assert.Error(t, err)
	})
}
