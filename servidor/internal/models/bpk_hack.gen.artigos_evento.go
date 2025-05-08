package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _ArtigosEventoMgr struct {
	*_BaseMgr
}

// ArtigosEventoMgr open func
func ArtigosEventoMgr(db *gorm.DB) *_ArtigosEventoMgr {
	if db == nil {
		panic(fmt.Errorf("ArtigosEventoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ArtigosEventoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("artigos_evento"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_ArtigosEventoMgr) Debug() *_ArtigosEventoMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ArtigosEventoMgr) GetTableName() string {
	return "artigos_evento"
}

// Reset 重置gorm会话
func (obj *_ArtigosEventoMgr) Reset() *_ArtigosEventoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ArtigosEventoMgr) Get() (result ArtigosEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ArtigosEvento{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("eventos").Where("id = ?", result.EventosID).Find(&result.Eventos).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("artigos").Where("id = ?", result.ArtigoID).Find(&result.Artigos).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_ArtigosEventoMgr) Gets() (results []*ArtigosEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ArtigosEvento{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("eventos").Where("id = ?", results[i].EventosID).Find(&results[i].Eventos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos").Where("id = ?", results[i].ArtigoID).Find(&results[i].Artigos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ArtigosEventoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ArtigosEvento{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ArtigosEventoMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithEventosID eventos_id获取
func (obj *_ArtigosEventoMgr) WithEventosID(eventosID int64) Option {
	return optionFunc(func(o *options) { o.query["eventos_id"] = eventosID })
}

// WithArtigoID artigo_id获取
func (obj *_ArtigosEventoMgr) WithArtigoID(artigoID int64) Option {
	return optionFunc(func(o *options) { o.query["artigo_id"] = artigoID })
}

// GetByOption 功能选项模式获取
func (obj *_ArtigosEventoMgr) GetByOption(opts ...Option) (result ArtigosEvento, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ArtigosEvento{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("eventos").Where("id = ?", result.EventosID).Find(&result.Eventos).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("artigos").Where("id = ?", result.ArtigoID).Find(&result.Artigos).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ArtigosEventoMgr) GetByOptions(opts ...Option) (results []*ArtigosEvento, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ArtigosEvento{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("eventos").Where("id = ?", results[i].EventosID).Find(&results[i].Eventos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos").Where("id = ?", results[i].ArtigoID).Find(&results[i].Artigos).Error; err != nil { //
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
func (obj *_ArtigosEventoMgr) GetFromID(id int64) (result ArtigosEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ArtigosEvento{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("eventos").Where("id = ?", result.EventosID).Find(&result.Eventos).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("artigos").Where("id = ?", result.ArtigoID).Find(&result.Artigos).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_ArtigosEventoMgr) GetBatchFromID(ids []int64) (results []*ArtigosEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ArtigosEvento{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("eventos").Where("id = ?", results[i].EventosID).Find(&results[i].Eventos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos").Where("id = ?", results[i].ArtigoID).Find(&results[i].Artigos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromEventosID 通过eventos_id获取内容
func (obj *_ArtigosEventoMgr) GetFromEventosID(eventosID int64) (results []*ArtigosEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ArtigosEvento{}).Where("`eventos_id` = ?", eventosID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("eventos").Where("id = ?", results[i].EventosID).Find(&results[i].Eventos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos").Where("id = ?", results[i].ArtigoID).Find(&results[i].Artigos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromEventosID 批量查找
func (obj *_ArtigosEventoMgr) GetBatchFromEventosID(eventosIDs []int64) (results []*ArtigosEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ArtigosEvento{}).Where("`eventos_id` IN (?)", eventosIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("eventos").Where("id = ?", results[i].EventosID).Find(&results[i].Eventos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos").Where("id = ?", results[i].ArtigoID).Find(&results[i].Artigos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromArtigoID 通过artigo_id获取内容
func (obj *_ArtigosEventoMgr) GetFromArtigoID(artigoID int64) (results []*ArtigosEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ArtigosEvento{}).Where("`artigo_id` = ?", artigoID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("eventos").Where("id = ?", results[i].EventosID).Find(&results[i].Eventos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos").Where("id = ?", results[i].ArtigoID).Find(&results[i].Artigos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromArtigoID 批量查找
func (obj *_ArtigosEventoMgr) GetBatchFromArtigoID(artigoIDs []int64) (results []*ArtigosEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ArtigosEvento{}).Where("`artigo_id` IN (?)", artigoIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("eventos").Where("id = ?", results[i].EventosID).Find(&results[i].Eventos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos").Where("id = ?", results[i].ArtigoID).Find(&results[i].Artigos).Error; err != nil { //
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
func (obj *_ArtigosEventoMgr) FetchByPrimaryKey(id int64) (result ArtigosEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ArtigosEvento{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("eventos").Where("id = ?", result.EventosID).Find(&result.Eventos).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("artigos").Where("id = ?", result.ArtigoID).Find(&result.Artigos).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueIndexByUniqueArtigoEvento primary or index 获取唯一内容
func (obj *_ArtigosEventoMgr) FetchUniqueIndexByUniqueArtigoEvento(eventosID int64, artigoID int64) (result ArtigosEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ArtigosEvento{}).Where("`eventos_id` = ? AND `artigo_id` = ?", eventosID, artigoID).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("eventos").Where("id = ?", result.EventosID).Find(&result.Eventos).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("artigos").Where("id = ?", result.ArtigoID).Find(&result.Artigos).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByFkArtigoEventosEventos1IDx  获取多个内容
func (obj *_ArtigosEventoMgr) FetchIndexByFkArtigoEventosEventos1IDx(eventosID int64) (results []*ArtigosEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ArtigosEvento{}).Where("`eventos_id` = ?", eventosID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("eventos").Where("id = ?", results[i].EventosID).Find(&results[i].Eventos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos").Where("id = ?", results[i].ArtigoID).Find(&results[i].Artigos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByFkArtigoEventosArtigo1IDx  获取多个内容
func (obj *_ArtigosEventoMgr) FetchIndexByFkArtigoEventosArtigo1IDx(artigoID int64) (results []*ArtigosEvento, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ArtigosEvento{}).Where("`artigo_id` = ?", artigoID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("eventos").Where("id = ?", results[i].EventosID).Find(&results[i].Eventos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos").Where("id = ?", results[i].ArtigoID).Find(&results[i].Artigos).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
