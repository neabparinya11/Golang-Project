package authhandler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/neabparinya11/Golang-Project/config"
	"github.com/neabparinya11/Golang-Project/modules/auth"
	authusecase "github.com/neabparinya11/Golang-Project/modules/auth/authUsecase"
	"github.com/neabparinya11/Golang-Project/pkg/request"
	"github.com/neabparinya11/Golang-Project/pkg/response"
)

type (
	AuthHttpHandlerService interface{
		Login(c echo.Context) error
		RefreshToken(c echo.Context) error
		Logout(c echo.Context) error
	}

	AuthHttpHandler struct {
		cfg *config.Config
		authUsecase authusecase.AuthUsecaseService
	}
)

func NewAuthHttpHandler(cfg *config.Config, authUsecase authusecase.AuthUsecaseService) AuthHttpHandlerService{
	return &AuthHttpHandler{authUsecase: authUsecase, cfg: cfg}
}

func (h *AuthHttpHandler) Login(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.NewContextWrapper(c)
	req := new(auth.PlayerLoginRequest)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.authUsecase.Login(ctx, h.cfg, req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *AuthHttpHandler) RefreshToken(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.NewContextWrapper(c)

	req := new(auth.PlayerRefreshToken)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.authUsecase.RefreshToken(ctx, h.cfg, req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *AuthHttpHandler) Logout(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.NewContextWrapper(c)

	req := new(auth.LogoutRequest)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.authUsecase.Logout(ctx, req.CredentialId)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, &response.MessageResponse{
		Message: fmt.Sprintf("Delete count: %d", res),
	})
}