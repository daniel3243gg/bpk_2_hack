package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _EventosMgr struct {
	*_BaseMgr
}

// EventosMgr open func
func EventosMgr(db *gorm.DB) *_EventosMgr {
	if db == nil {
		panic(fmt.Errorf("EventosMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EventosMgr{_BaseMgr: &_BaseMgr{DB: db.Table("eventos"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_EventosMgr) Debug() *_EventosMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EventosMgr) GetTableName() string {
	return "eventos"
}

// Reset 重置gorm会话
func (obj *_EventosMgr) Reset() *_EventosMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EventosMgr) Get() (result Eventos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Eventos{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("usuarios").Where("id = ?", result.UsuariosID).Find(&result.Usuarios).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_EventosMgr) Gets() (results []*Eventos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Eventos{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_EventosMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Eventos{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_EventosMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithNome nome获取
func (obj *_EventosMgr) WithNome(nome string) Option {
	return optionFunc(func(o *options) { o.query["nome"] = nome })
}

// WithBanner banner获取
func (obj *_EventosMgr) WithBanner(banner []byte) Option {
	return optionFunc(func(o *options) { o.query["banner"] = banner })
}

// WithDescricao descricao获取
func (obj *_EventosMgr) WithDescricao(descricao string) Option {
	return optionFunc(func(o *options) { o.query["descricao"] = descricao })
}

// WithUsuariosID usuarios_id获取
func (obj *_EventosMgr) WithUsuariosID(usuariosID int64) Option {
	return optionFunc(func(o *options) { o.query["usuarios_id"] = usuariosID })
}

// WithDataInicio data_inicio获取
func (obj *_EventosMgr) WithDataInicio(dataInicio time.Time) Option {
	return optionFunc(func(o *options) { o.query["data_inicio"] = dataInicio })
}

// WithDataFim data_fim获取
func (obj *_EventosMgr) WithDataFim(dataFim time.Time) Option {
	return optionFunc(func(o *options) { o.query["data_fim"] = dataFim })
}

// WithStatus status获取
func (obj *_EventosMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_EventosMgr) GetByOption(opts ...Option) (result Eventos, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Eventos{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("usuarios").Where("id = ?", result.UsuariosID).Find(&result.Usuarios).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EventosMgr) GetByOptions(opts ...Option) (results []*Eventos, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Eventos{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_EventosMgr) GetFromID(id int64) (result Eventos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Eventos{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("usuarios").Where("id = ?", result.UsuariosID).Find(&result.Usuarios).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_EventosMgr) GetBatchFromID(ids []int64) (results []*Eventos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Eventos{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromNome 通过nome获取内容
func (obj *_EventosMgr) GetFromNome(nome string) (results []*Eventos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Eventos{}).Where("`nome` = ?", nome).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromNome 批量查找
func (obj *_EventosMgr) GetBatchFromNome(nomes []string) (results []*Eventos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Eventos{}).Where("`nome` IN (?)", nomes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromBanner 通过banner获取内容
func (obj *_EventosMgr) GetFromBanner(banner []byte) (results []*Eventos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Eventos{}).Where("`banner` = ?", banner).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromBanner 批量查找
func (obj *_EventosMgr) GetBatchFromBanner(banners [][]byte) (results []*Eventos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Eventos{}).Where("`banner` IN (?)", banners).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDescricao 通过descricao获取内容
func (obj *_EventosMgr) GetFromDescricao(descricao string) (results []*Eventos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Eventos{}).Where("`descricao` = ?", descricao).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDescricao 批量查找
func (obj *_EventosMgr) GetBatchFromDescricao(descricaos []string) (results []*Eventos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Eventos{}).Where("`descricao` IN (?)", descricaos).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_EventosMgr) GetFromUsuariosID(usuariosID int64) (results []*Eventos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Eventos{}).Where("`usuarios_id` = ?", usuariosID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_EventosMgr) GetBatchFromUsuariosID(usuariosIDs []int64) (results []*Eventos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Eventos{}).Where("`usuarios_id` IN (?)", usuariosIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDataInicio 通过data_inicio获取内容
func (obj *_EventosMgr) GetFromDataInicio(dataInicio time.Time) (results []*Eventos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Eventos{}).Where("`data_inicio` = ?", dataInicio).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDataInicio 批量查找
func (obj *_EventosMgr) GetBatchFromDataInicio(dataInicios []time.Time) (results []*Eventos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Eventos{}).Where("`data_inicio` IN (?)", dataInicios).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDataFim 通过data_fim获取内容
func (obj *_EventosMgr) GetFromDataFim(dataFim time.Time) (results []*Eventos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Eventos{}).Where("`data_fim` = ?", dataFim).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDataFim 批量查找
func (obj *_EventosMgr) GetBatchFromDataFim(dataFims []time.Time) (results []*Eventos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Eventos{}).Where("`data_fim` IN (?)", dataFims).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStatus 通过status获取内容
func (obj *_EventosMgr) GetFromStatus(status string) (results []*Eventos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Eventos{}).Where("`status` = ?", status).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStatus 批量查找
func (obj *_EventosMgr) GetBatchFromStatus(statuss []string) (results []*Eventos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Eventos{}).Where("`status` IN (?)", statuss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_EventosMgr) FetchByPrimaryKey(id int64) (result Eventos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Eventos{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("usuarios").Where("id = ?", result.UsuariosID).Find(&result.Usuarios).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByFkEventosUsuarios1IDx  获取多个内容
func (obj *_EventosMgr) FetchIndexByFkEventosUsuarios1IDx(usuariosID int64) (results []*Eventos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Eventos{}).Where("`usuarios_id` = ?", usuariosID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
