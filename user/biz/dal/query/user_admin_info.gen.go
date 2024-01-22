// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/wheelergeo/g-otter-micro-user/biz/dal/model"
)

func newUserAdminInfo(db *gorm.DB, opts ...gen.DOOption) userAdminInfo {
	_userAdminInfo := userAdminInfo{}

	_userAdminInfo.userAdminInfoDo.UseDB(db, opts...)
	_userAdminInfo.userAdminInfoDo.UseModel(&model.UserAdminInfo{})

	tableName := _userAdminInfo.userAdminInfoDo.TableName()
	_userAdminInfo.ALL = field.NewAsterisk(tableName)
	_userAdminInfo.ID = field.NewInt64(tableName, "id")
	_userAdminInfo.UUID = field.NewInt64(tableName, "uuid")
	_userAdminInfo.UserName = field.NewString(tableName, "user_name")
	_userAdminInfo.PhoneNumber = field.NewString(tableName, "phone_number")
	_userAdminInfo.RealName = field.NewString(tableName, "real_name")
	_userAdminInfo.UserPassword = field.NewString(tableName, "user_password")
	_userAdminInfo.UserStatus = field.NewInt32(tableName, "user_status")
	_userAdminInfo.UserEmail = field.NewString(tableName, "user_email")
	_userAdminInfo.UserMode = field.NewInt32(tableName, "user_mode")
	_userAdminInfo.UserAddress = field.NewString(tableName, "user_address")
	_userAdminInfo.IdentityNum = field.NewString(tableName, "identity_num")
	_userAdminInfo.Sex = field.NewInt32(tableName, "sex")
	_userAdminInfo.Avatar = field.NewString(tableName, "avatar")
	_userAdminInfo.DeptID = field.NewInt32(tableName, "dept_id")
	_userAdminInfo.PostID = field.NewInt32(tableName, "post_id")
	_userAdminInfo.Remark = field.NewString(tableName, "remark")
	_userAdminInfo.Describe = field.NewString(tableName, "describe")
	_userAdminInfo.WxOpenID = field.NewString(tableName, "wx_open_id")
	_userAdminInfo.DyOpenID = field.NewString(tableName, "dy_open_id")
	_userAdminInfo.LastLoginIP = field.NewString(tableName, "last_login_ip")
	_userAdminInfo.LastLoginTime = field.NewTime(tableName, "last_login_time")
	_userAdminInfo.CreatedAt = field.NewTime(tableName, "created_at")
	_userAdminInfo.UpdatedAt = field.NewTime(tableName, "updated_at")
	_userAdminInfo.DeletedAt = field.NewField(tableName, "deleted_at")

	_userAdminInfo.fillFieldMap()

	return _userAdminInfo
}

type userAdminInfo struct {
	userAdminInfoDo userAdminInfoDo

	ALL           field.Asterisk
	ID            field.Int64
	UUID          field.Int64  // 人员唯一标识符
	UserName      field.String // 用户名
	PhoneNumber   field.String // +国家代码 手机号
	RealName      field.String // 真实姓名
	UserPassword  field.String // 登录密码;cmf_password加密
	UserStatus    field.Int32  // 用户状态;0:禁用,1:正常,2:未验证
	UserEmail     field.String // 用户邮箱
	UserMode      field.Int32  // 1:内部用户 2:外部用户
	UserAddress   field.String // 联系地址
	IdentityNum   field.String // 身份证号
	Sex           field.Int32  // 性别;0:保密,1:男,2:女
	Avatar        field.String // 用户头像
	DeptID        field.Int32  // 部门id
	PostID        field.Int32  // 岗位id
	Remark        field.String // 备注
	Describe      field.String // 描述信息
	WxOpenID      field.String // 微信openid
	DyOpenID      field.String // 抖音openid
	LastLoginIP   field.String // 最后登录ip
	LastLoginTime field.Time   // 最后登录时间
	CreatedAt     field.Time   // 创建时间
	UpdatedAt     field.Time   // 更新时间
	DeletedAt     field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (u userAdminInfo) Table(newTableName string) *userAdminInfo {
	u.userAdminInfoDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userAdminInfo) As(alias string) *userAdminInfo {
	u.userAdminInfoDo.DO = *(u.userAdminInfoDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userAdminInfo) updateTableName(table string) *userAdminInfo {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.UUID = field.NewInt64(table, "uuid")
	u.UserName = field.NewString(table, "user_name")
	u.PhoneNumber = field.NewString(table, "phone_number")
	u.RealName = field.NewString(table, "real_name")
	u.UserPassword = field.NewString(table, "user_password")
	u.UserStatus = field.NewInt32(table, "user_status")
	u.UserEmail = field.NewString(table, "user_email")
	u.UserMode = field.NewInt32(table, "user_mode")
	u.UserAddress = field.NewString(table, "user_address")
	u.IdentityNum = field.NewString(table, "identity_num")
	u.Sex = field.NewInt32(table, "sex")
	u.Avatar = field.NewString(table, "avatar")
	u.DeptID = field.NewInt32(table, "dept_id")
	u.PostID = field.NewInt32(table, "post_id")
	u.Remark = field.NewString(table, "remark")
	u.Describe = field.NewString(table, "describe")
	u.WxOpenID = field.NewString(table, "wx_open_id")
	u.DyOpenID = field.NewString(table, "dy_open_id")
	u.LastLoginIP = field.NewString(table, "last_login_ip")
	u.LastLoginTime = field.NewTime(table, "last_login_time")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")

	u.fillFieldMap()

	return u
}

func (u *userAdminInfo) WithContext(ctx context.Context) *userAdminInfoDo {
	return u.userAdminInfoDo.WithContext(ctx)
}

func (u userAdminInfo) TableName() string { return u.userAdminInfoDo.TableName() }

func (u userAdminInfo) Alias() string { return u.userAdminInfoDo.Alias() }

func (u *userAdminInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userAdminInfo) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 24)
	u.fieldMap["id"] = u.ID
	u.fieldMap["uuid"] = u.UUID
	u.fieldMap["user_name"] = u.UserName
	u.fieldMap["phone_number"] = u.PhoneNumber
	u.fieldMap["real_name"] = u.RealName
	u.fieldMap["user_password"] = u.UserPassword
	u.fieldMap["user_status"] = u.UserStatus
	u.fieldMap["user_email"] = u.UserEmail
	u.fieldMap["user_mode"] = u.UserMode
	u.fieldMap["user_address"] = u.UserAddress
	u.fieldMap["identity_num"] = u.IdentityNum
	u.fieldMap["sex"] = u.Sex
	u.fieldMap["avatar"] = u.Avatar
	u.fieldMap["dept_id"] = u.DeptID
	u.fieldMap["post_id"] = u.PostID
	u.fieldMap["remark"] = u.Remark
	u.fieldMap["describe"] = u.Describe
	u.fieldMap["wx_open_id"] = u.WxOpenID
	u.fieldMap["dy_open_id"] = u.DyOpenID
	u.fieldMap["last_login_ip"] = u.LastLoginIP
	u.fieldMap["last_login_time"] = u.LastLoginTime
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
}

