package service

import (
	"context"
	user "github.com/wheelergeo/g-otter-gen/user"
	"testing"
)

func TestRpcUserCommonDeptDelete_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRpcUserCommonDeptDeleteService(ctx)
	// init req and assert value

	req := &user.RpcUserCommonDeptDeleteReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
