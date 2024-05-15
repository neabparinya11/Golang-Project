package itemrepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/neabparinya11/Golang-Project/modules/item"
	"github.com/neabparinya11/Golang-Project/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	ItemRepositoryService interface {
		IsUniqueItem(pctx context.Context, title string) bool
		InsertOneItem(pctx context.Context, req *item.Item) (primitive.ObjectID, error)
		FindOneItem(pctx context.Context, itemId string) (*item.Item, error)
		FindManyItems(pctx context.Context, filter primitive.D, option []*options.FindOptions) ([]*item.ItemShowCase, error)
		CountItems(pctx context.Context, filter primitive.D) (int64, error)
	}

	ItemRepository struct {
		db *mongo.Client
	}
)

func NewItemRepository(db *mongo.Client) ItemRepositoryService {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) itemDbConn(pctx context.Context) *mongo.Database {
	_ = pctx
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

func (r *ItemRepository) FindOneItem(pctx context.Context, itemId string) (*item.Item, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.itemDbConn(ctx)
	col := db.Collection("items")

	result := new(item.Item)
	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjectId(itemId)}).Decode(result); err != nil {
		return nil, errors.New("error: Item not found")
	}

	return result, nil
}

func (r *ItemRepository) FindManyItems(pctx context.Context, filter primitive.D, option []*options.FindOptions) ([]*item.ItemShowCase, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.itemDbConn(ctx)
	col := db.Collection("items")

	cursors, err := col.Find(ctx, filter, option...)
	if err != nil {
		return make([]*item.ItemShowCase, 0), errors.New("error: Find many item failed")
	}

	results := make([]*item.ItemShowCase, 0)
	for cursors.Next(ctx) {
		result := new(item.Item)
		if err := cursors.Decode(result); err != nil {
			return make([]*item.ItemShowCase, 0), errors.New("error: Find many item failed")
		}

		results = append(results, &item.ItemShowCase{
			ItemId:   "item:" + result.Id.Hex(),
			Title:    result.Title,
			Price:    result.Price,
			Damage:   result.Damage,
			ImageUrl: result.ImageUrl,
		})
	}

	return results, nil
}

func (r *ItemRepository) CountItems(pctx context.Context, filter primitive.D) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.itemDbConn(ctx)
	col := db.Collection("items")

	counts, err := col.CountDocuments(ctx, filter)
	if err != nil {
		return -1, errors.New("error: Count item failed")
	}

	return counts, nil
}
