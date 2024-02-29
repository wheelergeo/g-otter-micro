package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	user "github.com/wheelergeo/g-otter-gen/user"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/model"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/mysql"
	"github.com/wheelergeo/g-otter-pkg/constants"
)

type RpcUserCommonKvUpdateService struct {
	ctx context.Context
} // NewRpcUserCommonKvUpdateService new RpcUserCommonKvUpdateService
func NewRpcUserCommonKvUpdateService(ctx context.Context) *RpcUserCommonKvUpdateService {
	return &RpcUserCommonKvUpdateService{ctx: ctx}
}

// Run create note info
func (s *RpcUserCommonKvUpdateService) Run(req *user.RpcUserCommonKvUpdateReq) (resp *user.RpcUserCommonKvUpdateResp, err error) {
	err = req.IsValid()
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.ParamInvalid
		resp.Status.Message = constants.ParamInvalidMsg
		return
	}
	resp = new(user.RpcUserCommonKvUpdateResp)
	resp.Status = new(user.BaseResp)

	kvModel := model.UserCommonKv{}
	err = mysql.DB.First(&kvModel, req.Key).Error
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.UserCommonKvUpdateFailed
		resp.Status.Message = constants.UserCommonKvUpdateFailedMsg
		return
	}

	kvModel.Value = req.Value

	err = mysql.DB.Save(&kvModel).Error
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.UserCommonKvUpdateFailed
		resp.Status.Message = constants.UserCommonKvUpdateFailedMsg
		return
	}

	resp.Status.Code = constants.UserCommonKvUpdateSuccess
	resp.Status.Message = constants.UserCommonKvUpdateSuccessMsg
	return
}
