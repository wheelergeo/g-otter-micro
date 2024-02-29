package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	user "github.com/wheelergeo/g-otter-gen/user"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/model"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/mysql"
	"github.com/wheelergeo/g-otter-pkg/constants"
)

type RpcUserCommonKvDeleteService struct {
	ctx context.Context
} // NewRpcUserCommonKvDeleteService new RpcUserCommonKvDeleteService
func NewRpcUserCommonKvDeleteService(ctx context.Context) *RpcUserCommonKvDeleteService {
	return &RpcUserCommonKvDeleteService{ctx: ctx}
}

// Run create note info
func (s *RpcUserCommonKvDeleteService) Run(req *user.RpcUserCommonKvDeleteReq) (resp *user.RpcUserCommonKvDeleteResp, err error) {
	err = req.IsValid()
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.ParamInvalid
		resp.Status.Message = constants.ParamInvalidMsg
		return
	}
	resp = new(user.RpcUserCommonKvDeleteResp)
	resp.Status = new(user.BaseResp)

	err = mysql.DB.Where("key IN ?", req.Keys).Delete(&model.UserCommonKv{}).Error
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.UserCommonKvDeleteFailed
		resp.Status.Message = constants.UserCommonKvDeleteFailedMsg
		return
	}

	resp.Status.Code = constants.UserCommonKvDeleteSuccess
	resp.Status.Message = constants.UserCommonKvDeleteSuccessMsg
	return
}
