package handler

import (
	authRepository "sistema_artigos_bpk/internal/app/repository/auth"
	"sistema_artigos_bpk/internal/app/service"
	"sistema_artigos_bpk/internal/models"
	erros "sistema_artigos_bpk/pkg/errors"
	"sistema_artigos_bpk/pkg/response"
	"github.com/labstack/echo/v4"
)



type AuthHandler struct {
    repositoryAuth authRepository.AuthRepositoryInterface
    repositoryJwt  authRepository.JwtRepositoryInterface
    serviceAuth    service.AuthServiceInterface
}

func NewAuthHandler(repositoryAuth authRepository.AuthRepositoryInterface, repositoryJwt authRepository.JwtRepositoryInterface, serviceAuth service.AuthServiceInterface) *AuthHandler {
    return &AuthHandler{
        repositoryAuth: repositoryAuth,
        repositoryJwt:  repositoryJwt,
        serviceAuth:    serviceAuth,
    }
}

func (h *AuthHandler) CreateAuth(c echo.Context) error {
    // Decodificando o JSON enviado no corpo da requisição
    var empresaCreated models.Usuarios
    if err := c.Bind(&empresaCreated); err != nil {
        return erros.HandleError(c, erros.CustomError{
            Message:    "Erro ao processar os dados",
            Error:      err.Error(),
            Orientacao: "Verifique o formato do JSON enviado",
        })
    }

    // Chamando o serviço para criar o estabelecimento
    jwt, err := h.serviceAuth.CreateUserWithToken(empresaCreated)
    if err.Message != "" {
        return erros.HandleError(c, err)
    }
    c.Logger().Printf("entro aqui222")
    return response.HandleResponse(c, jwt, "Usuario criado.")
}


func (h *AuthHandler) Login(c echo.Context) error {
	var empresa models.Usuarios
    if err := c.Bind(&empresa); err != nil {
        return erros.HandleError(c, erros.CustomError{
            Message:    "Erro ao processar os dados",
            Error:      err.Error(),
            Orientacao: "Verifique o formato do JSON enviado",
        })
    }

    // Chamando o serviço para criar o estabelecimento
    jwt, err := h.serviceAuth.LoginUserWithToken(empresa)
    if err.Message != "" {
        return erros.HandleError(c, err)
    }
    
    return response.HandleResponse(c, jwt, "usuario Localizado.")
}
