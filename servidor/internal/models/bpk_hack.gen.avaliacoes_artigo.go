package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AvaliacoesArtigoMgr struct {
	*_BaseMgr
}

// AvaliacoesArtigoMgr open func
func AvaliacoesArtigoMgr(db *gorm.DB) *_AvaliacoesArtigoMgr {
	if db == nil {
		panic(fmt.Errorf("AvaliacoesArtigoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AvaliacoesArtigoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("avaliacoes_artigo"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_AvaliacoesArtigoMgr) Debug() *_AvaliacoesArtigoMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AvaliacoesArtigoMgr) GetTableName() string {
	return "avaliacoes_artigo"
}

// Reset 重置gorm会话
func (obj *_AvaliacoesArtigoMgr) Reset() *_AvaliacoesArtigoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AvaliacoesArtigoMgr) Get() (result AvaliacoesArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliacoesArtigo{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("avaliador_evento").Where("id = ?", result.AvaliadorEventoID).Find(&result.AvaliadorEvento).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("artigos_evento").Where("id = ?", result.ArtigoEventoID).Find(&result.ArtigosEvento).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_AvaliacoesArtigoMgr) Gets() (results []*AvaliacoesArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliacoesArtigo{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("avaliador_evento").Where("id = ?", results[i].AvaliadorEventoID).Find(&results[i].AvaliadorEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos_evento").Where("id = ?", results[i].ArtigoEventoID).Find(&results[i].ArtigosEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AvaliacoesArtigoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AvaliacoesArtigo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_AvaliacoesArtigoMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAvaliadorEventoID avaliador_evento_id获取
func (obj *_AvaliacoesArtigoMgr) WithAvaliadorEventoID(avaliadorEventoID int64) Option {
	return optionFunc(func(o *options) { o.query["avaliador_evento_id"] = avaliadorEventoID })
}

// WithArtigoEventoID artigo_evento_id获取
func (obj *_AvaliacoesArtigoMgr) WithArtigoEventoID(artigoEventoID int64) Option {
	return optionFunc(func(o *options) { o.query["artigo_evento_id"] = artigoEventoID })
}

// WithAvaliacao avaliacao获取
func (obj *_AvaliacoesArtigoMgr) WithAvaliacao(avaliacao string) Option {
	return optionFunc(func(o *options) { o.query["avaliacao"] = avaliacao })
}

// WithNota nota获取
func (obj *_AvaliacoesArtigoMgr) WithNota(nota float64) Option {
	return optionFunc(func(o *options) { o.query["nota"] = nota })
}

// WithDataAvaliacao data_avaliacao获取
func (obj *_AvaliacoesArtigoMgr) WithDataAvaliacao(dataAvaliacao time.Time) Option {
	return optionFunc(func(o *options) { o.query["data_avaliacao"] = dataAvaliacao })
}

// WithRecomendacao recomendacao获取
func (obj *_AvaliacoesArtigoMgr) WithRecomendacao(recomendacao string) Option {
	return optionFunc(func(o *options) { o.query["recomendacao"] = recomendacao })
}

// GetByOption 功能选项模式获取
func (obj *_AvaliacoesArtigoMgr) GetByOption(opts ...Option) (result AvaliacoesArtigo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AvaliacoesArtigo{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("avaliador_evento").Where("id = ?", result.AvaliadorEventoID).Find(&result.AvaliadorEvento).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("artigos_evento").Where("id = ?", result.ArtigoEventoID).Find(&result.ArtigosEvento).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AvaliacoesArtigoMgr) GetByOptions(opts ...Option) (results []*AvaliacoesArtigo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AvaliacoesArtigo{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("avaliador_evento").Where("id = ?", results[i].AvaliadorEventoID).Find(&results[i].AvaliadorEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos_evento").Where("id = ?", results[i].ArtigoEventoID).Find(&results[i].ArtigosEvento).Error; err != nil { //
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
func (obj *_AvaliacoesArtigoMgr) GetFromID(id int64) (result AvaliacoesArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliacoesArtigo{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("avaliador_evento").Where("id = ?", result.AvaliadorEventoID).Find(&result.AvaliadorEvento).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("artigos_evento").Where("id = ?", result.ArtigoEventoID).Find(&result.ArtigosEvento).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_AvaliacoesArtigoMgr) GetBatchFromID(ids []int64) (results []*AvaliacoesArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliacoesArtigo{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("avaliador_evento").Where("id = ?", results[i].AvaliadorEventoID).Find(&results[i].AvaliadorEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos_evento").Where("id = ?", results[i].ArtigoEventoID).Find(&results[i].ArtigosEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromAvaliadorEventoID 通过avaliador_evento_id获取内容
func (obj *_AvaliacoesArtigoMgr) GetFromAvaliadorEventoID(avaliadorEventoID int64) (results []*AvaliacoesArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliacoesArtigo{}).Where("`avaliador_evento_id` = ?", avaliadorEventoID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("avaliador_evento").Where("id = ?", results[i].AvaliadorEventoID).Find(&results[i].AvaliadorEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos_evento").Where("id = ?", results[i].ArtigoEventoID).Find(&results[i].ArtigosEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromAvaliadorEventoID 批量查找
func (obj *_AvaliacoesArtigoMgr) GetBatchFromAvaliadorEventoID(avaliadorEventoIDs []int64) (results []*AvaliacoesArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliacoesArtigo{}).Where("`avaliador_evento_id` IN (?)", avaliadorEventoIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("avaliador_evento").Where("id = ?", results[i].AvaliadorEventoID).Find(&results[i].AvaliadorEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos_evento").Where("id = ?", results[i].ArtigoEventoID).Find(&results[i].ArtigosEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromArtigoEventoID 通过artigo_evento_id获取内容
func (obj *_AvaliacoesArtigoMgr) GetFromArtigoEventoID(artigoEventoID int64) (results []*AvaliacoesArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliacoesArtigo{}).Where("`artigo_evento_id` = ?", artigoEventoID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("avaliador_evento").Where("id = ?", results[i].AvaliadorEventoID).Find(&results[i].AvaliadorEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos_evento").Where("id = ?", results[i].ArtigoEventoID).Find(&results[i].ArtigosEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromArtigoEventoID 批量查找
func (obj *_AvaliacoesArtigoMgr) GetBatchFromArtigoEventoID(artigoEventoIDs []int64) (results []*AvaliacoesArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliacoesArtigo{}).Where("`artigo_evento_id` IN (?)", artigoEventoIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("avaliador_evento").Where("id = ?", results[i].AvaliadorEventoID).Find(&results[i].AvaliadorEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos_evento").Where("id = ?", results[i].ArtigoEventoID).Find(&results[i].ArtigosEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromAvaliacao 通过avaliacao获取内容
func (obj *_AvaliacoesArtigoMgr) GetFromAvaliacao(avaliacao string) (results []*AvaliacoesArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliacoesArtigo{}).Where("`avaliacao` = ?", avaliacao).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("avaliador_evento").Where("id = ?", results[i].AvaliadorEventoID).Find(&results[i].AvaliadorEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos_evento").Where("id = ?", results[i].ArtigoEventoID).Find(&results[i].ArtigosEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromAvaliacao 批量查找
func (obj *_AvaliacoesArtigoMgr) GetBatchFromAvaliacao(avaliacaos []string) (results []*AvaliacoesArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliacoesArtigo{}).Where("`avaliacao` IN (?)", avaliacaos).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("avaliador_evento").Where("id = ?", results[i].AvaliadorEventoID).Find(&results[i].AvaliadorEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos_evento").Where("id = ?", results[i].ArtigoEventoID).Find(&results[i].ArtigosEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromNota 通过nota获取内容
func (obj *_AvaliacoesArtigoMgr) GetFromNota(nota float64) (results []*AvaliacoesArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliacoesArtigo{}).Where("`nota` = ?", nota).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("avaliador_evento").Where("id = ?", results[i].AvaliadorEventoID).Find(&results[i].AvaliadorEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos_evento").Where("id = ?", results[i].ArtigoEventoID).Find(&results[i].ArtigosEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromNota 批量查找
func (obj *_AvaliacoesArtigoMgr) GetBatchFromNota(notas []float64) (results []*AvaliacoesArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliacoesArtigo{}).Where("`nota` IN (?)", notas).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("avaliador_evento").Where("id = ?", results[i].AvaliadorEventoID).Find(&results[i].AvaliadorEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos_evento").Where("id = ?", results[i].ArtigoEventoID).Find(&results[i].ArtigosEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDataAvaliacao 通过data_avaliacao获取内容
func (obj *_AvaliacoesArtigoMgr) GetFromDataAvaliacao(dataAvaliacao time.Time) (results []*AvaliacoesArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliacoesArtigo{}).Where("`data_avaliacao` = ?", dataAvaliacao).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("avaliador_evento").Where("id = ?", results[i].AvaliadorEventoID).Find(&results[i].AvaliadorEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos_evento").Where("id = ?", results[i].ArtigoEventoID).Find(&results[i].ArtigosEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDataAvaliacao 批量查找
func (obj *_AvaliacoesArtigoMgr) GetBatchFromDataAvaliacao(dataAvaliacaos []time.Time) (results []*AvaliacoesArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliacoesArtigo{}).Where("`data_avaliacao` IN (?)", dataAvaliacaos).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("avaliador_evento").Where("id = ?", results[i].AvaliadorEventoID).Find(&results[i].AvaliadorEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos_evento").Where("id = ?", results[i].ArtigoEventoID).Find(&results[i].ArtigosEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRecomendacao 通过recomendacao获取内容
func (obj *_AvaliacoesArtigoMgr) GetFromRecomendacao(recomendacao string) (results []*AvaliacoesArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliacoesArtigo{}).Where("`recomendacao` = ?", recomendacao).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("avaliador_evento").Where("id = ?", results[i].AvaliadorEventoID).Find(&results[i].AvaliadorEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos_evento").Where("id = ?", results[i].ArtigoEventoID).Find(&results[i].ArtigosEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRecomendacao 批量查找
func (obj *_AvaliacoesArtigoMgr) GetBatchFromRecomendacao(recomendacaos []string) (results []*AvaliacoesArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliacoesArtigo{}).Where("`recomendacao` IN (?)", recomendacaos).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("avaliador_evento").Where("id = ?", results[i].AvaliadorEventoID).Find(&results[i].AvaliadorEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos_evento").Where("id = ?", results[i].ArtigoEventoID).Find(&results[i].ArtigosEvento).Error; err != nil { //
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
func (obj *_AvaliacoesArtigoMgr) FetchByPrimaryKey(id int64) (result AvaliacoesArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliacoesArtigo{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("avaliador_evento").Where("id = ?", result.AvaliadorEventoID).Find(&result.AvaliadorEvento).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("artigos_evento").Where("id = ?", result.ArtigoEventoID).Find(&result.ArtigosEvento).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueIndexByUniqueAvaliacao primary or index 获取唯一内容
func (obj *_AvaliacoesArtigoMgr) FetchUniqueIndexByUniqueAvaliacao(avaliadorEventoID int64, artigoEventoID int64) (result AvaliacoesArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliacoesArtigo{}).Where("`avaliador_evento_id` = ? AND `artigo_evento_id` = ?", avaliadorEventoID, artigoEventoID).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("avaliador_evento").Where("id = ?", result.AvaliadorEventoID).Find(&result.AvaliadorEvento).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("artigos_evento").Where("id = ?", result.ArtigoEventoID).Find(&result.ArtigosEvento).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByFkAvaliacoesArtigoAvaliadorEvento1IDx  获取多个内容
func (obj *_AvaliacoesArtigoMgr) FetchIndexByFkAvaliacoesArtigoAvaliadorEvento1IDx(avaliadorEventoID int64) (results []*AvaliacoesArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliacoesArtigo{}).Where("`avaliador_evento_id` = ?", avaliadorEventoID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("avaliador_evento").Where("id = ?", results[i].AvaliadorEventoID).Find(&results[i].AvaliadorEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos_evento").Where("id = ?", results[i].ArtigoEventoID).Find(&results[i].ArtigosEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByFkAvaliacoesArtigoArtigosEvento1IDx  获取多个内容
func (obj *_AvaliacoesArtigoMgr) FetchIndexByFkAvaliacoesArtigoArtigosEvento1IDx(artigoEventoID int64) (results []*AvaliacoesArtigo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AvaliacoesArtigo{}).Where("`artigo_evento_id` = ?", artigoEventoID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("avaliador_evento").Where("id = ?", results[i].AvaliadorEventoID).Find(&results[i].AvaliadorEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("artigos_evento").Where("id = ?", results[i].ArtigoEventoID).Find(&results[i].ArtigosEvento).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
