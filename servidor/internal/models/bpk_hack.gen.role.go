package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RoleMgr struct {
	*_BaseMgr
}

// RoleMgr open func
func RoleMgr(db *gorm.DB) *_RoleMgr {
	if db == nil {
		panic(fmt.Errorf("RoleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RoleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("role"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_RoleMgr) Debug() *_RoleMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RoleMgr) GetTableName() string {
	return "role"
}

// Reset 重置gorm会话
func (obj *_RoleMgr) Reset() *_RoleMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RoleMgr) Get() (result Role, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Role{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RoleMgr) Gets() (results []*Role, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Role{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RoleMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Role{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_RoleMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithNome nome获取
func (obj *_RoleMgr) WithNome(nome string) Option {
	return optionFunc(func(o *options) { o.query["nome"] = nome })
}

// GetByOption 功能选项模式获取
func (obj *_RoleMgr) GetByOption(opts ...Option) (result Role, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Role{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RoleMgr) GetByOptions(opts ...Option) (results []*Role, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Role{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_RoleMgr) GetFromID(id int64) (result Role, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Role{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_RoleMgr) GetBatchFromID(ids []int64) (results []*Role, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Role{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromNome 通过nome获取内容
func (obj *_RoleMgr) GetFromNome(nome string) (result Role, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Role{}).Where("`nome` = ?", nome).First(&result).Error

	return
}

// GetBatchFromNome 批量查找
func (obj *_RoleMgr) GetBatchFromNome(nomes []string) (results []*Role, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Role{}).Where("`nome` IN (?)", nomes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RoleMgr) FetchByPrimaryKey(id int64) (result Role, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Role{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByNomeUnique primary or index 获取唯一内容
func (obj *_RoleMgr) FetchUniqueByNomeUnique(nome string) (result Role, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Role{}).Where("`nome` = ?", nome).First(&result).Error

	return
}
