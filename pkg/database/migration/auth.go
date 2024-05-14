package migration

import (
	"context"
	"log"

	"github.com/neabparinya11/Golang-Project/config"
	"github.com/neabparinya11/Golang-Project/modules/auth"
	"github.com/neabparinya11/Golang-Project/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthDbConn(pctx context.Context, cfg *config.Config) *mongo.Database{
	return database.DbConn(pctx, cfg).Database("auth_db")
}

func AuthMigrate(pctx context.Context, cfg *config.Config){
	db := AuthDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	col := db.Collection("auth")

	// indexs

	// auth
	indexs, _ := col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{ Keys: bson.D{{"_id", 1}}},
		{ Keys: bson.D{{"player_id", 1}}},
		{ Keys: bson.D{{"refresh_token", 1}}},
	})

	for _, index := range indexs {
		log.Printf("Index: %s", index)
	}

	// roles
	col = db.Collection("roles")

	indexs, _ = col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{ Keys: bson.D{{"_id", 1}}},
		{ Keys: bson.D{{"code", 1}}},
	})

	// roles data
	documents := func() []any {
		roles := []*auth.Role{
			{
				Title: "player",
				Code: 0,
			},
			{
				Title: "admin",
				Code: 1,
			},
		}
		docs := make([]any, 0)
		for _ , r := range roles{
			docs = append(docs, r)
		}
		return docs
	}()

	result, err := col.InsertMany(pctx, documents, nil)
	if err != nil {
		panic(err)
	}
	log.Println("Migrate auth complete: ", result.InsertedIDs)
}