package service

import (
	"context"
	user "github.com/wheelergeo/g-otter-gen/user"
	"testing"
)

func TestRpcUserCommonPostDelete_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRpcUserCommonPostDeleteService(ctx)
	// init req and assert value

	req := &user.RpcUserCommonPostDeleteReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
