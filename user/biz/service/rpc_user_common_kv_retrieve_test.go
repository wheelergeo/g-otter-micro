package service

import (
	"context"
	user "github.com/wheelergeo/g-otter-gen/user"
	"testing"
)

func TestRpcUserCommonKvRetrieve_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRpcUserCommonKvRetrieveService(ctx)
	// init req and assert value

	req := &user.RpcUserCommonKvRetrieveReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
