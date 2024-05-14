package inventoryusecase

import inventoryrepository "github.com/neabparinya11/Golang-Project/modules/inventory/inventoryRepository"

type (
	InventoryUsecaseService interface{}

	InventoryUsecase struct {
		inventoryRepository inventoryrepository.InventoryRepositoryService
	}
)

func NewInventoryRepository(inventoryRepository inventoryrepository.InventoryRepositoryService) InventoryUsecaseService{
	return &InventoryUsecase{inventoryRepository: inventoryRepository}
}