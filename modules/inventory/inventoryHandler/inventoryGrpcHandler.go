package inventoryhandler

import (
	"context"

	inventoryPb "github.com/neabparinya11/Golang-Project/modules/inventory/inventoryPb"
	inventoryusecase "github.com/neabparinya11/Golang-Project/modules/inventory/inventoryUsecase"
)

type (
	InventoryGrpcHandler struct {
		inventoryUsecase inventoryusecase.InventoryUsecaseService
		inventoryPb.UnimplementedInventoryGrpcServiceServer
	}
)

func NewInventoryGrpcHandler(inventoryUsecase inventoryusecase.InventoryUsecaseService) *InventoryGrpcHandler{
	return &InventoryGrpcHandler{inventoryUsecase: inventoryUsecase}
}

func (in *InventoryGrpcHandler) IsAvailableToSell(ctx context.Context, req *inventoryPb.IsAvailableToSellRequest) (*inventoryPb.IsAvailableToSellResponse, error){
	return nil, nil
}