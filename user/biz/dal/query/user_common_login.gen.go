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

func newUserCommonLogin(db *gorm.DB, opts ...gen.DOOption) userCommonLogin {
	_userCommonLogin := userCommonLogin{}

	_userCommonLogin.userCommonLoginDo.UseDB(db, opts...)
	_userCommonLogin.userCommonLoginDo.UseModel(&model.UserCommonLogin{})

	tableName := _userCommonLogin.userCommonLoginDo.TableName()
	_userCommonLogin.ALL = field.NewAsterisk(tableName)
	_userCommonLogin.ID = field.NewInt64(tableName, "id")
	_userCommonLogin.UUID = field.NewInt64(tableName, "uuid")
	_userCommonLogin.UserName = field.NewString(tableName, "user_name")
	_userCommonLogin.PhoneNumber = field.NewString(tableName, "phone_number")
	_userCommonLogin.IP = field.NewString(tableName, "ip")
	_userCommonLogin.Location = field.NewString(tableName, "location")
	_userCommonLogin.Browser = field.NewString(tableName, "browser")
	_userCommonLogin.Os = field.NewString(tableName, "os")
	_userCommonLogin.Status = field.NewInt32(tableName, "status")
	_userCommonLogin.Msg = field.NewString(tableName, "msg")
	_userCommonLogin.LoginAt = field.NewTime(tableName, "login_at")
	_userCommonLogin.Client = field.NewString(tableName, "client")

	_userCommonLogin.fillFieldMap()

	return _userCommonLogin
}

type userCommonLogin struct {
	userCommonLoginDo userCommonLoginDo

	ALL         field.Asterisk
	ID          field.Int64
	UUID        field.Int64  // 用户标示
	UserName    field.String // 登录用户名
	PhoneNumber field.String // 登录手机号
	IP          field.String // 登录IP地址
	Location    field.String // 登录地点
	Browser     field.String // 浏览器类型
	Os          field.String // 操作系统
	Status      field.Int32  // 登录状态（0成功 1失败）
	Msg         field.String // 提示消息
	LoginAt     field.Time   // 登录时间
	Client      field.String // 登录客户端：web，wx，dy，app

	fieldMap map[string]field.Expr
}

func (u userCommonLogin) Table(newTableName string) *userCommonLogin {
	u.userCommonLoginDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userCommonLogin) As(alias string) *userCommonLogin {
	u.userCommonLoginDo.DO = *(u.userCommonLoginDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userCommonLogin) updateTableName(table string) *userCommonLogin {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.UUID = field.NewInt64(table, "uuid")
	u.UserName = field.NewString(table, "user_name")
	u.PhoneNumber = field.NewString(table, "phone_number")
	u.IP = field.NewString(table, "ip")
	u.Location = field.NewString(table, "location")
	u.Browser = field.NewString(table, "browser")
	u.Os = field.NewString(table, "os")
	u.Status = field.NewInt32(table, "status")
	u.Msg = field.NewString(table, "msg")
	u.LoginAt = field.NewTime(table, "login_at")
	u.Client = field.NewString(table, "client")

	u.fillFieldMap()

	return u
}

func (u *userCommonLogin) WithContext(ctx context.Context) *userCommonLoginDo {
	return u.userCommonLoginDo.WithContext(ctx)
}

func (u userCommonLogin) TableName() string { return u.userCommonLoginDo.TableName() }

func (u userCommonLogin) Alias() string { return u.userCommonLoginDo.Alias() }

func (u *userCommonLogin) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userCommonLogin) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 12)
	u.fieldMap["id"] = u.ID
	u.fieldMap["uuid"] = u.UUID
	u.fieldMap["user_name"] = u.UserName
	u.fieldMap["phone_number"] = u.PhoneNumber
	u.fieldMap["ip"] = u.IP
	u.fieldMap["location"] = u.Location
	u.fieldMap["browser"] = u.Browser
	u.fieldMap["os"] = u.Os
	u.fieldMap["status"] = u.Status
	u.fieldMap["msg"] = u.Msg
	u.fieldMap["login_at"] = u.LoginAt
	u.fieldMap["client"] = u.Client
}

