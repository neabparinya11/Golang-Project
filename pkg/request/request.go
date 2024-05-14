package request

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	ContextWrapperService interface {
		Bind(data any) error
	}

	ContextWrapper struct {
		Context   echo.Context
		Validator *validator.Validate
	}
)

func NewContextWrapper(c echo.Context) ContextWrapperService{
	return &ContextWrapper{
		Context: c,
		Validator: validator.New(),
	}
}

func (cont *ContextWrapper) Bind(data any) error{
	if err := cont.Context.Bind(data); err != nil {
		log.Fatalf("Error: Deserialize data failed: %s ", err.Error())
	}

	if err := cont.Validator.Struct(data); err != nil {
		log.Fatalf("Error: Validate struct data failed: %s", err.Error())
	}
	return nil
}
