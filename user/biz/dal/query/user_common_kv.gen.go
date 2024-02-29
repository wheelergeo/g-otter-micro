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

func newUserCommonKv(db *gorm.DB, opts ...gen.DOOption) userCommonKv {
	_userCommonKv := userCommonKv{}

	_userCommonKv.userCommonKvDo.UseDB(db, opts...)
	_userCommonKv.userCommonKvDo.UseModel(&model.UserCommonKv{})

	tableName := _userCommonKv.userCommonKvDo.TableName()
	_userCommonKv.ALL = field.NewAsterisk(tableName)
	_userCommonKv.Key = field.NewString(tableName, "key")
	_userCommonKv.Value = field.NewString(tableName, "value")

	_userCommonKv.fillFieldMap()

	return _userCommonKv
}

type userCommonKv struct {
	userCommonKvDo userCommonKvDo

	ALL   field.Asterisk
	Key   field.String // key值
	Value field.String // value值

	fieldMap map[string]field.Expr
}

func (u userCommonKv) Table(newTableName string) *userCommonKv {
	u.userCommonKvDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userCommonKv) As(alias string) *userCommonKv {
	u.userCommonKvDo.DO = *(u.userCommonKvDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userCommonKv) updateTableName(table string) *userCommonKv {
	u.ALL = field.NewAsterisk(table)
	u.Key = field.NewString(table, "key")
	u.Value = field.NewString(table, "value")

	u.fillFieldMap()

	return u
}

func (u *userCommonKv) WithContext(ctx context.Context) *userCommonKvDo {
	return u.userCommonKvDo.WithContext(ctx)
}

func (u userCommonKv) TableName() string { return u.userCommonKvDo.TableName() }

func (u userCommonKv) Alias() string { return u.userCommonKvDo.Alias() }

func (u *userCommonKv) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userCommonKv) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 2)
	u.fieldMap["key"] = u.Key
	u.fieldMap["value"] = u.Value
}

func (u userCommonKv) clone(db *gorm.DB) userCommonKv {
	u.userCommonKvDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userCommonKv) replaceDB(db *gorm.DB) userCommonKv {
	u.userCommonKvDo.ReplaceDB(db)
	return u
}

type userCommonKvDo struct{ gen.DO }

func (u userCommonKvDo) Debug() *userCommonKvDo {
	return u.withDO(u.DO.Debug())
}

func (u userCommonKvDo) WithContext(ctx context.Context) *userCommonKvDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userCommonKvDo) ReadDB() *userCommonKvDo {
	return u.Clauses(dbresolver.Read)
}

func (u userCommonKvDo) WriteDB() *userCommonKvDo {
	return u.Clauses(dbresolver.Write)
}

func (u userCommonKvDo) Session(config *gorm.Session) *userCommonKvDo {
	return u.withDO(u.DO.Session(config))
}

func (u userCommonKvDo) Clauses(conds ...clause.Expression) *userCommonKvDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userCommonKvDo) Returning(value interface{}, columns ...string) *userCommonKvDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userCommonKvDo) Not(conds ...gen.Condition) *userCommonKvDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userCommonKvDo) Or(conds ...gen.Condition) *userCommonKvDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userCommonKvDo) Select(conds ...field.Expr) *userCommonKvDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userCommonKvDo) Where(conds ...gen.Condition) *userCommonKvDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userCommonKvDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *userCommonKvDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u userCommonKvDo) Order(conds ...field.Expr) *userCommonKvDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userCommonKvDo) Distinct(cols ...field.Expr) *userCommonKvDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userCommonKvDo) Omit(cols ...field.Expr) *userCommonKvDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userCommonKvDo) Join(table schema.Tabler, on ...field.Expr) *userCommonKvDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userCommonKvDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userCommonKvDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userCommonKvDo) RightJoin(table schema.Tabler, on ...field.Expr) *userCommonKvDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userCommonKvDo) Group(cols ...field.Expr) *userCommonKvDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userCommonKvDo) Having(conds ...gen.Condition) *userCommonKvDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userCommonKvDo) Limit(limit int) *userCommonKvDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userCommonKvDo) Offset(offset int) *userCommonKvDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userCommonKvDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userCommonKvDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userCommonKvDo) Unscoped() *userCommonKvDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userCommonKvDo) Create(values ...*model.UserCommonKv) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userCommonKvDo) CreateInBatches(values []*model.UserCommonKv, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userCommonKvDo) Save(values ...*model.UserCommonKv) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userCommonKvDo) First() (*model.UserCommonKv, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCommonKv), nil
	}
}

func (u userCommonKvDo) Take() (*model.UserCommonKv, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCommonKv), nil
	}
}

func (u userCommonKvDo) Last() (*model.UserCommonKv, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCommonKv), nil
	}
}

func (u userCommonKvDo) Find() ([]*model.UserCommonKv, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserCommonKv), err
}

func (u userCommonKvDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserCommonKv, err error) {
	buf := make([]*model.UserCommonKv, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userCommonKvDo) FindInBatches(result *[]*model.UserCommonKv, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userCommonKvDo) Attrs(attrs ...field.AssignExpr) *userCommonKvDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userCommonKvDo) Assign(attrs ...field.AssignExpr) *userCommonKvDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userCommonKvDo) Joins(fields ...field.RelationField) *userCommonKvDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userCommonKvDo) Preload(fields ...field.RelationField) *userCommonKvDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userCommonKvDo) FirstOrInit() (*model.UserCommonKv, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCommonKv), nil
	}
}

func (u userCommonKvDo) FirstOrCreate() (*model.UserCommonKv, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCommonKv), nil
	}
}

func (u userCommonKvDo) FindByPage(offset int, limit int) (result []*model.UserCommonKv, count int64, err error) {
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

func (u userCommonKvDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userCommonKvDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userCommonKvDo) Delete(models ...*model.UserCommonKv) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userCommonKvDo) withDO(do gen.Dao) *userCommonKvDo {
	u.DO = *do.(*gen.DO)
	return u
}