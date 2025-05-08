package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ArtigosMgr struct {
	*_BaseMgr
}

// ArtigosMgr open func
func ArtigosMgr(db *gorm.DB) *_ArtigosMgr {
	if db == nil {
		panic(fmt.Errorf("ArtigosMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ArtigosMgr{_BaseMgr: &_BaseMgr{DB: db.Table("artigos"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_ArtigosMgr) Debug() *_ArtigosMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ArtigosMgr) GetTableName() string {
	return "artigos"
}

// Reset 重置gorm会话
func (obj *_ArtigosMgr) Reset() *_ArtigosMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ArtigosMgr) Get() (result Artigos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Artigos{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ArtigosMgr) Gets() (results []*Artigos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Artigos{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ArtigosMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Artigos{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_ArtigosMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithArquivo arquivo获取
func (obj *_ArtigosMgr) WithArquivo(arquivo []byte) Option {
	return optionFunc(func(o *options) { o.query["arquivo"] = arquivo })
}

// WithTitulo titulo获取
func (obj *_ArtigosMgr) WithTitulo(titulo string) Option {
	return optionFunc(func(o *options) { o.query["titulo"] = titulo })
}

// WithResumo resumo获取
func (obj *_ArtigosMgr) WithResumo(resumo string) Option {
	return optionFunc(func(o *options) { o.query["resumo"] = resumo })
}

// WithPalavrasChave palavras_chave获取
func (obj *_ArtigosMgr) WithPalavrasChave(palavrasChave string) Option {
	return optionFunc(func(o *options) { o.query["palavras_chave"] = palavrasChave })
}

// WithAreaTematica area_tematica获取
func (obj *_ArtigosMgr) WithAreaTematica(areaTematica string) Option {
	return optionFunc(func(o *options) { o.query["area_tematica"] = areaTematica })
}

// WithStatus status获取
func (obj *_ArtigosMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithDataSubmissao data_submissao获取
func (obj *_ArtigosMgr) WithDataSubmissao(dataSubmissao time.Time) Option {
	return optionFunc(func(o *options) { o.query["data_submissao"] = dataSubmissao })
}

// GetByOption 功能选项模式获取
func (obj *_ArtigosMgr) GetByOption(opts ...Option) (result Artigos, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Artigos{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ArtigosMgr) GetByOptions(opts ...Option) (results []*Artigos, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Artigos{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_ArtigosMgr) GetFromID(id int64) (result Artigos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Artigos{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_ArtigosMgr) GetBatchFromID(ids []int64) (results []*Artigos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Artigos{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromArquivo 通过arquivo获取内容
func (obj *_ArtigosMgr) GetFromArquivo(arquivo []byte) (results []*Artigos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Artigos{}).Where("`arquivo` = ?", arquivo).Find(&results).Error

	return
}

// GetBatchFromArquivo 批量查找
func (obj *_ArtigosMgr) GetBatchFromArquivo(arquivos [][]byte) (results []*Artigos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Artigos{}).Where("`arquivo` IN (?)", arquivos).Find(&results).Error

	return
}

// GetFromTitulo 通过titulo获取内容
func (obj *_ArtigosMgr) GetFromTitulo(titulo string) (results []*Artigos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Artigos{}).Where("`titulo` = ?", titulo).Find(&results).Error

	return
}

// GetBatchFromTitulo 批量查找
func (obj *_ArtigosMgr) GetBatchFromTitulo(titulos []string) (results []*Artigos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Artigos{}).Where("`titulo` IN (?)", titulos).Find(&results).Error

	return
}

// GetFromResumo 通过resumo获取内容
func (obj *_ArtigosMgr) GetFromResumo(resumo string) (results []*Artigos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Artigos{}).Where("`resumo` = ?", resumo).Find(&results).Error

	return
}

// GetBatchFromResumo 批量查找
func (obj *_ArtigosMgr) GetBatchFromResumo(resumos []string) (results []*Artigos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Artigos{}).Where("`resumo` IN (?)", resumos).Find(&results).Error

	return
}

// GetFromPalavrasChave 通过palavras_chave获取内容
func (obj *_ArtigosMgr) GetFromPalavrasChave(palavrasChave string) (results []*Artigos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Artigos{}).Where("`palavras_chave` = ?", palavrasChave).Find(&results).Error

	return
}

// GetBatchFromPalavrasChave 批量查找
func (obj *_ArtigosMgr) GetBatchFromPalavrasChave(palavrasChaves []string) (results []*Artigos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Artigos{}).Where("`palavras_chave` IN (?)", palavrasChaves).Find(&results).Error

	return
}

// GetFromAreaTematica 通过area_tematica获取内容
func (obj *_ArtigosMgr) GetFromAreaTematica(areaTematica string) (results []*Artigos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Artigos{}).Where("`area_tematica` = ?", areaTematica).Find(&results).Error

	return
}

// GetBatchFromAreaTematica 批量查找
func (obj *_ArtigosMgr) GetBatchFromAreaTematica(areaTematicas []string) (results []*Artigos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Artigos{}).Where("`area_tematica` IN (?)", areaTematicas).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_ArtigosMgr) GetFromStatus(status string) (results []*Artigos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Artigos{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_ArtigosMgr) GetBatchFromStatus(statuss []string) (results []*Artigos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Artigos{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromDataSubmissao 通过data_submissao获取内容
func (obj *_ArtigosMgr) GetFromDataSubmissao(dataSubmissao time.Time) (results []*Artigos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Artigos{}).Where("`data_submissao` = ?", dataSubmissao).Find(&results).Error

	return
}

// GetBatchFromDataSubmissao 批量查找
func (obj *_ArtigosMgr) GetBatchFromDataSubmissao(dataSubmissaos []time.Time) (results []*Artigos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Artigos{}).Where("`data_submissao` IN (?)", dataSubmissaos).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ArtigosMgr) FetchByPrimaryKey(id int64) (result Artigos, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Artigos{}).Where("`id` = ?", id).First(&result).Error

	return
}
