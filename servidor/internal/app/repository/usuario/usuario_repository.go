package usuarioRepository

import (
	"sistema_artigos_bpk/internal/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UsuarioRepositoryInterface interface {
	QueryByIdRole(id int) (*models.Role, error)
}

type UsuarioRepository struct {
	db     *gorm.DB
	Logger echo.Logger
}


func NewUsuarioRepository(db *gorm.DB, log echo.Logger) UsuarioRepository {
	return UsuarioRepository{
		db:     db,
		Logger: log,
	}
}


func(r *UsuarioRepository) QueryByIdRole(id int)( *models.Role, error){
	var role models.Role

	if err := r.db.Where("id = ?", id).First(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}