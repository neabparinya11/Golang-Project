package middlewareusecase

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/neabparinya11/Golang-Project/config"
	middlewarerepository "github.com/neabparinya11/Golang-Project/modules/middleware/middlewareRepository"
	"github.com/neabparinya11/Golang-Project/pkg/jwtauth"
	"github.com/neabparinya11/Golang-Project/pkg/rbac"
)

type (
	MiddlewareUsecaseService interface{
		JwtAuthorization(c echo.Context, cfg *config.Config, accessToken string) (echo.Context, error)
		RbacAuthorization(c echo.Context, cfg *config.Config, expected []int) (echo.Context, error)
		PlayerIdParamValidation(c echo.Context) (echo.Context, error)
	}

	MiddlewareUsecase struct {
		middlewareRepository middlewarerepository.MiddlewareRepositoryService
	}
)

func NewMiddlewareUsecase(middlewareRepository middlewarerepository.MiddlewareRepositoryService) MiddlewareUsecaseService{
	return &MiddlewareUsecase{middlewareRepository: middlewareRepository}
}

func (u *MiddlewareUsecase) JwtAuthorization(c echo.Context, cfg *config.Config, accessToken string) (echo.Context, error) {
	ctx := c.Request().Context()

	claims, err := jwtauth.ParseToken(cfg.Jwt.AccessSecretKey, accessToken)
	if err != nil {
		return nil, err
	}

	if err := u.middlewareRepository.AccessTokenSearch(ctx, cfg.Grpc.AuthUrl, accessToken); err != nil {
		return nil, err
	}
	c.Set("player_id", claims.PlayerId)
	c.Set("role_code", claims.RoleCode)

	return c, nil
}

func (u *MiddlewareUsecase) RbacAuthorization(c echo.Context, cfg *config.Config, expected []int) (echo.Context, error) {
	ctx := c.Request().Context()

	playerRoleCode := c.Get("role_code").(int)

	roleCount, err := u.middlewareRepository.RoleCount(ctx, cfg.Grpc.AuthUrl)
	if err != nil {
		return nil, err
	}

	playerRoleBinary := rbac.IntToBinary(playerRoleCode, int(roleCount))
	for i := 0; i< int(roleCount); i++ {
		if playerRoleBinary[i] & expected[i] == 1{
			return c, nil
		}
	} 

	return nil, errors.New("error: Permission denided")
}

func (u *MiddlewareUsecase) PlayerIdParamValidation(c echo.Context) (echo.Context, error) {
	playerIdRequest := c.Param("player_id")
	playerIdToken := c.Get("player_id").(string)

	if playerIdToken != "" {
		return nil, errors.New("error: player id is required")
	}

	if playerIdToken != playerIdRequest {
		return nil, errors.New("error: player id not match")
	}
	return c, nil
}