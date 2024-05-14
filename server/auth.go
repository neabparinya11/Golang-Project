package server

import (
	"log"

	authhandler "github.com/neabparinya11/Golang-Project/modules/auth/authHandler"
	authPb "github.com/neabparinya11/Golang-Project/modules/auth/authPb"
	authrepository "github.com/neabparinya11/Golang-Project/modules/auth/authRepository"
	authusecase "github.com/neabparinya11/Golang-Project/modules/auth/authUsecase"
	"github.com/neabparinya11/Golang-Project/pkg/grpccon"
)

func (s *Server) AuthService() {
	repository := authrepository.NewAuthRepository(s.db)
	usecase := authusecase.NewAuthUsecase(repository)
	httpHandler := authhandler.NewAuthHttpHandler(s.cfg, usecase)
	grpcHandler := authhandler.NewAuthGrpcHandler(usecase)

	//gRPC
	go func(){
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.AuthUrl)
		authPb.RegisterAuthGrpcServiceServer(grpcServer, grpcHandler)
		log.Printf("Auth gRPC server listening: %s", s.cfg.Grpc.AuthUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler
	
	//Route path
	auth := s.app.Group("/auth_v1")

	//Health check
	auth.GET("/health", s.HealthCheckService)

	auth.POST("/auth/login", httpHandler.Login)
	auth.POST("/auth/refresh-token", httpHandler.RefreshToken)
	auth.POST("/auth/logout", httpHandler.Logout)
}