package server

import (
	"log"

	inventoryhandler "github.com/neabparinya11/Golang-Project/modules/inventory/inventoryHandler"
	inventoryPb "github.com/neabparinya11/Golang-Project/modules/inventory/inventoryPb"
	inventoryrepository "github.com/neabparinya11/Golang-Project/modules/inventory/inventoryRepository"
	inventoryusecase "github.com/neabparinya11/Golang-Project/modules/inventory/inventoryUsecase"
	"github.com/neabparinya11/Golang-Project/pkg/grpccon"
)

func (s *Server) InventoryService() {
	repository := inventoryrepository.NewInventoryRepository(s.db)
	usecase := inventoryusecase.NewInventoryRepository(repository)
	httpHandler := inventoryhandler.NewInventoryHttpHandler(s.cfg, usecase)
	grpcHandler := inventoryhandler.NewInventoryGrpcHandler(usecase)
	queueHandler := inventoryhandler.NewInventoryQueueHandler(s.cfg, usecase)

	// gRPC
	go func(){
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.InventoryUrl)
		inventoryPb.RegisterInventoryGrpcServiceServer(grpcServer, grpcHandler)
		log.Printf("Inventory gRPC server listening: %s", s.cfg.Grpc.InventoryUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler
	_ = queueHandler

	//Route path
	inventory := s.app.Group("/inventory_v1")

	//Health check
	inventory.GET("/health", s.HealthCheckService)
}