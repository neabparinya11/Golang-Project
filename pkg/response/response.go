package response

import "github.com/labstack/echo/v4"

type (
	MessageResponse struct {
		Message string `json:"message"`
	}
)

func ErrorResponse(c echo.Context, statuscode int, message string) error{
	return c.JSON(statuscode, &MessageResponse{Message: message})
}

func SuccessResponse(c echo.Context, statuscode int, data any) error{
	return c.JSON(statuscode, data)
}