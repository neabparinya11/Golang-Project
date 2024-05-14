package inventory

import (
	"github.com/neabparinya11/Golang-Project/modules/item"
	"github.com/neabparinya11/Golang-Project/modules/models"
)

type (
	UpdateInventoryRequest struct {
		PlayerId string `json:"player_id" validate:"required,max=64"`
		ItemId   string `json:"item_id" validate:"required,max=64"`
	}

	ItemInventory struct {
		InventoryId string `json:"inventory_id"`
		*item.ItemShowCase
	}

	PlayerInventory struct{
		PlayerId string `json:"player_id"`
		*models.PaginateResponse
	}
)
