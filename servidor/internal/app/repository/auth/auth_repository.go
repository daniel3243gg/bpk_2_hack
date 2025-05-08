package authRepository

import (
	"sistema_artigos_bpk/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthRepositoryInterface interface {
	GetAuthByEmailAndPassword(email string, senha string) (*models.Usuarios, error)
	CreateAuth(data models.Usuarios ) (*models.Usuarios ,error)
	LoginAuth(data models.Usuarios)(*models.Usuarios, error)
	LocateAuth(data models.Usuarios)([]models.Usuarios, error)
}


type AuthRepository struct {
	db 				*gorm.DB
	Logger         echo.Logger
}


func NewAuthRepository( db *gorm.DB,log echo.Logger) *AuthRepository {
	return &AuthRepository{
		db: db,
		Logger: log,
	}
}


func (r *AuthRepository) GetAuthByEmailAndPassword(email string, senha string) (*models.Usuarios, error) {
	var auth models.Usuarios

	err := r.db.Where("email = ? AND senha = ?", email, senha).First(&auth).Error
	if err != nil {
		return nil, err
	}

	return &auth, nil

}

func (r *AuthRepository) CreateAuth(data models.Usuarios) (*models.Usuarios, error) {
	if err := r.db.Create(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}


func (r *AuthRepository) LoginAuth(data models.Usuarios) (*models.Usuarios, error) {
	var user models.Usuarios

	err := r.db.Where("email = ? AND senha = ?", data.Email, data.Senha).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepository) LocateAuth(data models.Usuarios) ([]models.Usuarios, error) {
	var usuarios []models.Usuarios

	err := r.db.Where("email = ? OR telefone = ? ", data.Email, data.Telefone).Find(&usuarios).Error
	if err != nil {
		return nil, err
	}

	return usuarios, nil
}