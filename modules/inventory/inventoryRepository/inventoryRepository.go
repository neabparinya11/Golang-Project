package inventoryrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	InventoryRepositoryService interface{}

	InventoryRepository struct {
		db *mongo.Client
	}
)

func NewInventoryRepository(db *mongo.Client) InventoryRepositoryService{
	return &InventoryRepository{db: db}
}

func (in *InventoryRepository) inventoryDbConn(pctx context.Context) *mongo.Database{
	return in.db.Database("inventory_db")
}