package service

import (
	"fmt"
	authRepository "sistema_artigos_bpk/internal/app/repository/auth"
	eventoRepository "sistema_artigos_bpk/internal/app/repository/evento"
	usuarioRepository "sistema_artigos_bpk/internal/app/repository/usuario"
	"sistema_artigos_bpk/internal/models"
	erros "sistema_artigos_bpk/pkg/errors"
	"sistema_artigos_bpk/pkg/utils"
	"time"

	"github.com/labstack/echo/v4"
)

type EventoServiceInterface interface {

	CreateEvent(data *models.Eventos, jwt string)(*models.Eventos ,erros.CustomError)
	InsertBlobInEvent(id int,  blob []byte)(erros.CustomError)
	AllEvents() ([]eventoRepository.EventoFiltradoDTO, erros.CustomError)
	EditAEvent(data models.Eventos, jwt string)(*models.Eventos, erros.CustomError)
}

type EventoService struct {
	EventoRepository eventoRepository.EventoRepositoryInterface
	JwtRepository    authRepository.JwtRepositoryInterface
	UsuarioRepository usuarioRepository.UsuarioRepository
	Logger           echo.Logger
}

func NewEventoService(
	eventoRepository eventoRepository.EventoRepositoryInterface,
	jwtrepository authRepository.JwtRepositoryInterface,
	UsuarioRepository usuarioRepository.UsuarioRepository,
	log echo.Logger,

) EventoServiceInterface {
	return &EventoService{
		EventoRepository: eventoRepository,
		JwtRepository:    jwtrepository,
		UsuarioRepository: UsuarioRepository,
		Logger:           log,
	}
}


func (s *EventoService) CreateEvent(data *models.Eventos, jwt string)(*models.Eventos, erros.CustomError){

	s.Logger.Print(jwt)
	dataUser, err :=  s.JwtRepository.ParseJWT(jwt)
	if err != nil{
		return &models.Eventos{}, erros.CustomError{
			Message: fmt.Sprint("Erro ao ler JWT %v ", err.Error()) ,
			Error: erros.ErrNaoAutorizado.Error(),
			Orientacao: "Verifique o jwt enviado",
		}
	}
	s.Logger.Print(dataUser.RoleID)
	s.Logger.Print(int(dataUser.RoleID))
	roleData,err := s.UsuarioRepository.QueryByIdRole(int(dataUser.RoleID))
	
	if err != nil{
		return nil, erros.CustomError{
			Message: fmt.Sprint("Erro ao coletar o a role: ", err.Error()) ,
			Error:erros.ErrRecursoInvalido.Error(),
			Orientacao: "Verifique o Jwt, ou Usuario",
		}
	}

	if roleData.Nome != utils.Roles.Moderador{
		return nil, erros.CustomError{
			Message: "Usuario não tem permissao para criar eventos",
			Error: erros.ErrNaoAutorizado.Error(),
			Orientacao: "Verifique as permissoes de usuario",
		}
	}

	if err :=  data.Validate(); err.Message != ""{
		return nil, err
	}


	data.UsuariosID = dataUser.ID
	data.DataInicio = time.Now()
	dataCreated, err := s.EventoRepository.CreateEvent(*data)
	if err != nil {
		return nil, erros.CustomError{
			Message: fmt.Sprint("Erro na criação do evento %v", err) ,
			Error: erros.ErrInterno.Error(),
			Orientacao: "Verfique os dados do evento",
		}
	}

	return dataCreated, erros.CustomError{}
}

func (s *EventoService) InsertBlobInEvent(id int, blob []byte) erros.CustomError {
	err := s.EventoRepository.UpdateBannerById(id, blob)
    if err != nil {
        return erros.CustomError{
            Message:   fmt.Sprintf("Erro ao inserir blob no evento: %v", err),
            Error:     erros.ErrInterno.Error(),
            Orientacao: "Verifique os dados do evento e tente novamente",
        }
    }
    return erros.CustomError{}
}


func (s *EventoService) EditAEvent(data models.Eventos, jwt string) (*models.Eventos, erros.CustomError) {
	// Valida ID
	if err := data.ValidateID(); err.Message != "" {
		return nil, err
	}

	// JWT
	dataUser, err := s.JwtRepository.ParseJWT(jwt)
	if err != nil {
		return nil, erros.CustomError{
			Message:    "Ocorreu um erro no JWT",
			Error:      erros.ErrInterno.Error(),
			Orientacao: "Verifique seu JWT",
		}
	}

	// Busca evento existente
	existing, err := s.EventoRepository.BuscaPorID(int(data.ID))
	if err != nil {
		return nil, erros.CustomError{
			Message:    "Evento não encontrado",
			Error:      err.Error(),
			Orientacao: "Verifique o ID enviado",
		}
	}
	s.Logger.Print(existing)
	// Verifica se é dono
	if existing.UsuariosID != dataUser.ID {
		return nil, erros.CustomError{
			Message:    "Você não pode fazer essa ação",
			Error:      erros.ErrNaoAutorizado.Error(),
			Orientacao: "Verifique seu cargo!",
		}
	}

	// Atualiza somente campos enviados
	if data.Nome != "" {
		existing.Nome = data.Nome
	}
	if data.Descricao != "" {
		existing.Descricao = data.Descricao
	}
	if !data.DataFim.IsZero() {
		existing.DataFim = data.DataFim
	}

	// Salva
	dataMod, err := s.EventoRepository.EditEvent(existing)
	if err != nil {
		return nil, erros.CustomError{
			Message:    fmt.Sprint("Erro ao modificar: ", err.Error()),
			Error:      erros.ErrInterno.Error(),
			Orientacao: "Verifique os dados",
		}
	}

	return dataMod, erros.CustomError{}
}
func (s *EventoService) AllEvents() ([]eventoRepository.EventoFiltradoDTO, erros.CustomError){
	events, err := s.EventoRepository.Events()
	if err != nil{
		return nil, erros.CustomError{
			Message: fmt.Sprint("Houve um erro no retorno de eventos ", err) ,
			Error: erros.ErrInterno.Error(),
			Orientacao: "Verifique com o Dev",
		}
	}
	return events, erros.CustomError{}

}