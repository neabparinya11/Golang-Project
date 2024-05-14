package playerhandler

import (
	"context"

	playerPb "github.com/neabparinya11/Golang-Project/modules/player/playerPb"
	playerusecase "github.com/neabparinya11/Golang-Project/modules/player/playerUsecase"
)

type (
	PlayerGrpcHandler struct {
		playerUsecase playerusecase.PlayerUsecaseService
		playerPb.UnimplementedPlayerGrpcServiceServer
	}
)

func NewPlayerGrpcHandler(playerUsecase playerusecase.PlayerUsecaseService) *PlayerGrpcHandler{
	return &PlayerGrpcHandler{playerUsecase: playerUsecase}
}

func (g *PlayerGrpcHandler) CreadentialSearch(ctx context.Context, req *playerPb.CreadentialSearchRequest) (*playerPb.PlayerProfile, error){
	return g.playerUsecase.FindOnePlayerCredential(ctx, req.Password, req.Email)
}

func (g *PlayerGrpcHandler) FindOnePlayerProfileToRefresh(ctx context.Context, req *playerPb.FindOnePlayerProfileToRefreshRequest) (*playerPb.PlayerProfile, error){
	return g.playerUsecase.FindOnePlayerProfileToRefresh(ctx, req.PlayerId)
}

func (g *PlayerGrpcHandler) GetPlayerSavingAccout(ctx context.Context, req *playerPb.GetPlayerSavingAccoutRequest) (*playerPb.GetPlayerSavingAccoutResponse, error){
	return nil, nil
}