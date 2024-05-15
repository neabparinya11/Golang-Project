package authhandler

import (
	"context"

	authPb "github.com/neabparinya11/Golang-Project/modules/auth/authPb"
	authusecase "github.com/neabparinya11/Golang-Project/modules/auth/authUsecase"
)

type (
	AuthGrpcHandler struct {
		authPb.UnimplementedAuthGrpcServiceServer
		authUsecase authusecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(authUsercase authusecase.AuthUsecaseService) *AuthGrpcHandler{
	return &AuthGrpcHandler{authUsecase: authUsercase}
}

func (g *AuthGrpcHandler) AccessTokenSearch(ctx context.Context, req *authPb.AccessTokenSearchRequest) (*authPb.AccessTokenSearchResponse, error){
	return g.authUsecase.AccessTokenSearch(ctx, req.AccessToken)
}

func (g *AuthGrpcHandler) RolesCount(ctx context.Context, req *authPb.RolesCountRequest) (*authPb.RolesCountResponse, error){
	return g.authUsecase.RoleCount(ctx)
}