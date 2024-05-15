package middlewarerepository

import (
	"context"
	"errors"
	"time"

	authPb "github.com/neabparinya11/Golang-Project/modules/auth/authPb"
	"github.com/neabparinya11/Golang-Project/pkg/grpccon"
	"github.com/neabparinya11/Golang-Project/pkg/jwtauth"
)

type (
	MiddlewareRepositoryService interface{
		AccessTokenSearch(pctx context.Context, grpcUrl, accessToken string) error
		RoleCount(pctx context.Context, grpcUrl string) (int64, error)
	}

	MiddlewareRepository struct {
	}
)

func NewMiddlewareRepository() MiddlewareRepositoryService {
	return &MiddlewareRepository{}
}

func (r *MiddlewareRepository) AccessTokenSearch(pctx context.Context, grpcUrl, accessToken string) error {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	conn, err := grpccon.NewGrpcClient(grpcUrl)
	if err != nil {
		return errors.New("error: gRPC connection failed")
	}

	jwtauth.SetApiKeyInContext(&ctx)

	result, err := conn.Auth().AccessTokenSearch(ctx, &authPb.AccessTokenSearchRequest{
		AccessToken: accessToken,
	})
	if err != nil {
		return errors.New("error: email or password is incorrect")
	}

	if result == nil {
		return errors.New("error: access token is invalid")
	}

	if !result.IsValid {
		return errors.New("error: access token is invalid")
	}
	return nil
}

func (r *MiddlewareRepository) RoleCount(pctx context.Context, grpcUrl string) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	
	conn, err := grpccon.NewGrpcClient(grpcUrl)
	if err != nil {
		return -1 , errors.New("error: gRPC connection failed")
	}
	jwtauth.SetApiKeyInContext(&ctx)

	result, err := conn.Auth().RolesCount(ctx, &authPb.RolesCountRequest{})
	if err != nil {
		return -1, errors.New("error: email or password is incorrect")
	}

	if result == nil {
		return -1, errors.New("error: roles count is failed")
	}
	return result.Count, nil
}