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

type RpcUserCommonPostCreateService struct {
	ctx context.Context
} // NewRpcUserCommonPostCreateService new RpcUserCommonPostCreateService
func NewRpcUserCommonPostCreateService(ctx context.Context) *RpcUserCommonPostCreateService {
	return &RpcUserCommonPostCreateService{ctx: ctx}
}

// Run create note info
func (s *RpcUserCommonPostCreateService) Run(req *user.RpcUserCommonPostCreateReq) (resp *user.RpcUserCommonPostCreateResp, err error) {
	err = req.IsValid()
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.ParamInvalid
		resp.Status.Message = constants.ParamInvalidMsg
		return
	}

	resp = new(user.RpcUserCommonPostCreateResp)
	resp.Status = new(user.BaseResp)
	roleIds := ""

	for _, v := range req.RoleIds {
		roleIds += strconv.Itoa(int(v)) + ","
	}

	postModel := model.UserCommonPost{
		PostCode: req.PostCode,
		PostName: req.PostName,
		Level:    int32(req.Level),
		Status:   int32(req.Status),
		RoleIds:  roleIds,
	}

	if req.Remark != nil {
		postModel.Remark = *req.Remark
	}

	err = mysql.DB.Create(&postModel).Error
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.UserCommonPostCreateFailed
		resp.Status.Message = constants.UserCommonPostCreateFailedMsg
		return
	}

	resp.Status.Code = constants.UserCommonPostCreateSuccess
	resp.Status.Message = constants.UserCommonPostCreateSuccessMsg
	return
}
