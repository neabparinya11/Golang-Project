package middlewarehandler

import (
	"github.com/neabparinya11/Golang-Project/config"
	middlewareusecase "github.com/neabparinya11/Golang-Project/modules/middleware/middlewareUsecase"
)

type (
	MiddlewareHandlerService interface{}

	MiddlewareHandler struct {
		cgf                *config.Config
		middlewareUsecaase middlewareusecase.MiddlewareUsecaseService
	}
)

func NewMiddlewareHandler(cgf *config.Config, middlewareUsercase middlewareusecase.MiddlewareUsecaseService) MiddlewareHandlerService {
	return &MiddlewareHandler{cgf: cgf, middlewareUsecaase: middlewareUsercase}
}
