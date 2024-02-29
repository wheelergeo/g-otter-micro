package service

import (
	"context"
	user "github.com/wheelergeo/g-otter-gen/user"
	"testing"
)

func TestRpcUserCommonPostCreate_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRpcUserCommonPostCreateService(ctx)
	// init req and assert value

	req := &user.RpcUserCommonPostCreateReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
