package main

import (
	"context"
	"log"
	"os"

	"github.com/neabparinya11/Golang-Project/config"
	"github.com/neabparinya11/Golang-Project/pkg/database"
	"github.com/neabparinya11/Golang-Project/server"
)

func main() {
	ctx := context.Background()

	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2{
			log.Fatal("Error .env path required")
		} 
		return os.Args[1]
	}())

	//Database Connection
	db := database.DbConn(ctx, &cfg)
	defer db.Disconnect(ctx)

	//Start server
	server.Start(ctx, &cfg, db)
}