package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_test(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		srv := NewHelloService(NewMock())
		req := "test"
		ret := srv.GetGrpcMessage(req)
		assert.Equal(t, ret, req)
	})
	t.Run("실패", func(t *testing.T) {
		srv := NewHelloService(NewMock())
		req := ""
		ret := srv.GetGrpcMessage(req)
		assert.Equal(t, ret, req)
	})

}
