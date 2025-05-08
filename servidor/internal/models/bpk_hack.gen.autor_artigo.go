package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AutorArtigoMgr struct {
	*_BaseMgr
}

// AutorArtigoMgr open func
func AutorArtigoMgr(db *gorm.DB) *_AutorArtigoMgr {
	if db == nil {
		panic(fmt.Errorf("AutorArtigoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AutorArtigoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("autor_artigo"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_AutorArtigoMgr) Debug() *_AutorArtigoMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AutorArtigoMgr) GetTableName() string {
	return "autor_artigo"
}

// Reset 重置gorm会话
func (obj *_AutorArtigoMgr) Reset() *_AutorArtigoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AutorArtigoMgr) Get() (result AutorArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AutorArtigo{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("artigos").Where("id = ?", result.ArtigosID).Find(&result.Artigos).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("usuarios").Where("id = ?", result.UsuariosID).Find(&result.Usuarios).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_AutorArtigoMgr) Gets() (results []*AutorArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AutorArtigo{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("artigos").Where("id = ?", results[i].ArtigosID).Find(&results[i].Artigos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AutorArtigoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AutorArtigo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_AutorArtigoMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithArtigosID artigos_id获取
func (obj *_AutorArtigoMgr) WithArtigosID(artigosID int64) Option {
	return optionFunc(func(o *options) { o.query["artigos_id"] = artigosID })
}

// WithUsuariosID usuarios_id获取
func (obj *_AutorArtigoMgr) WithUsuariosID(usuariosID int64) Option {
	return optionFunc(func(o *options) { o.query["usuarios_id"] = usuariosID })
}

// WithOrdemAutoria ordem_autoria获取
func (obj *_AutorArtigoMgr) WithOrdemAutoria(ordemAutoria int) Option {
	return optionFunc(func(o *options) { o.query["ordem_autoria"] = ordemAutoria })
}

// WithCorrespondente correspondente获取
func (obj *_AutorArtigoMgr) WithCorrespondente(correspondente bool) Option {
	return optionFunc(func(o *options) { o.query["correspondente"] = correspondente })
}

// GetByOption 功能选项模式获取
func (obj *_AutorArtigoMgr) GetByOption(opts ...Option) (result AutorArtigo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AutorArtigo{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("artigos").Where("id = ?", result.ArtigosID).Find(&result.Artigos).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("usuarios").Where("id = ?", result.UsuariosID).Find(&result.Usuarios).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AutorArtigoMgr) GetByOptions(opts ...Option) (results []*AutorArtigo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AutorArtigo{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("artigos").Where("id = ?", results[i].ArtigosID).Find(&results[i].Artigos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_AutorArtigoMgr) GetFromID(id int64) (result AutorArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AutorArtigo{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("artigos").Where("id = ?", result.ArtigosID).Find(&result.Artigos).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("usuarios").Where("id = ?", result.UsuariosID).Find(&result.Usuarios).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_AutorArtigoMgr) GetBatchFromID(ids []int64) (results []*AutorArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AutorArtigo{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("artigos").Where("id = ?", results[i].ArtigosID).Find(&results[i].Artigos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromArtigosID 通过artigos_id获取内容
func (obj *_AutorArtigoMgr) GetFromArtigosID(artigosID int64) (results []*AutorArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AutorArtigo{}).Where("`artigos_id` = ?", artigosID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("artigos").Where("id = ?", results[i].ArtigosID).Find(&results[i].Artigos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromArtigosID 批量查找
func (obj *_AutorArtigoMgr) GetBatchFromArtigosID(artigosIDs []int64) (results []*AutorArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AutorArtigo{}).Where("`artigos_id` IN (?)", artigosIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("artigos").Where("id = ?", results[i].ArtigosID).Find(&results[i].Artigos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUsuariosID 通过usuarios_id获取内容
func (obj *_AutorArtigoMgr) GetFromUsuariosID(usuariosID int64) (results []*AutorArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AutorArtigo{}).Where("`usuarios_id` = ?", usuariosID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("artigos").Where("id = ?", results[i].ArtigosID).Find(&results[i].Artigos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUsuariosID 批量查找
func (obj *_AutorArtigoMgr) GetBatchFromUsuariosID(usuariosIDs []int64) (results []*AutorArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AutorArtigo{}).Where("`usuarios_id` IN (?)", usuariosIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("artigos").Where("id = ?", results[i].ArtigosID).Find(&results[i].Artigos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromOrdemAutoria 通过ordem_autoria获取内容
func (obj *_AutorArtigoMgr) GetFromOrdemAutoria(ordemAutoria int) (results []*AutorArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AutorArtigo{}).Where("`ordem_autoria` = ?", ordemAutoria).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("artigos").Where("id = ?", results[i].ArtigosID).Find(&results[i].Artigos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromOrdemAutoria 批量查找
func (obj *_AutorArtigoMgr) GetBatchFromOrdemAutoria(ordemAutorias []int) (results []*AutorArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AutorArtigo{}).Where("`ordem_autoria` IN (?)", ordemAutorias).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("artigos").Where("id = ?", results[i].ArtigosID).Find(&results[i].Artigos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCorrespondente 通过correspondente获取内容
func (obj *_AutorArtigoMgr) GetFromCorrespondente(correspondente bool) (results []*AutorArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AutorArtigo{}).Where("`correspondente` = ?", correspondente).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("artigos").Where("id = ?", results[i].ArtigosID).Find(&results[i].Artigos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCorrespondente 批量查找
func (obj *_AutorArtigoMgr) GetBatchFromCorrespondente(correspondentes []bool) (results []*AutorArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AutorArtigo{}).Where("`correspondente` IN (?)", correspondentes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("artigos").Where("id = ?", results[i].ArtigosID).Find(&results[i].Artigos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AutorArtigoMgr) FetchByPrimaryKey(id int64) (result AutorArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AutorArtigo{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("artigos").Where("id = ?", result.ArtigosID).Find(&result.Artigos).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("usuarios").Where("id = ?", result.UsuariosID).Find(&result.Usuarios).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByFkAutorArtigoArtigos1IDx  获取多个内容
func (obj *_AutorArtigoMgr) FetchIndexByFkAutorArtigoArtigos1IDx(artigosID int64) (results []*AutorArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AutorArtigo{}).Where("`artigos_id` = ?", artigosID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("artigos").Where("id = ?", results[i].ArtigosID).Find(&results[i].Artigos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByFkAutorArtigoUsuarios1IDx  获取多个内容
func (obj *_AutorArtigoMgr) FetchIndexByFkAutorArtigoUsuarios1IDx(usuariosID int64) (results []*AutorArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AutorArtigo{}).Where("`usuarios_id` = ?", usuariosID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("artigos").Where("id = ?", results[i].ArtigosID).Find(&results[i].Artigos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
