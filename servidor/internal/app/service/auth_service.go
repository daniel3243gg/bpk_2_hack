package service

import (
	"fmt"
	authRepository "sistema_artigos_bpk/internal/app/repository/auth"
	"sistema_artigos_bpk/internal/models"
	"sistema_artigos_bpk/pkg/errors"
	"sistema_artigos_bpk/pkg/utils"
	"time"

	"github.com/labstack/echo/v4"
)

type AuthServiceInterface interface {
    CreateUserWithToken(empresa models.Usuarios) (string,  erros.CustomError)
    PrepareUsuarioData(data *models.Usuarios) // <- aqui!
    HashPassword(password string)string
    LoginUserWithToken(empresa models.Usuarios)(string, erros.CustomError)
    ExistsAuths(data *models.Usuarios)(bool, error)
}

type AuthService struct {
	AuthRepository authRepository.AuthRepositoryInterface
	JwtRepository authRepository.JwtRepositoryInterface
    Logger         echo.Logger
}

func NewAuthService(authRepository authRepository.AuthRepositoryInterface, jwtrepository authRepository.JwtRepositoryInterface, log echo.Logger) AuthServiceInterface {
	return &AuthService{
		AuthRepository: authRepository,
		JwtRepository:  jwtrepository, 
        Logger: log,
	}
}


func (s *AuthService) CreateUserWithToken(data models.Usuarios) (string, erros.CustomError) {
    // Validação de dados
	if err := data.Validate(); err.Message != "" {
        return "", err
    }

    existe, err := s.ExistsAuths(&data)
    if err != nil {
        return "", handleRepositoryError(err, "Erro ao localizar Usuarios")
    }

    if existe{
        return "", erros.CustomError{
            Message: "Ja existe usuarios com esse cadastro.",
            Error: erros.ErrNaoAutorizado.Error(),
            Orientacao: "Informe outros dados",
        } 
    }
    // Preparação dos dados
	s.PrepareUsuarioData(&data)

    // Criação do usuário
    usuario, err := s.AuthRepository.CreateAuth(data)
    if err != nil || usuario.Email == ""{
        return "", handleRepositoryError(err, "Erro ao criar usuário")
    }

    // Geração do JWT
    jwt, err := s.generateJWT(usuario.Email, usuario.Telefone, usuario.Nome, int(usuario.ID), int(data.RoleID))
    if err != nil {
        return "", handleRepositoryError(err, "Erro na geração do JWT")
    }

    return jwt, erros.CustomError{}
}



func (s *AuthService) LoginUserWithToken(data models.Usuarios)(string,erros.CustomError){
    if err := data.ValidateEmailAndSenha(); err.Message != "" {
        return "", err
    }

    s.PrepareUsuarioData(&data)

    usuario, err := s.AuthRepository.LoginAuth(data)
    if err != nil{
        if err.Error() =="record not found"{
            return "",erros.CustomError{
                Message: "Não localizado usuario",
                Error: erros.ErrNaoAutorizado.Error(),
                Orientacao: "verifique os dados",
            }
        }
        return "",handleRepositoryError(err, "Erro na localização da Usuario")
    }

   
    
    jwt, err := s.generateJWT(usuario.Email, usuario.Telefone, usuario.Nome, int(usuario.ID), int(usuario.RoleID))
    if err != nil{
        return "", handleRepositoryError(err, "Erro na geração do jwt")
    }

    return jwt, erros.CustomError{}
}



func (s *AuthService) PrepareUsuarioData(data *models.Usuarios) {
    data.DataCadastro = time.Now()
    data.Senha = s.HashPassword(data.Senha)
}

func (s *AuthService) ExistsAuths(data *models.Usuarios) (bool, error){
    empresas, err := s.AuthRepository.LocateAuth(*data)
    if err != nil{
        return false, err
    }

    if len(empresas) > 0{
        return true, nil
    }
    return false,nil
}


func (s *AuthService) HashPassword(password string) string {
    return utils.HashString(password)
}

func (s *AuthService) generateJWT(email, cnpjCpf, nome string, id, role_id int) (string, error) {
    jwt, err := s.JwtRepository.GenerateJWT(email, cnpjCpf, nome, id,role_id)
    if err != nil {
        return "", fmt.Errorf("Erro ao gerar token JWT: %v", err)
    }
    return jwt, nil
}

func handleRepositoryError(err error, message string) erros.CustomError {
    return erros.CustomError{
        Message:   fmt.Sprintf("%s: %v", message, err),
        Error:     erros.ErrInterno.Error(),
        Orientacao: "Verifique os dados informados",
    }
}

