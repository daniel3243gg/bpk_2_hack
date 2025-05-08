package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UsuariosMgr struct {
	*_BaseMgr
}

// UsuariosMgr open func
func UsuariosMgr(db *gorm.DB) *_UsuariosMgr {
	if db == nil {
		panic(fmt.Errorf("UsuariosMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UsuariosMgr{_BaseMgr: &_BaseMgr{DB: db.Table("usuarios"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// Debug open debug.打开debug模式查看sql语句
func (obj *_UsuariosMgr) Debug() *_UsuariosMgr {
	obj._BaseMgr.DB = obj._BaseMgr.DB.Debug()
	return obj
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UsuariosMgr) GetTableName() string {
	return "usuarios"
}

// Reset 重置gorm会话
func (obj *_UsuariosMgr) Reset() *_UsuariosMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UsuariosMgr) Get() (result Usuarios, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Usuarios{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("role").Where("id = ?", result.RoleID).Find(&result.Role).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_UsuariosMgr) Gets() (results []*Usuarios, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Usuarios{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("role").Where("id = ?", results[i].RoleID).Find(&results[i].Role).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UsuariosMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Usuarios{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UsuariosMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithNome nome获取
func (obj *_UsuariosMgr) WithNome(nome string) Option {
	return optionFunc(func(o *options) { o.query["nome"] = nome })
}

// WithEmail email获取
func (obj *_UsuariosMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithSenha senha获取
func (obj *_UsuariosMgr) WithSenha(senha string) Option {
	return optionFunc(func(o *options) { o.query["senha"] = senha })
}

// WithDataCadastro data_cadastro获取
func (obj *_UsuariosMgr) WithDataCadastro(dataCadastro time.Time) Option {
	return optionFunc(func(o *options) { o.query["data_cadastro"] = dataCadastro })
}

// WithRoleID role_id获取
func (obj *_UsuariosMgr) WithRoleID(roleID int64) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// WithTelefone telefone获取
func (obj *_UsuariosMgr) WithTelefone(telefone string) Option {
	return optionFunc(func(o *options) { o.query["telefone"] = telefone })
}

// WithArea area获取
func (obj *_UsuariosMgr) WithArea(area string) Option {
	return optionFunc(func(o *options) { o.query["area"] = area })
}

// GetByOption 功能选项模式获取
func (obj *_UsuariosMgr) GetByOption(opts ...Option) (result Usuarios, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Usuarios{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("role").Where("id = ?", result.RoleID).Find(&result.Role).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UsuariosMgr) GetByOptions(opts ...Option) (results []*Usuarios, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Usuarios{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("role").Where("id = ?", results[i].RoleID).Find(&results[i].Role).Error; err != nil { //
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
func (obj *_UsuariosMgr) GetFromID(id int64) (result Usuarios, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Usuarios{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("role").Where("id = ?", result.RoleID).Find(&result.Role).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_UsuariosMgr) GetBatchFromID(ids []int64) (results []*Usuarios, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Usuarios{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("role").Where("id = ?", results[i].RoleID).Find(&results[i].Role).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromNome 通过nome获取内容
func (obj *_UsuariosMgr) GetFromNome(nome string) (results []*Usuarios, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Usuarios{}).Where("`nome` = ?", nome).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("role").Where("id = ?", results[i].RoleID).Find(&results[i].Role).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromNome 批量查找
func (obj *_UsuariosMgr) GetBatchFromNome(nomes []string) (results []*Usuarios, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Usuarios{}).Where("`nome` IN (?)", nomes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("role").Where("id = ?", results[i].RoleID).Find(&results[i].Role).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromEmail 通过email获取内容
func (obj *_UsuariosMgr) GetFromEmail(email string) (result Usuarios, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Usuarios{}).Where("`email` = ?", email).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("role").Where("id = ?", result.RoleID).Find(&result.Role).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromEmail 批量查找
func (obj *_UsuariosMgr) GetBatchFromEmail(emails []string) (results []*Usuarios, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Usuarios{}).Where("`email` IN (?)", emails).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("role").Where("id = ?", results[i].RoleID).Find(&results[i].Role).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromSenha 通过senha获取内容
func (obj *_UsuariosMgr) GetFromSenha(senha string) (results []*Usuarios, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Usuarios{}).Where("`senha` = ?", senha).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("role").Where("id = ?", results[i].RoleID).Find(&results[i].Role).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromSenha 批量查找
func (obj *_UsuariosMgr) GetBatchFromSenha(senhas []string) (results []*Usuarios, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Usuarios{}).Where("`senha` IN (?)", senhas).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("role").Where("id = ?", results[i].RoleID).Find(&results[i].Role).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDataCadastro 通过data_cadastro获取内容
func (obj *_UsuariosMgr) GetFromDataCadastro(dataCadastro time.Time) (results []*Usuarios, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Usuarios{}).Where("`data_cadastro` = ?", dataCadastro).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("role").Where("id = ?", results[i].RoleID).Find(&results[i].Role).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDataCadastro 批量查找
func (obj *_UsuariosMgr) GetBatchFromDataCadastro(dataCadastros []time.Time) (results []*Usuarios, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Usuarios{}).Where("`data_cadastro` IN (?)", dataCadastros).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("role").Where("id = ?", results[i].RoleID).Find(&results[i].Role).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRoleID 通过role_id获取内容
func (obj *_UsuariosMgr) GetFromRoleID(roleID int64) (results []*Usuarios, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Usuarios{}).Where("`role_id` = ?", roleID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("role").Where("id = ?", results[i].RoleID).Find(&results[i].Role).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRoleID 批量查找
func (obj *_UsuariosMgr) GetBatchFromRoleID(roleIDs []int64) (results []*Usuarios, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Usuarios{}).Where("`role_id` IN (?)", roleIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("role").Where("id = ?", results[i].RoleID).Find(&results[i].Role).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromTelefone 通过telefone获取内容
func (obj *_UsuariosMgr) GetFromTelefone(telefone string) (results []*Usuarios, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Usuarios{}).Where("`telefone` = ?", telefone).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("role").Where("id = ?", results[i].RoleID).Find(&results[i].Role).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromTelefone 批量查找
func (obj *_UsuariosMgr) GetBatchFromTelefone(telefones []string) (results []*Usuarios, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Usuarios{}).Where("`telefone` IN (?)", telefones).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("role").Where("id = ?", results[i].RoleID).Find(&results[i].Role).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromArea 通过area获取内容
func (obj *_UsuariosMgr) GetFromArea(area string) (results []*Usuarios, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Usuarios{}).Where("`area` = ?", area).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("role").Where("id = ?", results[i].RoleID).Find(&results[i].Role).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromArea 批量查找
func (obj *_UsuariosMgr) GetBatchFromArea(areas []string) (results []*Usuarios, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Usuarios{}).Where("`area` IN (?)", areas).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("role").Where("id = ?", results[i].RoleID).Find(&results[i].Role).Error; err != nil { //
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
func (obj *_UsuariosMgr) FetchByPrimaryKey(id int64) (result Usuarios, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Usuarios{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("role").Where("id = ?", result.RoleID).Find(&result.Role).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueByEmailUnique primary or index 获取唯一内容
func (obj *_UsuariosMgr) FetchUniqueByEmailUnique(email string) (result Usuarios, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Usuarios{}).Where("`email` = ?", email).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("role").Where("id = ?", result.RoleID).Find(&result.Role).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByFkUsuariosRoleIDx  获取多个内容
func (obj *_UsuariosMgr) FetchIndexByFkUsuariosRoleIDx(roleID int64) (results []*Usuarios, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Usuarios{}).Where("`role_id` = ?", roleID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("role").Where("id = ?", results[i].RoleID).Find(&results[i].Role).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
