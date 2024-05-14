package server

import (
	"log"

	playerhandler "github.com/neabparinya11/Golang-Project/modules/player/playerHandler"
	playerPb "github.com/neabparinya11/Golang-Project/modules/player/playerPb"
	playerrepository "github.com/neabparinya11/Golang-Project/modules/player/playerRepository"
	playerusecase "github.com/neabparinya11/Golang-Project/modules/player/playerUsecase"
	"github.com/neabparinya11/Golang-Project/pkg/grpccon"
)

func (s *Server) PlayerService() {
	repository := playerrepository.NewPlayerRepository(s.db)
	usecase := playerusecase.NewPlayerUsecase(repository)
	httpHandler := playerhandler.NewPlayerHttpHandler(s.cfg, usecase)
	grpcHandler := playerhandler.NewPlayerGrpcHandler(usecase)
	queueHandler := playerhandler.NewPlayerQueueHandler(s.cfg, usecase)

	// gRPC
	go func(){
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.PlayerUrl)
		playerPb.RegisterPlayerGrpcServiceServer(grpcServer, grpcHandler)
		log.Printf("Player gRPC server listening: %s", s.cfg.Grpc.PlayerUrl)
		grpcServer.Serve(lis)
	}()

	_ = grpcHandler
	_ = queueHandler

	//Route path
	player := s.app.Group("/player_v1")

	//Health check
	player.GET("/health", s.HealthCheckService)

	player.POST("/player/register", httpHandler.CreatePlayer)
	player.POST("/player/add-money", httpHandler.AddPlayerMoney)
	player.GET("/player/:player_id", httpHandler.FindOnePlayerProfile)
	player.GET("/player/account/:player_id", httpHandler.GetPlayerSavingAccout)
}