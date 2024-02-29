package service

import (
	"context"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
	user "github.com/wheelergeo/g-otter-gen/user"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/model"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/mysql"
	"github.com/wheelergeo/g-otter-pkg/constants"
)

type RpcUserCommonPostUpdateService struct {
	ctx context.Context
} // NewRpcUserCommonPostUpdateService new RpcUserCommonPostUpdateService
func NewRpcUserCommonPostUpdateService(ctx context.Context) *RpcUserCommonPostUpdateService {
	return &RpcUserCommonPostUpdateService{ctx: ctx}
}

// Run create note info
func (s *RpcUserCommonPostUpdateService) Run(req *user.RpcUserCommonPostUpdateReq) (resp *user.RpcUserCommonPostUpdateResp, err error) {
	err = req.IsValid()
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.ParamInvalid
		resp.Status.Message = constants.ParamInvalidMsg
		return
	}
	resp = new(user.RpcUserCommonPostUpdateResp)
	resp.Status = new(user.BaseResp)

	postModel := model.UserCommonPost{}
	err = mysql.DB.First(&postModel, req.Id).Error
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.UserCommonPostUpdateFailed
		resp.Status.Message = constants.UserCommonPostUpdateFailedMsg
		return
	}

	if req.PostCode != nil {
		postModel.PostCode = *req.PostCode
	}
	if req.PostName != nil {
		postModel.PostName = *req.PostName
	}
	if req.Level != nil {
		postModel.Level = int32(*req.Level)
	}
	if req.Status != nil {
		postModel.Status = int32(*req.Status)
	}
	if req.Remark != nil {
		postModel.Remark = *req.Remark
	}
	if len(req.RoleIds) != 0 {
		roleIds := ""
		for _, v := range req.RoleIds {
			roleIds += strconv.Itoa(int(v)) + ","
		}
		postModel.RoleIds = roleIds
	}

	err = mysql.DB.Save(&postModel).Error
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.UserCommonPostUpdateFailed
		resp.Status.Message = constants.UserCommonPostUpdateFailedMsg
		return
	}

	resp.Status.Code = constants.UserCommonPostUpdateSuccess
	resp.Status.Message = constants.UserCommonPostUpdateSuccessMsg
	return
}
