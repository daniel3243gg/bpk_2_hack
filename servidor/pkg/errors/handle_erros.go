package erros

import "github.com/labstack/echo/v4"





func HandleError(c echo.Context, errorCustom CustomError) error {
	
	switch errorCustom.Error {
	case ErrNaoEncontrado.Error():
		err := c.JSON(404, errorCustom)
		if err != nil {
			return err
		}
	case ErrNaoAutorizado.Error():
		err := c.JSON(401, errorCustom)
		if err != nil {
			return err
		}
	case ErrInterno.Error():
		err := c.JSON(500, errorCustom)
		if err != nil {
			return err
		}
	case ErrRecursoInvalido.Error():
		err := c.JSON(400, errorCustom)
		if err != nil {
			return err
		}
	default:
		err := c.JSON(500, errorCustom)
		if err != nil {
			return err
		}
	}
	return nil
}