package migration

import (
	"context"
	"log"

	"github.com/neabparinya11/Golang-Project/config"
	"github.com/neabparinya11/Golang-Project/modules/player"
	"github.com/neabparinya11/Golang-Project/pkg/database"
	"github.com/neabparinya11/Golang-Project/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func PlayerDbConn(pctx context.Context, cfg *config.Config) *mongo.Database{
	return database.DbConn(pctx, cfg).Database("player_db")
}

func PlayerMigrate(pctx context.Context, cfg *config.Config){
	db := PlayerDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	col := db.Collection("player_transactions")

	// player
	indexs, _ := col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{ Keys: bson.D{{"_id", 1}}},
		{ Keys: bson.D{{"player_id", 1}}},
	})
	log.Println(indexs)

	col = db.Collection("players")

	indexs , _ = col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"_id", 1}}},
		{ Keys: bson.D{{"email", 1}}},
	})
	log.Println(indexs)

	documents := func() []any{
		roles := []*player.Player{
			{
				Email: "neabparinya@gmail.com",
				Password: "123456",
				Username: "Neab012",
				PlayerRoles: []player.PlayerRole{
					{
						RoleTitle: "player",
						RoleCode: 0,
					},
					{
						RoleTitle: "admin",
						RoleCode: 1,
					},
				},
				CreateAt: utils.LocalTime(),
				UpdateAt: utils.LocalTime(),
			},
		}

		docs := make([]any, 0)
		for _, r := range roles{
			docs = append(docs, r)
		}
		return docs
	}()

	results, err := col.InsertMany(pctx, documents, nil)
	if err != nil {
		panic(err)
	}

	log.Println("Migrate player complete: ", results)
	
	playerTransactions := make([]any, 0)
	for _, p := range results.InsertedIDs {
		playerTransactions = append(playerTransactions, &player.PlayerTransaction{
			PlayerId: "player:" + p.(primitive.ObjectID).Hex(),
			Amount: 10000,
			CreateAt: utils.LocalTime(),
		})
	}
	col = db.Collection("player_transactions")
	results, err = col.InsertMany(pctx, playerTransactions, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate player_transactions completed: ", results)

	col = db.Collection("player_transactions_queue")
	result, err := col.InsertOne(pctx, bson.M{"offset": -1}, nil)
	if err != nil {
		panic(err)
	}

	log.Println("Migrate player_transactions_queue completed: ", result)
}