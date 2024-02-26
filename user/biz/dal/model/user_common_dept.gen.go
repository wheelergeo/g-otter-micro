// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUserCommonDept = "user_common_dept"

// UserCommonDept mapped from table <user_common_dept>
type UserCommonDept struct {
	ID          int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 部门id
	ParentID    int64          `gorm:"column:parent_id;not null" json:"parent_id"`        // 父级id
	AncestorID  int64          `gorm:"column:ancestor_id;not null" json:"ancestor_id"`    // 祖先id
	DeptName    string         `gorm:"column:dept_name;not null" json:"dept_name"`        // 部门名称
	OrderNum    int32          `gorm:"column:order_num;not null" json:"order_num"`        // 显示顺序
	Leader      string         `gorm:"column:leader;not null" json:"leader"`              // 负责人
	PhoneNumber string         `gorm:"column:phone_number;not null" json:"phone_number"`  // 联系电话
	Email       string         `gorm:"column:email;not null" json:"email"`                // 邮箱
	Status      int32          `gorm:"column:status;not null" json:"status"`              // 部门状态（0正常 1停用）
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"` // 删除时间
}

// TableName UserCommonDept's table name
func (*UserCommonDept) TableName() string {
	return TableNameUserCommonDept
}