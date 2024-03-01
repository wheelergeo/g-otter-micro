package service

import (
	"context"
	"strings"

	"github.com/cloudwego/kitex/pkg/klog"
	user "github.com/wheelergeo/g-otter-gen/user"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/model"
	"github.com/wheelergeo/g-otter-micro-user/biz/dal/mysql"
	"github.com/wheelergeo/g-otter-pkg/constants"
)

type RpcUserCommonPostRetrieveService struct {
	ctx context.Context
} // NewRpcUserCommonPostRetrieveService new RpcUserCommonPostRetrieveService
func NewRpcUserCommonPostRetrieveService(ctx context.Context) *RpcUserCommonPostRetrieveService {
	return &RpcUserCommonPostRetrieveService{ctx: ctx}
}

// Run create note info
func (s *RpcUserCommonPostRetrieveService) Run(req *user.RpcUserCommonPostRetrieveReq) (resp *user.RpcUserCommonPostRetrieveResp, err error) {
	err = req.IsValid()
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.ParamInvalid
		resp.Status.Message = constants.ParamInvalidMsg
		return
	}
	resp = new(user.RpcUserCommonPostRetrieveResp)
	resp.Status = new(user.BaseResp)

	postModels := []model.UserCommonPost{}
	roleModels := []model.UserCommonRole{}
	db := mysql.DB
	if req.PostCode != nil {
		db = db.Where("post_code LIKE ?", "%"+*req.PostCode+"%")
	}
	if req.PostName != nil {
		db = db.Where("post_name LIKE ?", "%"+*req.PostName+"%")
	}
	if req.Level != nil {
		db = db.Where("level = ?", *req.Level)
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}
	if req.Id != nil {
		db = db.Where("id = ?", *req.Id)
	}

	err = db.Find(&postModels).Error
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.UserCommonPostRetrieveFailed
		resp.Status.Message = constants.UserCommonPostRetrieveFailedMsg
		return
	}

	for _, v := range postModels {
		roleIds := strings.Split(v.RoleIds, ",")
		err = mysql.DB.Where("id IN ?", roleIds[:len(roleIds)-1]).Find(&roleModels).Error
		if err != nil {
			klog.CtxErrorf(s.ctx, err.Error())
			resp.Status.Code = constants.UserCommonPostRetrieveFailed
			resp.Status.Message = constants.UserCommonPostRetrieveFailedMsg
			return
		}

		roles := make(map[int32]string, len(roleModels))
		for _, v := range roleModels {
			roles[v.ID] = v.Name
		}

		resp.List = append(resp.List, &user.UserPostInfo{
			Id:       int32(v.ID),
			PostCode: v.PostCode,
			PostName: v.PostName,
			Level:    int8(v.Level),
			Status:   int8(v.Status),
			Remark:   v.Remark,
			Roles:    roles,
			CreateAt: v.CreatedAt.String(),
			UpdateAt: v.UpdatedAt.String(),
		})
	}

	resp.Status.Code = constants.UserCommonPostRetrieveSuccess
	resp.Status.Message = constants.UserCommonPostRetrieveSuccessMsg
	return
}
