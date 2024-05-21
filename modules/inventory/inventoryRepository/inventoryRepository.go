package inventoryrepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/neabparinya11/Golang-Project/modules/inventory"
	itemPb "github.com/neabparinya11/Golang-Project/modules/item/itemPb"
	"github.com/neabparinya11/Golang-Project/modules/models"
	"github.com/neabparinya11/Golang-Project/pkg/grpccon"
	"github.com/neabparinya11/Golang-Project/pkg/jwtauth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	InventoryRepositoryService interface{
		FindItemInIds(pctx context.Context, grpcUrl string, req *itemPb.FindItemInIdsRequest) (*itemPb.FindItemInIdsResponse , error)
		FindPlayerItems(pctx context.Context, filter primitive.D, option []*options.FindOptions) ([]*inventory.Inventory, error)
		CountPlayerItems(pctx context.Context, playerId string) (int64, error)
		GetOffset(pctx context.Context) (int64, error)
		UpserOffset(pctx context.Context, kafkaOffset int64) error
	}

	InventoryRepository struct {
		db *mongo.Client
	}
)

func NewInventoryRepository(db *mongo.Client) InventoryRepositoryService{
	return &InventoryRepository{db: db}
}

func (r *InventoryRepository) inventoryDbConn(pctx context.Context) *mongo.Database{
	_ = pctx
	return r.db.Database("inventory_db")
}

func (r *InventoryRepository) FindItemInIds(pctx context.Context, grpcUrl string, req *itemPb.FindItemInIdsRequest) (*itemPb.FindItemInIdsResponse , error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	jwtauth.SetApiKeyInContext(&ctx)
	conn, err := grpccon.NewGrpcClient(grpcUrl)
	if err != nil {
		return nil, errors.New("error: gRPC connected failed")
	}

	result, err := conn.Item().FindItemInIds(ctx, req)
	if err != nil {
		return nil, errors.New("error: Item not found")
	}

	if result == nil {
		return nil, errors.New("error: Item not found")
	}

	if len(result.Items) == 0 {
		return nil, errors.New("error: Item not found")
	}

	return result, nil
}

func (r *InventoryRepository) FindPlayerItems(pctx context.Context, filter primitive.D, option []*options.FindOptions) ([]*inventory.Inventory, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.inventoryDbConn(ctx)
	col := db.Collection("players_inventory")

	cursors, err := col.Find(ctx, filter, option...)
	if err != nil {
		return nil, errors.New("error: Player item not found")
	}

	results := make([]*inventory.Inventory, 0)
	for cursors.Next(ctx) {
		result := new(inventory.Inventory)
		if err := cursors.Decode(result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

func (r *InventoryRepository) CountPlayerItems(pctx context.Context, playerId string) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.inventoryDbConn(ctx)
	col := db.Collection("players_inventory")

	counts, err := col.CountDocuments(ctx, bson.M{"player_id": playerId})
	if err != nil {
		return -1, errors.New("error: Count player items failed")
	}

	return counts, nil
}

func (r *InventoryRepository) GetOffset(pctx context.Context) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.inventoryDbConn(ctx)
	col := db.Collection("players_inventory_queue")

	result := new(models.KafkaOffset)
	if err := col.FindOne(ctx, bson.M{}).Decode(result); err != nil {
		return -1, errors.New("error")
	}

	return result.Offset, nil
}

func (r *InventoryRepository) UpserOffset(pctx context.Context, kafkaOffset int64) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.inventoryDbConn(ctx)
	col := db.Collection("players_inventory_queue")

	result, err := col.UpdateOne(ctx, bson.M{}, bson.M{"$set": bson.M{"offset": kafkaOffset}})
	if err != nil {
		return errors.New("error: Upseroffset failed")
	}

	log.Printf("Info: Upseroffset result: %v", result)

	return nil
}