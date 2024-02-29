package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	user "github.com/wheelergeo/g-otter-gen/user"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/model"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/mysql"
	"github.com/wheelergeo/g-otter-pkg/constants"
)

type RpcUserCommonKvCreateService struct {
	ctx context.Context
} // NewRpcUserCommonKvCreateService new RpcUserCommonKvCreateService
func NewRpcUserCommonKvCreateService(ctx context.Context) *RpcUserCommonKvCreateService {
	return &RpcUserCommonKvCreateService{ctx: ctx}
}

// Run create note info
func (s *RpcUserCommonKvCreateService) Run(req *user.RpcUserCommonKvCreateReq) (resp *user.RpcUserCommonKvCreateResp, err error) {
	err = req.IsValid()
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.ParamInvalid
		resp.Status.Message = constants.ParamInvalidMsg
		return
	}

	resp = new(user.RpcUserCommonKvCreateResp)
	resp.Status = new(user.BaseResp)
	kvModel := model.UserCommonKv{
		Key:   req.Key,
		Value: req.Value,
	}

	err = mysql.DB.Create(&kvModel).Error
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.UserCommonKvCreateFailed
		resp.Status.Message = constants.UserCommonKvCreateFailedMsg
		return
	}

	resp.Status.Code = constants.UserCommonKvCreateSuccess
	resp.Status.Message = constants.UserCommonKvCreateSuccessMsg
	return
}
