package migration

import (
	"context"
	"log"

	"github.com/neabparinya11/Golang-Project/config"
	"github.com/neabparinya11/Golang-Project/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func PaymentDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return database.DbConn(pctx, cfg).Database("payment_db")
}

func PaymentMigrate(pctx context.Context, cfg *config.Config){
	db := PaymentDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	col := db.Collection("payment_queue")

	result, err := col.InsertOne(pctx, bson.M{"offset": -1}, nil)
	if err != nil {
		panic(err)
	}

	log.Println("Migrate payment completed: ", result)
}