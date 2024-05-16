package inventoryusecase

import (
	"context"
	"fmt"

	"github.com/neabparinya11/Golang-Project/modules/inventory"
	inventoryrepository "github.com/neabparinya11/Golang-Project/modules/inventory/inventoryRepository"
	"github.com/neabparinya11/Golang-Project/modules/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	InventoryUsecaseService interface{
		FindPlayerItems(pctx context.Context, basePaginateUrl, playerId string, req *inventory.InventorySearchRequest) (*models.PaginateResponse, error)
	}

	InventoryUsecase struct {
		inventoryRepository inventoryrepository.InventoryRepositoryService
	}
)

func NewInventoryRepository(inventoryRepository inventoryrepository.InventoryRepositoryService) InventoryUsecaseService{
	return &InventoryUsecase{inventoryRepository: inventoryRepository}
}

func (u *InventoryUsecase) FindPlayerItems(pctx context.Context, basePaginateUrl, playerId string, req *inventory.InventorySearchRequest) (*models.PaginateResponse, error) {
	findPlayerItemOptions := make([]*options.FindOptions, 0)
	findPlayerItemOptions = append(findPlayerItemOptions, options.Find().SetSort(bson.D{{"_id", 1}}))
	findPlayerItemOptions = append(findPlayerItemOptions, options.Find().SetLimit(int64(req.Limit)))

	resultFindPlayerItem, err := u.inventoryRepository.FindPlayerItems(pctx, playerId, findPlayerItemOptions)
	if err != nil {
		return nil, err
	}

	countFindPlayerItem, err := u.inventoryRepository.CountPlayerItems(pctx, playerId)
	if err != nil {
		return nil, err
	}

	if len(resultFindPlayerItem) == 0 {
		return &models.PaginateResponse{
			Data: make([]*inventory.ItemInventory, 0),
			Total: countFindPlayerItem,
			Limit: req.Limit,
			First: models.FirstPaginate{
				Href: fmt.Sprintf("%s?limit=%d", basePaginateUrl, req.Limit),
			},
			Next: models.NextPaginate{
				Start: "",
				Href: "",
			},
		}, nil
	}

	return &models.PaginateResponse{
		Data: resultFindPlayerItem,
		Total: countFindPlayerItem,
		Limit: req.Limit,
		First: models.FirstPaginate{
			Href: fmt.Sprintf("%s?limit=%d", basePaginateUrl, req.Limit),
		},
		Next: models.NextPaginate{
			Start: resultFindPlayerItem[len(resultFindPlayerItem) -1].ItemId,
			Href: fmt.Sprintf("%s?limit=%d&start=%s", basePaginateUrl, req.Limit, resultFindPlayerItem[len(resultFindPlayerItem) -1].ItemId),
		},
	}, nil
}