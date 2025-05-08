package response
import "github.com/labstack/echo/v4"

func HandleResponse(c echo.Context, data interface{}, message string) error {

	return c.JSON(200, CustomResponse{
		Data:    data,
		Message: message,
	})
}