package handler

import (
	"io"
	"sistema_artigos_bpk/internal/app/repository/evento"
	"sistema_artigos_bpk/internal/app/service"
	"sistema_artigos_bpk/internal/models"
	erros "sistema_artigos_bpk/pkg/errors"
	"sistema_artigos_bpk/pkg/response"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type EventoHandler struct {
	repositoryEvento eventoRepository.EventoRepositoryInterface
	serviceEvento   service.EventoServiceInterface
}

func NewEventoHandler(repositoryEvento eventoRepository.EventoRepositoryInterface, serviceEvento service.EventoServiceInterface) *EventoHandler {
	return &EventoHandler{
		repositoryEvento: repositoryEvento,
		serviceEvento:    serviceEvento,
	}
}


func (h *EventoHandler) CreateEvento(c echo.Context) error {
    var eventoCreated models.Eventos

    if err := c.Bind(&eventoCreated); err != nil {
        return erros.HandleError(c, erros.CustomError{
            Message:    "Erro ao processar os dados",
            Error:      err.Error(),
            Orientacao: "Verifique o formato do JSON enviado",
        })
    }
	token := c.Request().Header.Get("Authorization")
	
    _, err := h.serviceEvento.CreateEvent(&eventoCreated ,token ) 
    if err.Message != "" {
        return erros.HandleError(c, err)
    }

    return response.HandleResponse(c, eventoCreated.Nome, "Evento criado criado.")
}

func (h *EventoHandler)AllEvents(c echo.Context) error{
	events, err := h.serviceEvento.AllEvents()
	if err.Message !=""{
		return erros.HandleError(c, err)
	}
	return response.HandleResponse(c, events, "Eventos retornados")

}


func (h *EventoHandler) EditEvento(c echo.Context)(error){
	var evento models.Eventos
    if err := c.Bind(&evento); err != nil {
        return erros.HandleError(c, erros.CustomError{
            Message:    "Erro ao processar os dados",
            Error:      err.Error(),
            Orientacao: "Verifique o formato do JSON enviado",
        })

    }
	c.Logger().Print(evento)
	token := c.Request().Header.Get("Authorization")

	eventoAtualizado,er := h.serviceEvento.EditAEvent(evento, token)
	if er.Message != "" {
        return erros.HandleError(c, er)
    }

	return response.HandleResponse(c,eventoAtualizado,"Editado com sucesso" )
}

func (h *EventoHandler) UploadPDF(c echo.Context) error {
    eventoIDStr := c.Param("id")
    eventoID, err := strconv.Atoi(eventoIDStr)
    if err != nil {
        return erros.HandleError(c, erros.CustomError{
            Message:    "ID inválido",
            Error:      err.Error(),
            Orientacao: "Verifique o ID enviado na URL",
        })
    }

    file, err := c.FormFile("banner")
    if err != nil {
        return erros.HandleError(c, erros.CustomError{
            Message:    "Arquivo não enviado",
            Error:      err.Error(),
            Orientacao: "Verifique o arquivo",
        })
    }

	if !strings.HasSuffix(strings.ToLower(file.Filename), ".jpg") &&
		!strings.HasSuffix(strings.ToLower(file.Filename), ".jpeg") &&
		!strings.HasSuffix(strings.ToLower(file.Filename), ".png") &&
		!strings.HasSuffix(strings.ToLower(file.Filename), ".gif") {
			return erros.HandleError(c, erros.CustomError{
				Message:    "Somente imagens são permitidas",
				Error:      "Formato inválido",
				Orientacao: "Envie um arquivo com extensão .jpg, .jpeg, .png ou .gif",
			})
	}

    src, err := file.Open()
    if err != nil {
        return erros.HandleError(c, erros.CustomError{
            Message:    "Erro ao abrir o arquivo",
            Error:      err.Error(),
            Orientacao: "Tente novamente",
        })
    }
    defer src.Close()

    pdfBytes, err := io.ReadAll(src)
    if err != nil {
        return erros.HandleError(c, erros.CustomError{
            Message:    "Erro ao ler o arquivo",
            Error:      err.Error(),
            Orientacao: "Tente novamente",
        })
    }


    if customErr := h.serviceEvento.InsertBlobInEvent(eventoID, pdfBytes); customErr.Message != "" {
        return erros.HandleError(c, customErr)
    }
	return response.HandleResponse(c, nil, "Sucesso na inserção")

    
}


