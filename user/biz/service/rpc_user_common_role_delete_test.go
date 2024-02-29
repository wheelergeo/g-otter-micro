package service

import (
	"context"
	user "github.com/wheelergeo/g-otter-gen/user"
	"testing"
)

func TestRpcUserCommonRoleDelete_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRpcUserCommonRoleDeleteService(ctx)
	// init req and assert value

	req := &user.RpcUserCommonRoleDeleteReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
