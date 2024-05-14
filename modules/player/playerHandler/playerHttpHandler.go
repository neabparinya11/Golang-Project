package playerhandler

import (
	"context"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/neabparinya11/Golang-Project/config"
	"github.com/neabparinya11/Golang-Project/modules/player"
	playerusecase "github.com/neabparinya11/Golang-Project/modules/player/playerUsecase"
	"github.com/neabparinya11/Golang-Project/pkg/request"
	"github.com/neabparinya11/Golang-Project/pkg/response"
)

type (
	PlayerHttpHandlerService interface{
		CreatePlayer(c echo.Context) error
		FindOnePlayerProfile(c echo.Context) error
		AddPlayerMoney(c echo.Context) error
		GetPlayerSavingAccout(c echo.Context) error
	}

	PlayerHttpHandler struct {
		cfg           *config.Config
		playerUsecase playerusecase.PlayerUsecaseService
	}
)

func NewPlayerHttpHandler(cfg *config.Config, playerUsecase playerusecase.PlayerUsecaseService) PlayerHttpHandlerService {
	return &PlayerHttpHandler{playerUsecase: playerUsecase, cfg: cfg}
}

func (h *PlayerHttpHandler) CreatePlayer(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.NewContextWrapper(c)

	req := new(player.CreatePlayerRequest)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.playerUsecase.CreatePlayer(ctx, req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}

func (h *PlayerHttpHandler) FindOnePlayerProfile(c echo.Context) error{
	ctx := context.Background()

	playerId := strings.TrimPrefix(c.Param("player_id"), "player:")

	res, err := h.playerUsecase.FindOnePlayerProfile(ctx, playerId)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusAccepted, res)
}

func (h *PlayerHttpHandler) AddPlayerMoney(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.NewContextWrapper(c)

	req := new(player.CreatePlayerTransactionRequest)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.playerUsecase.AddPlayerMoney(ctx, req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusAccepted, res)
}

func (h *PlayerHttpHandler) GetPlayerSavingAccout(c echo.Context) error {
	ctx := context.Background()

	playerId := c.Param("player_id")

	res, err := h.playerUsecase.GetPlayerSavingAccout(ctx, playerId)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusAccepted, res)
}