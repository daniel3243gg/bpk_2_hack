package erros

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/labstack/echo/v4"
    "github.com/stretchr/testify/assert"
)

func TestHandleError_NotFound(t *testing.T) {
    e := echo.New()
    req := httptest.NewRequest(http.MethodGet, "/", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    errorCustom := CustomError{
        Message:    "Recurso não encontrado",
        Error:      ErrNaoEncontrado.Error(),
        Orientacao: "Verifique o ID enviado",
    }

    err := HandleError(c, errorCustom)
    assert.NoError(t, err)
    assert.Equal(t, http.StatusNotFound, rec.Code)

    var resp CustomError
    err = json.Unmarshal(rec.Body.Bytes(), &resp)
    assert.NoError(t, err)
    assert.Equal(t, errorCustom, resp)
}

func TestHandleError_Unauthorized(t *testing.T) {
    e := echo.New()
    req := httptest.NewRequest(http.MethodGet, "/", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    errorCustom := CustomError{
        Message:    "Não autorizado",
        Error:      ErrNaoAutorizado.Error(),
        Orientacao: "Faça login para continuar",
    }

    err := HandleError(c, errorCustom)
    assert.NoError(t, err)
    assert.Equal(t, http.StatusUnauthorized, rec.Code)

    var resp CustomError
    err = json.Unmarshal(rec.Body.Bytes(), &resp)
    assert.NoError(t, err)
    assert.Equal(t, errorCustom, resp)
}

func TestHandleError_InternalError(t *testing.T) {
    e := echo.New()
    req := httptest.NewRequest(http.MethodGet, "/", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    errorCustom := CustomError{
        Message:    "Erro interno no servidor",
        Error:      ErrInterno.Error(),
        Orientacao: "Tente novamente mais tarde",
    }

    err := HandleError(c, errorCustom)
    assert.NoError(t, err)
    assert.Equal(t, http.StatusInternalServerError, rec.Code)

    var resp CustomError
    err = json.Unmarshal(rec.Body.Bytes(), &resp)
    assert.NoError(t, err)
    assert.Equal(t, errorCustom, resp)
}

func TestHandleError_BadRequest(t *testing.T) {
    e := echo.New()
    req := httptest.NewRequest(http.MethodGet, "/", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    errorCustom := CustomError{
        Message:    "Requisição inválida",
        Error:      ErrRecursoInvalido.Error(),
        Orientacao: "Verifique os dados enviados",
    }

    err := HandleError(c, errorCustom)
    assert.NoError(t, err)
    assert.Equal(t, http.StatusBadRequest, rec.Code)

    var resp CustomError
    err = json.Unmarshal(rec.Body.Bytes(), &resp)
    assert.NoError(t, err)
    assert.Equal(t, errorCustom, resp)
}

func TestHandleError_Default(t *testing.T) {
    e := echo.New()
    req := httptest.NewRequest(http.MethodGet, "/", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    errorCustom := CustomError{
        Message:    "Erro desconhecido",
        Error:      "Erro não mapeado",
        Orientacao: "Entre em contato com o suporte",
    }

    err := HandleError(c, errorCustom)
    assert.NoError(t, err)
    assert.Equal(t, http.StatusInternalServerError, rec.Code)

    var resp CustomError
    err = json.Unmarshal(rec.Body.Bytes(), &resp)
    assert.NoError(t, err)
    assert.Equal(t, errorCustom, resp)
}