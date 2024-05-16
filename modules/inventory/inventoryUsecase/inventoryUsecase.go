package inventoryusecase

import (
	"context"
	"fmt"

	"github.com/neabparinya11/Golang-Project/config"
	"github.com/neabparinya11/Golang-Project/modules/inventory"
	inventoryrepository "github.com/neabparinya11/Golang-Project/modules/inventory/inventoryRepository"
	"github.com/neabparinya11/Golang-Project/modules/item"
	"github.com/neabparinya11/Golang-Project/modules/models"
	"github.com/neabparinya11/Golang-Project/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	itemPb "github.com/neabparinya11/Golang-Project/modules/item/itemPb"
)

type (
	InventoryUsecaseService interface{
		FindPlayerItems(pctx context.Context, cfg *config.Config, playerId string, req *inventory.InventorySearchRequest) (*models.PaginateResponse, error)
	}

	InventoryUsecase struct {
		inventoryRepository inventoryrepository.InventoryRepositoryService
	}
)

func NewInventoryRepository(inventoryRepository inventoryrepository.InventoryRepositoryService) InventoryUsecaseService{
	return &InventoryUsecase{inventoryRepository: inventoryRepository}
}

func (u *InventoryUsecase) FindPlayerItems(pctx context.Context, cfg *config.Config, playerId string, req *inventory.InventorySearchRequest) (*models.PaginateResponse, error) {
	findPlayerItemFilter := bson.D{}

	if req.Start != "" {
		findPlayerItemFilter = append(findPlayerItemFilter, bson.E{"_id", bson.D{{"$gt", utils.ConvertToObjectId(req.Start)}}})
	}
	findPlayerItemFilter = append(findPlayerItemFilter, bson.E{"player_id", playerId})

	findPlayerItemOptions := make([]*options.FindOptions, 0)
	findPlayerItemOptions = append(findPlayerItemOptions, options.Find().SetSort(bson.D{{"_id", 1}}))
	findPlayerItemOptions = append(findPlayerItemOptions, options.Find().SetLimit(int64(req.Limit)))

	resultFindPlayerItem, err := u.inventoryRepository.FindPlayerItems(pctx, findPlayerItemFilter, findPlayerItemOptions)
	if err != nil {
		return nil, err
	}

	itemData, err := u.inventoryRepository.FindItemInIds(pctx, cfg.Grpc.ItemUrl, &itemPb.FindItemInIdsRequest{
		Ids: func() []string {
			itemIds := make([]string, 0)
			for _, v := range resultFindPlayerItem {
				itemIds = append(itemIds, v.ItemId)
			}
			return itemIds
		}(),
	})

	itemMaps := make(map[string]*item.ItemShowCase)
	for _, v := range itemData.Items {
		itemMaps[v.Id] = &item.ItemShowCase{
			ItemId: v.Id,
			Title: v.Title,
			Price: v.Price,
			ImageUrl: v.ImageUrl,
			Damage: int(v.Damage),
		}
	}

	results := make([]*inventory.ItemInventory, 0)
	for _, v := range resultFindPlayerItem {
		results = append(results, &inventory.ItemInventory{
			InventoryId: v.Id.Hex(),
			PlayerId: v.PlayerId,
			ItemShowCase: &item.ItemShowCase{
				ItemId: v.ItemId,
				Title: itemMaps[v.ItemId].Title,
				Price: itemMaps[v.ItemId].Price,
				Damage: itemMaps[v.ItemId].Damage,
				ImageUrl: itemMaps[v.ItemId].ImageUrl,

			},
		})
	}

	countFindPlayerItem, err := u.inventoryRepository.CountPlayerItems(pctx, playerId)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return &models.PaginateResponse{
			Data: make([]*inventory.ItemInventory, 0),
			Total: countFindPlayerItem,
			Limit: req.Limit,
			First: models.FirstPaginate{
				Href: fmt.Sprintf("%s/%s?limit=%d", cfg.Paginate.InventoryNextBaseUrl, playerId, req.Limit),
			},
			Next: models.NextPaginate{
				Start: "",
				Href: "",
			},
		}, nil
	}

	return &models.PaginateResponse{
		Data: results,
		Total: countFindPlayerItem,
		Limit: req.Limit,
		First: models.FirstPaginate{
			Href: fmt.Sprintf("%s/%s?limit=%d", cfg.Paginate.InventoryNextBaseUrl, playerId, req.Limit),
		},
		Next: models.NextPaginate{
			Start: results[len(resultFindPlayerItem) -1].InventoryId,
			Href: fmt.Sprintf("%s/%s?limit=%d&start=%s", cfg.Paginate.InventoryNextBaseUrl, playerId, req.Limit, results[len(resultFindPlayerItem) -1].InventoryId),
		},
	}, nil
}