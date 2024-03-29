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

func newUserCommonPost(db *gorm.DB, opts ...gen.DOOption) userCommonPost {
	_userCommonPost := userCommonPost{}

	_userCommonPost.userCommonPostDo.UseDB(db, opts...)
	_userCommonPost.userCommonPostDo.UseModel(&model.UserCommonPost{})

	tableName := _userCommonPost.userCommonPostDo.TableName()
	_userCommonPost.ALL = field.NewAsterisk(tableName)
	_userCommonPost.ID = field.NewInt64(tableName, "id")
	_userCommonPost.PostCode = field.NewString(tableName, "post_code")
	_userCommonPost.PostName = field.NewString(tableName, "post_name")
	_userCommonPost.Level = field.NewInt32(tableName, "level")
	_userCommonPost.Status = field.NewInt32(tableName, "status")
	_userCommonPost.Remark = field.NewString(tableName, "remark")
	_userCommonPost.RoleIds = field.NewString(tableName, "role_ids")
	_userCommonPost.CreatedAt = field.NewTime(tableName, "created_at")
	_userCommonPost.UpdatedAt = field.NewTime(tableName, "updated_at")

	_userCommonPost.fillFieldMap()

	return _userCommonPost
}

type userCommonPost struct {
	userCommonPostDo userCommonPostDo

	ALL       field.Asterisk
	ID        field.Int64  // 岗位ID
	PostCode  field.String // 岗位编码
	PostName  field.String // 岗位名称
	Level     field.Int32  // 层级
	Status    field.Int32  // 状态（0正常 1停用）
	Remark    field.String // 备注
	RoleIds   field.String // 关联角色数据
	CreatedAt field.Time   // 创建时间
	UpdatedAt field.Time   // 修改时间

	fieldMap map[string]field.Expr
}

func (u userCommonPost) Table(newTableName string) *userCommonPost {
	u.userCommonPostDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userCommonPost) As(alias string) *userCommonPost {
	u.userCommonPostDo.DO = *(u.userCommonPostDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userCommonPost) updateTableName(table string) *userCommonPost {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.PostCode = field.NewString(table, "post_code")
	u.PostName = field.NewString(table, "post_name")
	u.Level = field.NewInt32(table, "level")
	u.Status = field.NewInt32(table, "status")
	u.Remark = field.NewString(table, "remark")
	u.RoleIds = field.NewString(table, "role_ids")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")

	u.fillFieldMap()

	return u
}

func (u *userCommonPost) WithContext(ctx context.Context) *userCommonPostDo {
	return u.userCommonPostDo.WithContext(ctx)
}

func (u userCommonPost) TableName() string { return u.userCommonPostDo.TableName() }

func (u userCommonPost) Alias() string { return u.userCommonPostDo.Alias() }

func (u *userCommonPost) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userCommonPost) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 9)
	u.fieldMap["id"] = u.ID
	u.fieldMap["post_code"] = u.PostCode
	u.fieldMap["post_name"] = u.PostName
	u.fieldMap["level"] = u.Level
	u.fieldMap["status"] = u.Status
	u.fieldMap["remark"] = u.Remark
	u.fieldMap["role_ids"] = u.RoleIds
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
}

func (u userCommonPost) clone(db *gorm.DB) userCommonPost {
	u.userCommonPostDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userCommonPost) replaceDB(db *gorm.DB) userCommonPost {
	u.userCommonPostDo.ReplaceDB(db)
	return u
}

type userCommonPostDo struct{ gen.DO }

func (u userCommonPostDo) Debug() *userCommonPostDo {
	return u.withDO(u.DO.Debug())
}

func (u userCommonPostDo) WithContext(ctx context.Context) *userCommonPostDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userCommonPostDo) ReadDB() *userCommonPostDo {
	return u.Clauses(dbresolver.Read)
}

func (u userCommonPostDo) WriteDB() *userCommonPostDo {
	return u.Clauses(dbresolver.Write)
}

func (u userCommonPostDo) Session(config *gorm.Session) *userCommonPostDo {
	return u.withDO(u.DO.Session(config))
}

func (u userCommonPostDo) Clauses(conds ...clause.Expression) *userCommonPostDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userCommonPostDo) Returning(value interface{}, columns ...string) *userCommonPostDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userCommonPostDo) Not(conds ...gen.Condition) *userCommonPostDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userCommonPostDo) Or(conds ...gen.Condition) *userCommonPostDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userCommonPostDo) Select(conds ...field.Expr) *userCommonPostDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userCommonPostDo) Where(conds ...gen.Condition) *userCommonPostDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userCommonPostDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *userCommonPostDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u userCommonPostDo) Order(conds ...field.Expr) *userCommonPostDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userCommonPostDo) Distinct(cols ...field.Expr) *userCommonPostDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userCommonPostDo) Omit(cols ...field.Expr) *userCommonPostDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userCommonPostDo) Join(table schema.Tabler, on ...field.Expr) *userCommonPostDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userCommonPostDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userCommonPostDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userCommonPostDo) RightJoin(table schema.Tabler, on ...field.Expr) *userCommonPostDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userCommonPostDo) Group(cols ...field.Expr) *userCommonPostDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userCommonPostDo) Having(conds ...gen.Condition) *userCommonPostDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userCommonPostDo) Limit(limit int) *userCommonPostDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userCommonPostDo) Offset(offset int) *userCommonPostDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userCommonPostDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userCommonPostDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userCommonPostDo) Unscoped() *userCommonPostDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userCommonPostDo) Create(values ...*model.UserCommonPost) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userCommonPostDo) CreateInBatches(values []*model.UserCommonPost, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userCommonPostDo) Save(values ...*model.UserCommonPost) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userCommonPostDo) First() (*model.UserCommonPost, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCommonPost), nil
	}
}

func (u userCommonPostDo) Take() (*model.UserCommonPost, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCommonPost), nil
	}
}

func (u userCommonPostDo) Last() (*model.UserCommonPost, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCommonPost), nil
	}
}

func (u userCommonPostDo) Find() ([]*model.UserCommonPost, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserCommonPost), err
}

func (u userCommonPostDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserCommonPost, err error) {
	buf := make([]*model.UserCommonPost, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userCommonPostDo) FindInBatches(result *[]*model.UserCommonPost, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userCommonPostDo) Attrs(attrs ...field.AssignExpr) *userCommonPostDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userCommonPostDo) Assign(attrs ...field.AssignExpr) *userCommonPostDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userCommonPostDo) Joins(fields ...field.RelationField) *userCommonPostDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userCommonPostDo) Preload(fields ...field.RelationField) *userCommonPostDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userCommonPostDo) FirstOrInit() (*model.UserCommonPost, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCommonPost), nil
	}
}

func (u userCommonPostDo) FirstOrCreate() (*model.UserCommonPost, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCommonPost), nil
	}
}

func (u userCommonPostDo) FindByPage(offset int, limit int) (result []*model.UserCommonPost, count int64, err error) {
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

func (u userCommonPostDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userCommonPostDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userCommonPostDo) Delete(models ...*model.UserCommonPost) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userCommonPostDo) withDO(do gen.Dao) *userCommonPostDo {
	u.DO = *do.(*gen.DO)
	return u
}
