package authRepository

import (
	"errors"
	"sistema_artigos_bpk/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("asasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

type JwtRepositoryInterface interface {
	GenerateJWT(email, telefone, nome string, id, role_id int) (string, error)
	ParseJWT(tokenString string) (models.Usuarios, error)
}

type JwtRepository struct {
	jwtKey []byte
}

func NewJwtRepository() *JwtRepository {
	return &JwtRepository{
		jwtKey: jwtKey,
	}
}

func (r *JwtRepository) GenerateJWT(email, telefone, nome string, id, role_id int) (string, error) {
	claims := jwt.MapClaims{
		"usuario_id": id,
		"role_id":    role_id,
		"email":      email,
		"telefone":   telefone,
		"exp":        time.Now().Add(time.Hour * 24).Unix(), // expira em 24h
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(r.jwtKey)
}

func (r *JwtRepository) ParseJWT(tokenString string) (models.Usuarios, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("asasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"), nil
	})

	if err != nil || !token.Valid {
		return models.Usuarios{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return models.Usuarios{}, errors.New("invalid token claims")
	}

	empresaID, ok := claims["usuario_id"].(float64)
	if !ok {
		return models.Usuarios{}, errors.New("invalid empresa_id type")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return models.Usuarios{}, errors.New("invalid email type")
	}


	role_id, ok := claims["role_id"].(float64)
	if !ok {
		return models.Usuarios{}, errors.New("invalid role type")
	}

	empresa := models.Usuarios{
		ID:     int64(empresaID),
		Email:  email,
		RoleID: int64(role_id),
	}

	return empresa, nil
}
