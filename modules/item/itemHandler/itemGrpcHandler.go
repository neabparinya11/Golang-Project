package itemhandler

import (
	"context"

	itemPb "github.com/neabparinya11/Golang-Project/modules/item/itemPb"
	itemusecase "github.com/neabparinya11/Golang-Project/modules/item/itemUsecase"
)

type (
	ItemGrpcHandler struct {
		itemUsecase itemusecase.ItemUsecaseService
		itemPb.UnimplementedItemGrpcServiceServer
	}
)

func NewItemGrpcHandler(itemUsecase itemusecase.ItemUsecaseService) *ItemGrpcHandler{
	return &ItemGrpcHandler{itemUsecase: itemUsecase}
}

func (i *ItemGrpcHandler) FindItemInIds(ct context.Context, req *itemPb.FindItemInIdsRequest) (*itemPb.FindItemInIdsResponse, error){
	return nil, nil
}