func (u userAdminInfo) clone(db *gorm.DB) userAdminInfo {
	u.userAdminInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userAdminInfo) replaceDB(db *gorm.DB) userAdminInfo {
	u.userAdminInfoDo.ReplaceDB(db)
	return u
}

type userAdminInfoDo struct{ gen.DO }

func (u userAdminInfoDo) Debug() *userAdminInfoDo {
	return u.withDO(u.DO.Debug())
}

func (u userAdminInfoDo) WithContext(ctx context.Context) *userAdminInfoDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userAdminInfoDo) ReadDB() *userAdminInfoDo {
	return u.Clauses(dbresolver.Read)
}

func (u userAdminInfoDo) WriteDB() *userAdminInfoDo {
	return u.Clauses(dbresolver.Write)
}

func (u userAdminInfoDo) Session(config *gorm.Session) *userAdminInfoDo {
	return u.withDO(u.DO.Session(config))
}

func (u userAdminInfoDo) Clauses(conds ...clause.Expression) *userAdminInfoDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userAdminInfoDo) Returning(value interface{}, columns ...string) *userAdminInfoDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userAdminInfoDo) Not(conds ...gen.Condition) *userAdminInfoDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userAdminInfoDo) Or(conds ...gen.Condition) *userAdminInfoDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userAdminInfoDo) Select(conds ...field.Expr) *userAdminInfoDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userAdminInfoDo) Where(conds ...gen.Condition) *userAdminInfoDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userAdminInfoDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *userAdminInfoDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u userAdminInfoDo) Order(conds ...field.Expr) *userAdminInfoDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userAdminInfoDo) Distinct(cols ...field.Expr) *userAdminInfoDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userAdminInfoDo) Omit(cols ...field.Expr) *userAdminInfoDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userAdminInfoDo) Join(table schema.Tabler, on ...field.Expr) *userAdminInfoDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userAdminInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userAdminInfoDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userAdminInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) *userAdminInfoDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userAdminInfoDo) Group(cols ...field.Expr) *userAdminInfoDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userAdminInfoDo) Having(conds ...gen.Condition) *userAdminInfoDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userAdminInfoDo) Limit(limit int) *userAdminInfoDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userAdminInfoDo) Offset(offset int) *userAdminInfoDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userAdminInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userAdminInfoDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userAdminInfoDo) Unscoped() *userAdminInfoDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userAdminInfoDo) Create(values ...*model.UserAdminInfo) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userAdminInfoDo) CreateInBatches(values []*model.UserAdminInfo, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userAdminInfoDo) Save(values ...*model.UserAdminInfo) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userAdminInfoDo) First() (*model.UserAdminInfo, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAdminInfo), nil
	}
}

func (u userAdminInfoDo) Take() (*model.UserAdminInfo, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAdminInfo), nil
	}
}

func (u userAdminInfoDo) Last() (*model.UserAdminInfo, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAdminInfo), nil
	}
}

func (u userAdminInfoDo) Find() ([]*model.UserAdminInfo, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserAdminInfo), err
}

func (u userAdminInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserAdminInfo, err error) {
	buf := make([]*model.UserAdminInfo, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userAdminInfoDo) FindInBatches(result *[]*model.UserAdminInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userAdminInfoDo) Attrs(attrs ...field.AssignExpr) *userAdminInfoDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userAdminInfoDo) Assign(attrs ...field.AssignExpr) *userAdminInfoDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userAdminInfoDo) Joins(fields ...field.RelationField) *userAdminInfoDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userAdminInfoDo) Preload(fields ...field.RelationField) *userAdminInfoDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userAdminInfoDo) FirstOrInit() (*model.UserAdminInfo, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAdminInfo), nil
	}
}

func (u userAdminInfoDo) FirstOrCreate() (*model.UserAdminInfo, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAdminInfo), nil
	}
}

func (u userAdminInfoDo) FindByPage(offset int, limit int) (result []*model.UserAdminInfo, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userAdminInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userAdminInfoDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userAdminInfoDo) Delete(models ...*model.UserAdminInfo) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userAdminInfoDo) withDO(do gen.Dao) *userAdminInfoDo {
	u.DO = *do.(*gen.DO)
	return u
}
