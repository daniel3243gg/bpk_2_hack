package response

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/labstack/echo/v4"
    "github.com/stretchr/testify/assert"
)

func TestHandleResponse_Success(t *testing.T) {
    e := echo.New()
    req := httptest.NewRequest(http.MethodGet, "/", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    data := map[string]interface{}{"key": "value"}
    message := "Operação realizada com sucesso"

    err := HandleResponse(c, data, message)
    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, rec.Code)

    var resp CustomResponse
    err = json.Unmarshal(rec.Body.Bytes(), &resp)
    assert.NoError(t, err)
    assert.Equal(t, data, resp.Data)
    assert.Equal(t, message, resp.Message)
}

func TestHandleResponse_EmptyData(t *testing.T) {
    e := echo.New()
    req := httptest.NewRequest(http.MethodGet, "/", nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    message := "Nenhum dado encontrado"

    err := HandleResponse(c, nil, message)
    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, rec.Code)

    var resp CustomResponse
    err = json.Unmarshal(rec.Body.Bytes(), &resp)
    assert.NoError(t, err)
    assert.Nil(t, resp.Data)
    assert.Equal(t, message, resp.Message)
}