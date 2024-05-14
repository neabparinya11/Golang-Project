package inventoryhandler

import (
	"github.com/neabparinya11/Golang-Project/config"
	inventoryusecase "github.com/neabparinya11/Golang-Project/modules/inventory/inventoryUsecase"
)

type (
	InventoryQueueHandlerService interface{}

	InventoryQueueHandler struct {
		cfg              *config.Config
		inventoryUsecase inventoryusecase.InventoryUsecaseService
	}
)

func NewInventoryQueueHandler(cfg *config.Config, inventoryUsecase inventoryusecase.InventoryUsecaseService) InventoryQueueHandlerService{
	return &InventoryQueueHandler{cfg: cfg, inventoryUsecase: inventoryUsecase}
}