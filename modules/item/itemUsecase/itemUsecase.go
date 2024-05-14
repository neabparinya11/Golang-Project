package itemusecase

import itemrepository "github.com/neabparinya11/Golang-Project/modules/item/itemRepository"

type (
	ItemUsecaseService interface{}

	ItemUsecase struct {
		itemRepository itemrepository.ItemRepositoryService
	}
)

func NewItemUsecase(itemRepository itemrepository.ItemRepositoryService) ItemUsecaseService {
	return &ItemUsecase{itemRepository: itemRepository}
}