package erros

import "errors"

var (
	ErrNaoEncontrado     = errors.New("Recurso não encontrado")
	ErrNaoAutorizado = errors.New("Não autorizado")
	ErrInterno     = errors.New("Erro interno do servidor")
    ErrRecursoInvalido     = errors.New("Recurso inválido")
)


type CustomError struct {
    Message string `json:"message"`
    Error string `json:"error"`
    Orientacao string `json:"orientacao"`
}
