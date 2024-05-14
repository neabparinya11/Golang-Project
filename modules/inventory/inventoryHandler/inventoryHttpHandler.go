package inventoryhandler

import (
	"github.com/neabparinya11/Golang-Project/config"
	inventoryusecase "github.com/neabparinya11/Golang-Project/modules/inventory/inventoryUsecase"
)

type (
	InventoryHttpHandlerService interface{}

	InventoryHttpHandler struct {
		cfg *config.Config
		inventoryUsecase inventoryusecase.InventoryUsecaseService
	}
)

func NewInventoryHttpHandler(cfg *config.Config, inventoryUsecase inventoryusecase.InventoryUsecaseService) InventoryHttpHandlerService {
	return &InventoryHttpHandler{cfg: cfg, inventoryUsecase: inventoryUsecase}
}