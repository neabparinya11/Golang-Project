package item

import "github.com/neabparinya11/Golang-Project/modules/models"

type (
	CreateItemRequest struct {
		Title    string  `json:"title" validate:"required,max=64"`
		Price    float64 `json:"price" validate:"required"`
		Damage   int     `json:"damage" validate:"required"`
		ImageUrl string  `json:"image_url" validate:"required,max=255"`
	}

	ItemShowCase struct {
		ItemId   string  `json:"item_id"`
		Title    string  `json:"title"`
		Price    float64 `json:"price"`
		Damage   int     `json:"damage"`
		ImageUrl string  `json:"image_url"`
	}

	ItemSearchRequest struct {
		Title string `json:"title"`
		models.PaginateRequest
	}

	ItemUpdateRequest struct{
		Title    string  `json:"title" validate:"required,max=64"`
		Price    float64 `json:"price" validate:"required"`
		Damage   int     `json:"damage" validate:"required"`
		ImageUrl string  `json:"image_url" validate:"required,max=255"`
	}

	EnableOrDisableItemRequest struct{
		UsageStatus bool `json:"status"`
	}
)
