package itemhandler

import (
	"github.com/neabparinya11/Golang-Project/config"
	itemusecase "github.com/neabparinya11/Golang-Project/modules/item/itemUsecase"
)

type (
	ItemHttpHandlerService interface{}

	ItemHttpHandler struct {
		cfg *config.Config
		itemUsecase itemusecase.ItemUsecaseService
	}
)

func NewItemHttpHandler(cfg *config.Config, itemUsecase itemusecase.ItemUsecaseService) ItemHttpHandlerService{
	return &ItemHttpHandler{cfg: cfg, itemUsecase: itemUsecase}
}