package auth

import (
	"time"

	"github.com/neabparinya11/Golang-Project/modules/player"
)

type (
	PlayerLoginRequest struct {
		Email    string `json:"email" form:"email" validate:"required,email,max=255"`
		Password string `json:"password" form:"password" validate:"required,max=32"`
	}

	PlayerRefreshToken struct {
		CredentialId string `json:"credential_id" form:"credential_id" validate:"required,max=64"`
		RefreshToken string `json:"refresh_token" form:"refresh_token" validate:"required,max=500"`
	}

	InsertPlayerRole struct {
		PlayerId string `json:"player_id" validate:"required"`
		RoleCode []int  `json:"role_id" validate:"required"`
	}

	ProfileIntercepter struct {
		*player.PlayerProfile
		Credential *CredentialResponse `json:"credential"`
	}

	CredentialResponse struct {
		Id           string    `json:"_id"`
		PlayerId     string    `json:"player_id"`
		RoleCode     int       `json:"role_code"`
		AccessToken  string    `json:"access_token"`
		RefreshToken string    `json:"refresh_token"`
		CreateAt     time.Time `json:"create_at"`
		UpdateAt     time.Time `json:"update_at"`
	}

	LogoutRequest struct {
		CredentialId string `json:"credential_id" form:"credential_id" validate:"required,max=64"`
	}
)
