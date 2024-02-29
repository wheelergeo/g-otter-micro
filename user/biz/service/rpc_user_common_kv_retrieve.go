package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	user "github.com/wheelergeo/g-otter-gen/user"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/model"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/mysql"
	"github.com/wheelergeo/g-otter-pkg/constants"
)

type RpcUserCommonKvRetrieveService struct {
	ctx context.Context
} // NewRpcUserCommonKvRetrieveService new RpcUserCommonKvRetrieveService
func NewRpcUserCommonKvRetrieveService(ctx context.Context) *RpcUserCommonKvRetrieveService {
	return &RpcUserCommonKvRetrieveService{ctx: ctx}
}

// Run create note info
func (s *RpcUserCommonKvRetrieveService) Run(req *user.RpcUserCommonKvRetrieveReq) (resp *user.RpcUserCommonKvRetrieveResp, err error) {
	err = req.IsValid()
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.ParamInvalid
		resp.Status.Message = constants.ParamInvalidMsg
		return
	}
	resp = new(user.RpcUserCommonKvRetrieveResp)
	resp.Status = new(user.BaseResp)

	kvModels := []model.UserCommonKv{}
	db := mysql.DB
	if req.Key != nil {
		db = db.Where("key = ?", *req.Key)
	}

	err = mysql.DB.Find(&kvModels).Error
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.UserCommonKvRetrieveFailed
		resp.Status.Message = constants.UserCommonKvRetrieveFailedMsg
		return
	}

	resp.Value = make(map[string]string, len(kvModels))
	for _, v := range kvModels {
		resp.Value[v.Key] = v.Value
	}

	resp.Status.Code = constants.UserCommonKvRetrieveSuccess
	resp.Status.Message = constants.UserCommonKvRetrieveSuccessMsg
	return
}
