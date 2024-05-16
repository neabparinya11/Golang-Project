package itemhandler

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/neabparinya11/Golang-Project/config"
	"github.com/neabparinya11/Golang-Project/modules/item"
	itemusecase "github.com/neabparinya11/Golang-Project/modules/item/itemUsecase"
	"github.com/neabparinya11/Golang-Project/pkg/request"
	"github.com/neabparinya11/Golang-Project/pkg/response"
)

type (
	ItemHttpHandlerService interface{
		CreateItem(c echo.Context) error
		FindOneItem(c echo.Context) error
		FindManyItems(c echo.Context) error
		EditItem(c echo.Context) error
		EnableOrDisableItem(c echo.Context) error
	}

	ItemHttpHandler struct {
		cfg *config.Config
		itemUsecase itemusecase.ItemUsecaseService
	}
)

func NewItemHttpHandler(cfg *config.Config, itemUsecase itemusecase.ItemUsecaseService) ItemHttpHandlerService{
	return &ItemHttpHandler{cfg: cfg, itemUsecase: itemUsecase}
}

func (h *ItemHttpHandler) CreateItem(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.NewContextWrapper(c)

	req := new(item.CreateItemRequest)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.itemUsecase.CreateItem(ctx, req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusCreated, res)
}

func (h *ItemHttpHandler) FindOneItem(c echo.Context) error {
	ctx := context.Background()

	itemId := strings.TrimPrefix(c.Param("item_id"), "item:")

	res, err := h.itemUsecase.FindOneItem(ctx, itemId)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *ItemHttpHandler) FindManyItems(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.NewContextWrapper(c)

	req := new(item.ItemSearchRequest)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.itemUsecase.FindManyItems(ctx, h.cfg.Paginate.ItemNextPageBaseUrl, req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *ItemHttpHandler) EditItem(c echo.Context) error {
	ctx := context.Background()

	itemId := strings.TrimPrefix(c.Param("item_id"), "item:")

	wrapper := request.NewContextWrapper(c)

	req := new(item.ItemUpdateRequest)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	res, err := h.itemUsecase.EditItem(ctx, itemId, req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *ItemHttpHandler) EnableOrDisableItem(c echo.Context) error {
	ctx := context.Background()

	itemId := strings.TrimPrefix(c.Param("item_id"), "item:")
	
	res, err := h.itemUsecase.EnableOrDisableItem(ctx, itemId)
	if err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	return response.SuccessResponse(c, http.StatusOK, &response.MessageResponse{
		Message: fmt.Sprintf("item_id: %s is successful is activate to %v", itemId, res),
	})
}