package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	user "github.com/wheelergeo/g-otter-gen/user"
	"go.opentelemetry.io/otel/baggage"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	klog.CtxInfof(s.ctx, "baggage:%s", baggage.FromContext(s.ctx).String())
	klog.CtxInfof(s.ctx, "phone:%s, password:%s",
		req.PhoneNum, req.Password)

	return
}
