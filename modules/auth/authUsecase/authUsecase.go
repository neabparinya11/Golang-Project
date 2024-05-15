package authusecase

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/neabparinya11/Golang-Project/config"
	"github.com/neabparinya11/Golang-Project/modules/auth"
	authrepository "github.com/neabparinya11/Golang-Project/modules/auth/authRepository"
	"github.com/neabparinya11/Golang-Project/modules/player"
	playerPb "github.com/neabparinya11/Golang-Project/modules/player/playerPb"
	authPb "github.com/neabparinya11/Golang-Project/modules/auth/authPb"
	"github.com/neabparinya11/Golang-Project/pkg/jwtauth"
	"github.com/neabparinya11/Golang-Project/pkg/utils"
)

type (
	AuthUsecaseService interface{
		Login(pctx context.Context, cfg *config.Config, req *auth.PlayerLoginRequest) (*auth.ProfileIntercepter, error)
		RefreshToken(pctx context.Context, cfg *config.Config, req *auth.PlayerRefreshToken) (*auth.ProfileIntercepter, error)
		Logout(pctx context.Context, credentialId string) (int64, error)
		AccessTokenSearch(pctx context.Context, accessToken string) (*authPb.AccessTokenSearchResponse, error)
		RoleCount(pctx context.Context) (*authPb.RolesCountResponse, error)
	}

	AuthUsecase struct {
		authRepository authrepository.AuthRepositoryService
	}
)

func NewAuthUsecase(authRepository authrepository.AuthRepositoryService) AuthUsecaseService {
	return &AuthUsecase{authRepository: authRepository}
}

func (u *AuthUsecase) Login(pctx context.Context, cfg *config.Config, req *auth.PlayerLoginRequest) (*auth.ProfileIntercepter, error) {
	profile, err := u.authRepository.CredentialSearch(pctx, cfg.Grpc.PlayerUrl, &playerPb.CreadentialSearchRequest{
		Email: req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	profile.Id = "player:" + profile.Id

	accessToken := jwtauth.NewAccessToken(cfg.Jwt.AccessSecretKey, cfg.Jwt.AccessDuration, &jwtauth.Claims{
		PlayerId: profile.Id,
		RoleCode: int(profile.RoleCode),
	}).SignToken()

	refreshToken := jwtauth.NewRefreshToken(cfg.Jwt.RefreshSecretKey, cfg.Jwt.RefreshDuration, &jwtauth.Claims{
		PlayerId: profile.Id,
		RoleCode: int(profile.RoleCode),
	}).SignToken()

	credentialId, err := u.authRepository.InsertOnePlayerCredential(pctx, &auth.Credential{
		PlayerId: profile.Id,
		RoleCode: int(profile.RoleCode),
		AccessToken: accessToken,
		RefreshToken: refreshToken,
		CreateAt: utils.LocalTime(),
		UpdateAt: utils.LocalTime(),
	})
	if err != nil {
		return nil, err
	}

	credential, err := u.authRepository.FindOnePlayerCredential(pctx, credentialId.Hex())
	if err != nil {
		return nil, err
	}

	loc, _ := time.LoadLocation("Asia/Bangkok")

	return &auth.ProfileIntercepter{
		PlayerProfile: &player.PlayerProfile{
			Id: profile.Id,
			Email: profile.Email,
			Username: profile.Username,
			CreateAt: utils.ConvertStringTimeToTimelayout(profile.CreateAt).In(loc),
			UpdateAt: utils.ConvertStringTimeToTimelayout(profile.UpdateAt).In(loc),
		},
		Credential: &auth.CredentialResponse{
			Id: credential.Id.Hex(),
			PlayerId: credential.PlayerId,
			RoleCode: credential.RoleCode,
			AccessToken: credential.AccessToken,
			RefreshToken: credential.RefreshToken,
			CreateAt: credential.CreateAt.In(loc),
			UpdateAt: credential.UpdateAt.In(loc),
		},
	}, nil
}

func (u *AuthUsecase) RefreshToken(pctx context.Context, cfg *config.Config, req *auth.PlayerRefreshToken) (*auth.ProfileIntercepter, error) {
	claims, err := jwtauth.ParseToken(cfg.Jwt.RefreshSecretKey, req.RefreshToken)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	profile, err := u.authRepository.FindOnePlayerProfileToRefresh(pctx, cfg.Grpc.AuthUrl, &playerPb.FindOnePlayerProfileToRefreshRequest{
		PlayerId: strings.TrimPrefix(claims.PlayerId, "player:"),
	})
	if err != nil {
		return nil, err
	}

	accessToken := jwtauth.NewAccessToken(cfg.Jwt.AccessSecretKey, cfg.Jwt.AccessDuration, &jwtauth.Claims{
		PlayerId: profile.Id,
		RoleCode: int(profile.RoleCode),
	}).SignToken()

	reloadToken := jwtauth.ReloadToken(cfg.Jwt.RefreshSecretKey, claims.ExpiresAt.Unix(), &jwtauth.Claims{
		PlayerId: profile.Id,
		RoleCode: int(profile.RoleCode),
	})

	if err := u.authRepository.UpdateOnePlayerCredential(pctx, req.CredentialId, &auth.UpdateRefreshTokenRequest{
		PlayerId: profile.Id,
		AccessToken: accessToken,
		RefreshToken: reloadToken,
		UpdateAt: utils.LocalTime(),
	}); err != nil {
		return nil, err
	}

	credential, err := u.authRepository.FindOnePlayerCredential(pctx, req.CredentialId)
	if err != nil {
		return nil, err
	}

	loc, _ := time.LoadLocation("Asia/Bangkok")

	return &auth.ProfileIntercepter{
		PlayerProfile: &player.PlayerProfile{
			Id: "player:" + profile.Id,
			Email: profile.Email,
			Username: profile.Username,
			CreateAt: utils.ConvertStringTimeToTimelayout(profile.CreateAt).In(loc),
			UpdateAt: utils.ConvertStringTimeToTimelayout(profile.UpdateAt).In(loc),
		},
		Credential: &auth.CredentialResponse{
			Id: credential.Id.Hex(),
			PlayerId: credential.PlayerId,
			RoleCode: credential.RoleCode,
			AccessToken: credential.AccessToken,
			RefreshToken: credential.RefreshToken,
			CreateAt: credential.CreateAt.In(loc),
			UpdateAt: credential.UpdateAt.In(loc),
		},
	}, nil
}

func (u *AuthUsecase) Logout(pctx context.Context, credentialId string) (int64, error) {
	return u.authRepository.DeleteOnePlayerCredential(pctx, credentialId)
}

func (u *AuthUsecase) AccessTokenSearch(pctx context.Context, accessToken string) (*authPb.AccessTokenSearchResponse, error) {
	credential, err := u.authRepository.FindOneAccessToken(pctx, accessToken)
	if err != nil {
		return &authPb.AccessTokenSearchResponse{
			IsValid: false,
		},err
	}

	if credential == nil {
		return &authPb.AccessTokenSearchResponse{
			IsValid: false,
		}, errors.New("error: access token is invalid")
	}
	return &authPb.AccessTokenSearchResponse{
		IsValid: true,
	},nil
}

func (u *AuthUsecase) RoleCount(pctx context.Context) (*authPb.RolesCountResponse, error) {
	result, err := u.authRepository.RoleCount(pctx)
	if err != nil {
		return nil, err
	}

	return &authPb.RolesCountResponse{
		Count: result,
	},nil
}