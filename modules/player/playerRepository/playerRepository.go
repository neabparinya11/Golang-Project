package playerrepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/neabparinya11/Golang-Project/modules/models"
	"github.com/neabparinya11/Golang-Project/modules/player"
	"github.com/neabparinya11/Golang-Project/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	PlayerRepositoryService interface{
		IsUniquePlayer(pctx context.Context, email, username string) bool
		InsertOnePlayer(pctx context.Context, req *player.Player) (primitive.ObjectID, error)
		FindOnePlayerProfile(pctx context.Context, playerId string) (*player.PlayerProfileBson, error)
		InsertOnePlayerTransaction(pctx context.Context, req *player.PlayerTransaction) error
		GetPlayerSavingAccout(pctx context.Context, playerId string) (*player.PlayerSavingAccout, error)
		FindOnePlayerCredential(pctx context.Context, email string) (*player.Player, error)
		FindOnePlayerProfileToRefresh(pctx context.Context, playerId string) (*player.Player, error)
		GetOffset(pctx context.Context) (int64, error)
		UpserOffset(pctx context.Context, kafkaOffset int64) error
	}

	PlayerRepository struct {
		db *mongo.Client
	}
)

func NewPlayerRepository(db *mongo.Client) PlayerRepositoryService{
	return &PlayerRepository{db: db}
}

func (r *PlayerRepository) playerDbConn(pctx context.Context) *mongo.Database{
	_ = pctx
	return r.db.Database("player_db")
}

func (r *PlayerRepository) IsUniquePlayer(pctx context.Context, email, username string) bool {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection("players")

	player := new(player.Player)
	if err := col.FindOne(
		ctx,
		bson.M{"$or": []bson.M{
			{ "username": username},
			{ "email": email},
		}},
	).Decode(player); err != nil {
		log.Printf("Error: IsUniquePlayer: %s", err.Error())
		return true
	}
	return false
}

func (r *PlayerRepository) InsertOnePlayer(pctx context.Context, req *player.Player) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection("players")

	playerId, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: InsertOnePlayer failed: %s", err.Error())
		return primitive.NilObjectID, errors.New("error: insert one player failed")
	}
	return playerId.InsertedID.(primitive.ObjectID), nil
}

func (r *PlayerRepository) FindOnePlayerProfile(pctx context.Context, playerId string) (*player.PlayerProfileBson, error){
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection("players")

	result := new(player.PlayerProfileBson)
	if err := col.FindOne(
		ctx,
		bson.M{"_id": utils.ConvertToObjectId(playerId)},
		options.FindOne().SetProjection(bson.M{
			"_id": 1,
			"email": 1,
			"username": 1,
			"create_at": 1,
			"update_at": 1,
		}),
	).Decode(result); err != nil {
		log.Printf("Error: FindOnePlayerProfile %s", err.Error())
		return nil, errors.New("error: player not found")
	}

	return result, nil
}

func (r *PlayerRepository) InsertOnePlayerTransaction(pctx context.Context, req *player.PlayerTransaction) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection("player_transactions")

	_, err := col.InsertOne(ctx, req)
	if err != nil {
		return errors.New("error: Insert one player transaction failed")
	}

	return nil
}

func (r *PlayerRepository) GetPlayerSavingAccout(pctx context.Context, playerId string) (*player.PlayerSavingAccout, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection("player_transactions")

	result := new(player.PlayerSavingAccout)

	filter := bson.A{
		bson.D{{"$match", bson.D{{"player_id", playerId}}}},
		bson.D{{"$group", bson.D{{"_id", "$player_id"}, {"balance", bson.D{{"$sum", "$amount"}}}}}},
		bson.D{{"$project", bson.D{{"player_id", "$_id"}, {"_id", 0}, {"balance", 1}}}},
	}

	cursors, err := col.Aggregate(ctx, filter)
	if err != nil {
		return nil, errors.New("error: Failed to get saving account")
	}

	for cursors.Next(ctx){
		if err := cursors.Decode(result); err != nil {
			return nil, errors.New("error: Failed to decode saving account")
		}
	}

	return result, nil
}

func (r *PlayerRepository) FindOnePlayerCredential(pctx context.Context, email string) (*player.Player, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection("players")

	result := new(player.Player)

	if err := col.FindOne(ctx, bson.M{"email": email}).Decode(result); err != nil {
		return nil, errors.New("error: Plyer credential search not found")
	}

	return result, nil
}

func (r *PlayerRepository) FindOnePlayerProfileToRefresh(pctx context.Context, playerId string) (*player.Player, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection("players")

	result := new(player.Player)

	if err := col.FindOne(ctx, bson.M{"_id": utils.ConvertToObjectId(playerId)}).Decode(result); err != nil {
		return nil, errors.New("error: Plyer not found")
	}

	return result, nil
}

func (r *PlayerRepository) GetOffset(pctx context.Context) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection("player_transactions_queue")

	result := new(models.KafkaOffset)
	if err := col.FindOne(ctx, bson.M{}).Decode(result); err != nil {
		return -1, errors.New("error")
	}

	return result.Offset, nil
}

func (r *PlayerRepository) UpserOffset(pctx context.Context, kafkaOffset int64) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection("player_transactions_queue")

	result, err := col.UpdateOne(ctx, bson.M{}, bson.M{"$set": bson.M{"offset": kafkaOffset}})
	if err != nil {
		return errors.New("error: Upseroffset failed")
	}

	log.Printf("Info: Upseroffset result: %v", result)

	return nil
}