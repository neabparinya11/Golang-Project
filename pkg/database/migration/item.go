package migration

import (
	"context"
	"log"

	"github.com/neabparinya11/Golang-Project/config"
	"github.com/neabparinya11/Golang-Project/modules/item"
	"github.com/neabparinya11/Golang-Project/pkg/database"
	"github.com/neabparinya11/Golang-Project/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ItemDbConn(pctx context.Context, cfg *config.Config) *mongo.Database{
	return database.DbConn(pctx, cfg).Database("item_db")
}

func ItemMigrate(pctx context.Context, cfg *config.Config){
	db := ItemDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	col := db.Collection("items")

	indexs, _ := col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{ Keys: bson.D{{"_id", 1}}},
		{ Keys: bson.D{{"title", 1}}},
	})

	for _, index := range indexs{
		log.Printf("Index: %s", index)
	}

	documents := func() []any {
		roles := []*item.Item{
			{
				Title: "Diamond Sword",
				Price: 1000,
				ImageUrl: "",
				UsageStatus: true,
				Damage: 100,
				CreateAt: utils.LocalTime(),
				UpdateAt: utils.LocalTime(),
			},
			{
				Title: "Wooden Sword",
				Price: 100,
				ImageUrl: "",
				UsageStatus: true,
				Damage: 20,
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

	log.Println("Migrate item completed: ", results)
}