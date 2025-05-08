package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AvaliadorEventoMgr struct {
	*_BaseMgr
}

// AvaliadorEventoMgr open func
func AvaliadorEventoMgr(db *gorm.DB) *_AvaliadorEventoMgr {
	if db == nil {
		panic(fmt.Errorf("AvaliadorEventoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AvaliadorEventoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("avaliador_evento"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_AvaliadorEventoMgr) Debug() *_AvaliadorEventoMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AvaliadorEventoMgr) GetTableName() string {
	return "avaliador_evento"
}

// Reset 重置gorm会话
func (obj *_AvaliadorEventoMgr) Reset() *_AvaliadorEventoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AvaliadorEventoMgr) Get() (result AvaliadorEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliadorEvento{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("usuarios").Where("id = ?", result.UsuariosID).Find(&result.Usuarios).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("eventos").Where("id = ?", result.EventoID).Find(&result.Eventos).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_AvaliadorEventoMgr) Gets() (results []*AvaliadorEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliadorEvento{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("eventos").Where("id = ?", results[i].EventoID).Find(&results[i].Eventos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AvaliadorEventoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AvaliadorEvento{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_AvaliadorEventoMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUsuariosID usuarios_id获取
func (obj *_AvaliadorEventoMgr) WithUsuariosID(usuariosID int64) Option {
	return optionFunc(func(o *options) { o.query["usuarios_id"] = usuariosID })
}

// WithEventoID evento_id获取
func (obj *_AvaliadorEventoMgr) WithEventoID(eventoID int64) Option {
	return optionFunc(func(o *options) { o.query["evento_id"] = eventoID })
}

// WithAreaEspecializacao area_especializacao获取
func (obj *_AvaliadorEventoMgr) WithAreaEspecializacao(areaEspecializacao string) Option {
	return optionFunc(func(o *options) { o.query["area_especializacao"] = areaEspecializacao })
}

// GetByOption 功能选项模式获取
func (obj *_AvaliadorEventoMgr) GetByOption(opts ...Option) (result AvaliadorEvento, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AvaliadorEvento{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("usuarios").Where("id = ?", result.UsuariosID).Find(&result.Usuarios).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("eventos").Where("id = ?", result.EventoID).Find(&result.Eventos).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AvaliadorEventoMgr) GetByOptions(opts ...Option) (results []*AvaliadorEvento, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AvaliadorEvento{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("eventos").Where("id = ?", results[i].EventoID).Find(&results[i].Eventos).Error; err != nil { //
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
func (obj *_AvaliadorEventoMgr) GetFromID(id int64) (result AvaliadorEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliadorEvento{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("usuarios").Where("id = ?", result.UsuariosID).Find(&result.Usuarios).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("eventos").Where("id = ?", result.EventoID).Find(&result.Eventos).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_AvaliadorEventoMgr) GetBatchFromID(ids []int64) (results []*AvaliadorEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliadorEvento{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("eventos").Where("id = ?", results[i].EventoID).Find(&results[i].Eventos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUsuariosID 通过usuarios_id获取内容
func (obj *_AvaliadorEventoMgr) GetFromUsuariosID(usuariosID int64) (results []*AvaliadorEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliadorEvento{}).Where("`usuarios_id` = ?", usuariosID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("eventos").Where("id = ?", results[i].EventoID).Find(&results[i].Eventos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUsuariosID 批量查找
func (obj *_AvaliadorEventoMgr) GetBatchFromUsuariosID(usuariosIDs []int64) (results []*AvaliadorEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliadorEvento{}).Where("`usuarios_id` IN (?)", usuariosIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("eventos").Where("id = ?", results[i].EventoID).Find(&results[i].Eventos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromEventoID 通过evento_id获取内容
func (obj *_AvaliadorEventoMgr) GetFromEventoID(eventoID int64) (results []*AvaliadorEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliadorEvento{}).Where("`evento_id` = ?", eventoID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("eventos").Where("id = ?", results[i].EventoID).Find(&results[i].Eventos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromEventoID 批量查找
func (obj *_AvaliadorEventoMgr) GetBatchFromEventoID(eventoIDs []int64) (results []*AvaliadorEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliadorEvento{}).Where("`evento_id` IN (?)", eventoIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("eventos").Where("id = ?", results[i].EventoID).Find(&results[i].Eventos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromAreaEspecializacao 通过area_especializacao获取内容
func (obj *_AvaliadorEventoMgr) GetFromAreaEspecializacao(areaEspecializacao string) (results []*AvaliadorEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliadorEvento{}).Where("`area_especializacao` = ?", areaEspecializacao).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("eventos").Where("id = ?", results[i].EventoID).Find(&results[i].Eventos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromAreaEspecializacao 批量查找
func (obj *_AvaliadorEventoMgr) GetBatchFromAreaEspecializacao(areaEspecializacaos []string) (results []*AvaliadorEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliadorEvento{}).Where("`area_especializacao` IN (?)", areaEspecializacaos).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("eventos").Where("id = ?", results[i].EventoID).Find(&results[i].Eventos).Error; err != nil { //
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
func (obj *_AvaliadorEventoMgr) FetchByPrimaryKey(id int64) (result AvaliadorEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliadorEvento{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("usuarios").Where("id = ?", result.UsuariosID).Find(&result.Usuarios).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("eventos").Where("id = ?", result.EventoID).Find(&result.Eventos).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueIndexByUniqueAvaliadorEvento primary or index 获取唯一内容
func (obj *_AvaliadorEventoMgr) FetchUniqueIndexByUniqueAvaliadorEvento(usuariosID int64, eventoID int64) (result AvaliadorEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliadorEvento{}).Where("`usuarios_id` = ? AND `evento_id` = ?", usuariosID, eventoID).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("usuarios").Where("id = ?", result.UsuariosID).Find(&result.Usuarios).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("eventos").Where("id = ?", result.EventoID).Find(&result.Eventos).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByFkAvaliadorEventoUsuarios1IDx  获取多个内容
func (obj *_AvaliadorEventoMgr) FetchIndexByFkAvaliadorEventoUsuarios1IDx(usuariosID int64) (results []*AvaliadorEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliadorEvento{}).Where("`usuarios_id` = ?", usuariosID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("eventos").Where("id = ?", results[i].EventoID).Find(&results[i].Eventos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByFkAvaliadorEventoEventos1IDx  获取多个内容
func (obj *_AvaliadorEventoMgr) FetchIndexByFkAvaliadorEventoEventos1IDx(eventoID int64) (results []*AvaliadorEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliadorEvento{}).Where("`evento_id` = ?", eventoID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("usuarios").Where("id = ?", results[i].UsuariosID).Find(&results[i].Usuarios).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("eventos").Where("id = ?", results[i].EventoID).Find(&results[i].Eventos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
