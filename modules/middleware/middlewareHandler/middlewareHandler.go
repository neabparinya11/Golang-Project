package middlewarehandler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/neabparinya11/Golang-Project/config"
	middlewareusecase "github.com/neabparinya11/Golang-Project/modules/middleware/middlewareUsecase"
	"github.com/neabparinya11/Golang-Project/pkg/response"
)

type (
	MiddlewareHandlerService interface{
		JwtAuthorization(next echo.HandlerFunc) echo.HandlerFunc
		RbacAuthorization(next echo.HandlerFunc, expected []int) echo.HandlerFunc
		PlayerIdParamValidation(next echo.HandlerFunc) echo.HandlerFunc
	}

	MiddlewareHandler struct {
		cgf                *config.Config
		middlewareUsecaase middlewareusecase.MiddlewareUsecaseService
	}
)

func NewMiddlewareHandler(cgf *config.Config, middlewareUsercase middlewareusecase.MiddlewareUsecaseService) MiddlewareHandlerService {
	return &MiddlewareHandler{cgf: cgf, middlewareUsecaase: middlewareUsercase}
}

func (h *MiddlewareHandler) JwtAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		accessToken := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")

		nextCtx, err := h.middlewareUsecaase.JwtAuthorization(c, h.cgf, accessToken)
		if err != nil {
			return response.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		}
		return next(nextCtx)
	}
}

func (h *MiddlewareHandler) RbacAuthorization(next echo.HandlerFunc, expected []int) echo.HandlerFunc {
	return func(c echo.Context) error {
		nextCtx, err := h.middlewareUsecaase.RbacAuthorization(c, h.cgf, expected)
		if err != nil {
			return response.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		}
		return next(nextCtx)
	}
}

func (h *MiddlewareHandler) PlayerIdParamValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		nextCtx, err := h.middlewareUsecaase.PlayerIdParamValidation(c)
		if err != nil {
			return response.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		}
		return next(nextCtx)
	}
}