package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	user "github.com/wheelergeo/g-otter-gen/user"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/model"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/mysql"
	"github.com/wheelergeo/g-otter-pkg/constants"
)

type RpcUserCommonPostDeleteService struct {
	ctx context.Context
} // NewRpcUserCommonPostDeleteService new RpcUserCommonPostDeleteService
func NewRpcUserCommonPostDeleteService(ctx context.Context) *RpcUserCommonPostDeleteService {
	return &RpcUserCommonPostDeleteService{ctx: ctx}
}

// Run create note info
func (s *RpcUserCommonPostDeleteService) Run(req *user.RpcUserCommonPostDeleteReq) (resp *user.RpcUserCommonPostDeleteResp, err error) {
	err = req.IsValid()
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.ParamInvalid
		resp.Status.Message = constants.ParamInvalidMsg
		return
	}
	resp = new(user.RpcUserCommonPostDeleteResp)
	resp.Status = new(user.BaseResp)

	err = mysql.DB.Delete(&model.UserCommonPost{}, req.Ids).Error
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.UserCommonPostDeleteFailed
		resp.Status.Message = constants.UserCommonPostDeleteFailedMsg
		return
	}

	resp.Status.Code = constants.UserCommonPostDeleteSuccess
	resp.Status.Message = constants.UserCommonPostDeleteSuccessMsg
	return
}
