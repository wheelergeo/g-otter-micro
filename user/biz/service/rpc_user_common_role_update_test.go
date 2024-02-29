package service

import (
	"context"
	user "github.com/wheelergeo/g-otter-gen/user"
	"testing"
)

func TestRpcUserCommonRoleUpdate_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRpcUserCommonRoleUpdateService(ctx)
	// init req and assert value

	req := &user.RpcUserCommonRoleUpdateReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
