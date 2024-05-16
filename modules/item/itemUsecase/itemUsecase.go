package itemusecase

import (
	"context"
	"errors"
	"fmt"
	"strings"

	itemPb "github.com/neabparinya11/Golang-Project/modules/item/itemPb"
	"github.com/neabparinya11/Golang-Project/modules/item"
	itemrepository "github.com/neabparinya11/Golang-Project/modules/item/itemRepository"
	"github.com/neabparinya11/Golang-Project/modules/models"
	"github.com/neabparinya11/Golang-Project/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	ItemUsecaseService interface{
		CreateItem(pctx context.Context, req *item.CreateItemRequest) (*item.ItemShowCase, error)
		FindOneItem(pctx context.Context, itemId string) (*item.ItemShowCase, error)
		FindManyItems(pctx context.Context, basePaginateUrl string, req *item.ItemSearchRequest) (*models.PaginateResponse, error)
		EditItem(pctx context.Context, itemId string, req *item.ItemUpdateRequest) (*item.ItemShowCase, error)
		EnableOrDisableItem(pctx context.Context, itemId string) (bool, error)
		FindItemInIds(pctx context.Context, req *itemPb.FindItemInIdsRequest) (*itemPb.FindItemInIdsResponse, error)
	}

	ItemUsecase struct {
		itemRepository itemrepository.ItemRepositoryService
	}
)

func NewItemUsecase(itemRepository itemrepository.ItemRepositoryService) ItemUsecaseService {
	return &ItemUsecase{itemRepository: itemRepository}
}

func (u *ItemUsecase) CreateItem(pctx context.Context, req *item.CreateItemRequest) (*item.ItemShowCase, error) {
	if !u.itemRepository.IsUniqueItem(pctx, req.Title){
		return nil, errors.New("error: Title has already")
	}

	result, err := u.itemRepository.InsertOneItem(pctx, &item.Item{
		Title: req.Title,
		Price: req.Price,
		Damage: req.Damage,
		ImageUrl: req.ImageUrl,
		UsageStatus: true,
		CreateAt: utils.LocalTime(),
		UpdateAt: utils.LocalTime(),
	})
	if err != nil {
		return nil, err
	}

	return u.FindOneItem(pctx, result.Hex())
}

func (u *ItemUsecase) FindOneItem(pctx context.Context, itemId string) (*item.ItemShowCase, error) {
	result, err := u.itemRepository.FindOneItem(pctx, itemId)
	if err != nil {
		return nil, err
	}

	return &item.ItemShowCase{
		ItemId: "item:" + result.Id.Hex(),
		Title: result.Title,
		Price: result.Price,
		Damage: result.Damage,
		ImageUrl: result.ImageUrl,
	}, nil
}

func (u *ItemUsecase) FindManyItems(pctx context.Context, basePaginateUrl string, req *item.ItemSearchRequest) (*models.PaginateResponse, error) {
	findItemFilter := bson.D{}
	findItemOptions := make([]*options.FindOptions, 0)

	countItemFilter := bson.D{}

	if req.Start != "" {
		req.Start = strings.TrimPrefix(req.Start, "item:")
		findItemFilter = append(findItemFilter, bson.E{"_id", bson.D{{"$gt", utils.ConvertToObjectId(req.Start)}}})
	}

	if req.Title != "" {
		findItemFilter = append(findItemFilter, bson.E{"title", primitive.Regex{ Pattern: req.Title, Options: "i"}})
		countItemFilter = append(countItemFilter, bson.E{"title", primitive.Regex{ Pattern: req.Title, Options: "i"}})
	}

	findItemFilter = append(findItemFilter, bson.E{"usage_status", true})
	countItemFilter = append(countItemFilter, bson.E{"usage_status", true})

	findItemOptions = append(findItemOptions, options.Find().SetSort(bson.D{{"_id", 1}}))
	findItemOptions = append(findItemOptions, options.Find().SetLimit(int64(req.Limit)))

	resultFindItem, err := u.itemRepository.FindManyItems(pctx, findItemFilter, findItemOptions)
	if err != nil {
		return nil, err
	}

	countFindItem, err := u.itemRepository.CountItems(pctx, countItemFilter)
	if err != nil {
		return nil, err
	}

	if len(resultFindItem) == 0 {
		return &models.PaginateResponse{
			Data: make([]*item.ItemShowCase, 0),
			Total: countFindItem,
			Limit: req.Limit,
			First: models.FirstPaginate{
				Href: fmt.Sprintf("%s?limit=%d&title=%s", basePaginateUrl, req.Limit, req.Title),
			},
			Next: models.NextPaginate{
				Start: "",
				Href: "",
			},
		}, nil
	}

	return &models.PaginateResponse{
		Data: resultFindItem,
		Total: countFindItem,
		Limit: req.Limit,
		First: models.FirstPaginate{
			Href: fmt.Sprintf("%s?limit=%d&title=%s", basePaginateUrl, req.Limit, req.Title),
		},
		Next: models.NextPaginate{
			Start: resultFindItem[len(resultFindItem) -1].ItemId,
			Href: fmt.Sprintf("%s?limit=%d&title=%s&start=%s", basePaginateUrl, req.Limit, req.Title, resultFindItem[len(resultFindItem) -1].ItemId),
		},
	}, nil
}

func (u *ItemUsecase) EditItem(pctx context.Context, itemId string, req *item.ItemUpdateRequest) (*item.ItemShowCase, error) {
	updateReq := bson.M{}

	if req.Title != "" {
		if !u.itemRepository.IsUniqueItem(pctx, req.Title) {
			return nil, errors.New("error: Title is already exist")
		}
		updateReq["title"] = req.Title
	}

	if req.ImageUrl != "" {
		updateReq["image_url"] = req.ImageUrl
	}

	if req.Damage > 0 {
		updateReq["damage"] = req.Damage
	}

	if req.Price >= 0 {
		updateReq["price"] = req.Price
	}

	if err := u.itemRepository.UpdateOneItems(pctx, itemId, updateReq); err != nil {
		return nil, err
	}
	
	return u.FindOneItem(pctx, itemId)
}

func (u *ItemUsecase) EnableOrDisableItem(pctx context.Context, itemId string) (bool, error) {
	result, err := u.itemRepository.FindOneItem(pctx, itemId)
	if err != nil {
		return false, err
	}

	if err := u.itemRepository.EnableOrDisableItem(pctx, itemId, !result.UsageStatus); err != nil {
		return false, err
	}

	return !result.UsageStatus, nil
}

func (u *ItemUsecase) FindItemInIds(pctx context.Context, req *itemPb.FindItemInIdsRequest) (*itemPb.FindItemInIdsResponse, error) {
	filter := bson.D{}

	objectIds := make([]primitive.ObjectID, 0)
	for _, itemId := range req.Ids {
		objectIds = append(objectIds, utils.ConvertToObjectId(strings.TrimPrefix(itemId, "item:")))
	}
	
	filter = append(filter, bson.E{"_id", bson.D{{"$in", objectIds}}})
	filter = append(filter, bson.E{"usage_status", true})

	results, err := u.itemRepository.FindManyItems(pctx, filter, nil)
	if err != nil {
		return nil, err
	}

	resultToRes := make([]*itemPb.Item, 0)
	for _, result := range results{
		resultToRes = append(resultToRes, &itemPb.Item{
			Id: result.ItemId,
			Title: result.Title,
			Price: result.Price,
			ImageUrl: result.ImageUrl,
			Damage: int32(result.Damage),
		})
	}
	return &itemPb.FindItemInIdsResponse{
		Items: resultToRes,
	}, nil
}