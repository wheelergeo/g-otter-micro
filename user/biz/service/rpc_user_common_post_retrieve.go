package service

import (
	"context"
	"strconv"
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

	err = mysql.DB.Find(&postModels).Error
	if err != nil {
		klog.CtxErrorf(s.ctx, err.Error())
		resp.Status.Code = constants.UserCommonPostRetrieveFailed
		resp.Status.Message = constants.UserCommonPostRetrieveFailedMsg
		return
	}

	for _, v := range postModels {
		roleStrIds := strings.Split(v.RoleIds, ",")
		roleIds := make([]int32, 0, len(roleStrIds)-1)
		for _, v := range roleStrIds {
			if v == " " {
				continue
			}
			num, _ := strconv.Atoi(v)
			roleIds = append(roleIds, int32(num))
		}

		resp.List = append(resp.List, &user.UserPostInfo{
			Id:       int32(v.ID),
			PostCode: v.PostCode,
			PostName: v.PostName,
			Level:    int8(v.Level),
			Status:   int8(v.Status),
			Remark:   v.Remark,
			RoleIds:  roleIds,
			CreateAt: v.CreatedAt.String(),
			UpdateAt: v.UpdatedAt.String(),
		})
	}

	resp.Status.Code = constants.UserCommonPostRetrieveSuccess
	resp.Status.Message = constants.UserCommonPostRetrieveSuccessMsg
	return
}
