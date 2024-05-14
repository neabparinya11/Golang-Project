package itemrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	ItemRepositoryService interface{}

	ItemRepository struct {
		db *mongo.Client
	}
)

func NewItemRepository(db *mongo.Client) ItemRepositoryService{
	return &ItemRepository{db: db}
}

func (i *ItemRepository) itemDbConn(pctx context.Context) *mongo.Database{
	return i.db.Database("item_db")
}