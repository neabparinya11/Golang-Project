package itemrepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/neabparinya11/Golang-Project/modules/item"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	ItemRepositoryService interface{
		IsUniqueItem(pctx context.Context, title string) bool
	}

	ItemRepository struct {
		db *mongo.Client
	}
)

func NewItemRepository(db *mongo.Client) ItemRepositoryService{
	return &ItemRepository{db: db}
}

func (r *ItemRepository) itemDbConn(pctx context.Context) *mongo.Database{
	return r.db.Database("item_db")
}

func (r *ItemRepository) IsUniqueItem(pctx context.Context, title string) bool {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.itemDbConn(ctx)
	col := db.Collection("items")

	item := new(item.Item)
	if err := col.FindOne(ctx, bson.M{"title": title}).Decode(item); err != nil {
		log.Printf("Error: IsUniqueItem: %s", err.Error())
		return true
	}
	return false
}

func (r *ItemRepository) InsertOneItem(pctx context.Context, req *item.Item) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.itemDbConn(ctx)
	col := db.Collection("items")

	itemId, err := col.InsertOne(ctx, req)
	if err != nil {
		return primitive.NilObjectID, errors.New("error: Insert one item failed")
	}

	return itemId.InsertedID.(primitive.ObjectID), nil
}