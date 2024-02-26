package service

import (
	"context"
	"testing"

	user "github.com/wheelergeo/g-otter-gen/user"
	mysql2 "github.com/wheelergeo/g-otter-micro-user/biz/dal/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestRpcUserCommonDeptCreate_Run(t *testing.T) {
	mysql2.DB, _ = gorm.Open(
		mysql.Open("root:123456@tcp(127.0.0.1:3306)/g_otter?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	ctx := context.Background()
	s := NewRpcUserCommonDeptCreateService(ctx)
	// init req and assert value

	req := &user.RpcUserCommonDeptCreateReq{
		ParentId:   0,
		AncestorId: 0,
		DeptName:   "测试",
		Status:     1,
		Leader:     nil,
		PhoneNum:   nil,
		Email:      nil,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
