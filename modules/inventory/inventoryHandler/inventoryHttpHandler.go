package inventoryhandler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/neabparinya11/Golang-Project/config"
	"github.com/neabparinya11/Golang-Project/modules/inventory"
	inventoryusecase "github.com/neabparinya11/Golang-Project/modules/inventory/inventoryUsecase"
	"github.com/neabparinya11/Golang-Project/pkg/request"
	"github.com/neabparinya11/Golang-Project/pkg/response"
)

type (
	InventoryHttpHandlerService interface{
		FindPlayerItems(c echo.Context) error
	}

	InventoryHttpHandler struct {
		cfg *config.Config
		inventoryUsecase inventoryusecase.InventoryUsecaseService
	}
)

func NewInventoryHttpHandler(cfg *config.Config, inventoryUsecase inventoryusecase.InventoryUsecaseService) InventoryHttpHandlerService {
	return &InventoryHttpHandler{cfg: cfg, inventoryUsecase: inventoryUsecase}
}

func (h *InventoryHttpHandler) FindPlayerItems(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.NewContextWrapper(c)

	playerId := c.Param("player_id")

	req := new(inventory.InventorySearchRequest)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.inventoryUsecase.FindPlayerItems(ctx, h.cfg.Paginate.ItemNextPageBaseUrl, playerId, req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}