func (u userCommonLogin) clone(db *gorm.DB) userCommonLogin {
	u.userCommonLoginDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userCommonLogin) replaceDB(db *gorm.DB) userCommonLogin {
	u.userCommonLoginDo.ReplaceDB(db)
	return u
}

type userCommonLoginDo struct{ gen.DO }

func (u userCommonLoginDo) Debug() *userCommonLoginDo {
	return u.withDO(u.DO.Debug())
}

func (u userCommonLoginDo) WithContext(ctx context.Context) *userCommonLoginDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userCommonLoginDo) ReadDB() *userCommonLoginDo {
	return u.Clauses(dbresolver.Read)
}

func (u userCommonLoginDo) WriteDB() *userCommonLoginDo {
	return u.Clauses(dbresolver.Write)
}

func (u userCommonLoginDo) Session(config *gorm.Session) *userCommonLoginDo {
	return u.withDO(u.DO.Session(config))
}

func (u userCommonLoginDo) Clauses(conds ...clause.Expression) *userCommonLoginDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userCommonLoginDo) Returning(value interface{}, columns ...string) *userCommonLoginDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userCommonLoginDo) Not(conds ...gen.Condition) *userCommonLoginDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userCommonLoginDo) Or(conds ...gen.Condition) *userCommonLoginDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userCommonLoginDo) Select(conds ...field.Expr) *userCommonLoginDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userCommonLoginDo) Where(conds ...gen.Condition) *userCommonLoginDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userCommonLoginDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *userCommonLoginDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u userCommonLoginDo) Order(conds ...field.Expr) *userCommonLoginDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userCommonLoginDo) Distinct(cols ...field.Expr) *userCommonLoginDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userCommonLoginDo) Omit(cols ...field.Expr) *userCommonLoginDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userCommonLoginDo) Join(table schema.Tabler, on ...field.Expr) *userCommonLoginDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userCommonLoginDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userCommonLoginDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userCommonLoginDo) RightJoin(table schema.Tabler, on ...field.Expr) *userCommonLoginDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userCommonLoginDo) Group(cols ...field.Expr) *userCommonLoginDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userCommonLoginDo) Having(conds ...gen.Condition) *userCommonLoginDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userCommonLoginDo) Limit(limit int) *userCommonLoginDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userCommonLoginDo) Offset(offset int) *userCommonLoginDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userCommonLoginDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userCommonLoginDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userCommonLoginDo) Unscoped() *userCommonLoginDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userCommonLoginDo) Create(values ...*model.UserCommonLogin) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userCommonLoginDo) CreateInBatches(values []*model.UserCommonLogin, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userCommonLoginDo) Save(values ...*model.UserCommonLogin) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userCommonLoginDo) First() (*model.UserCommonLogin, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCommonLogin), nil
	}
}

func (u userCommonLoginDo) Take() (*model.UserCommonLogin, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCommonLogin), nil
	}
}

func (u userCommonLoginDo) Last() (*model.UserCommonLogin, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCommonLogin), nil
	}
}

func (u userCommonLoginDo) Find() ([]*model.UserCommonLogin, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserCommonLogin), err
}

func (u userCommonLoginDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserCommonLogin, err error) {
	buf := make([]*model.UserCommonLogin, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userCommonLoginDo) FindInBatches(result *[]*model.UserCommonLogin, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userCommonLoginDo) Attrs(attrs ...field.AssignExpr) *userCommonLoginDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userCommonLoginDo) Assign(attrs ...field.AssignExpr) *userCommonLoginDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userCommonLoginDo) Joins(fields ...field.RelationField) *userCommonLoginDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userCommonLoginDo) Preload(fields ...field.RelationField) *userCommonLoginDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userCommonLoginDo) FirstOrInit() (*model.UserCommonLogin, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCommonLogin), nil
	}
}

func (u userCommonLoginDo) FirstOrCreate() (*model.UserCommonLogin, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCommonLogin), nil
	}
}

func (u userCommonLoginDo) FindByPage(offset int, limit int) (result []*model.UserCommonLogin, count int64, err error) {
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

func (u userCommonLoginDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userCommonLoginDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userCommonLoginDo) Delete(models ...*model.UserCommonLogin) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userCommonLoginDo) withDO(do gen.Dao) *userCommonLoginDo {
	u.DO = *do.(*gen.DO)
	return u
}