package server

import (
	"log"

	itemhandler "github.com/neabparinya11/Golang-Project/modules/item/itemHandler"
	itemPb "github.com/neabparinya11/Golang-Project/modules/item/itemPb"
	itemrepository "github.com/neabparinya11/Golang-Project/modules/item/itemRepository"
	itemusecase "github.com/neabparinya11/Golang-Project/modules/item/itemUsecase"
	"github.com/neabparinya11/Golang-Project/pkg/grpccon"
)

func (s *Server) ItemService() {
	repository := itemrepository.NewItemRepository(s.db)
	usecase := itemusecase.NewItemUsecase(repository)
	httpHandler := itemhandler.NewItemHttpHandler(s.cfg, usecase)
	grpcHandler := itemhandler.NewItemGrpcHandler(usecase)

	// gRPC
	go func(){
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.ItemUrl)
		itemPb.RegisterItemGrpcServiceServer(grpcServer, grpcHandler)
		log.Printf("Item gRPC server listening: %s", s.cfg.Grpc.ItemUrl)
		grpcServer.Serve(lis)
	}()
	_ = httpHandler

	//Route path
	item := s.app.Group("/item_v1")

	//Health check
	item.GET("/health", s.HealthCheckService)
	item.GET("/item/:item_id", httpHandler.FindOneItem)
	item.GET("/item", httpHandler.FindManyItems)
	item.PATCH("/item/:item_id", s.middleware.JwtAuthorization(s.middleware.RbacAuthorization(httpHandler.EditItem, []int{1, 0})))
	item.PATCH("/item/:item_id/is-active", s.middleware.JwtAuthorization(s.middleware.RbacAuthorization(httpHandler.EnableOrDisableItem, []int{1, 0})))
	item.POST("/item", s.middleware.JwtAuthorization(s.middleware.RbacAuthorization(httpHandler.CreateItem, []int{1, 0})))
}