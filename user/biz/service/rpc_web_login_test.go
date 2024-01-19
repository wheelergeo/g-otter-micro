package service

import (
	"context"
	user "github.com/wheelergeo/g-otter-gen/user"
	"testing"
)

func TestRpcWebLogin_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRpcWebLoginService(ctx)
	// init req and assert value

	req := &user.RpcWebLoginReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
