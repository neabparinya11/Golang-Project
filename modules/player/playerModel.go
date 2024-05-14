package player

import "time"

type (
	PlayerProfile struct {
		Id       string    `json:"id"`
		Email    string    `json:"email"`
		Username string    `json:"username"`
		CreateAt time.Time `json:"create_at"`
		UpdateAt time.Time `json:"update_at"`
	}

	PlayerClaims struct {
		Id       string `json:"id"`
		RoleCode int    `json:"role_code"`
	}

	CreatePlayerRequest struct {
		Email    string `json:"email" form:"email" validate:"required,email,max=255"`
		Password string `json:"password" form:"password" validate:"required,max=32"`
		Username string `json:"username" form:"username" validate:"required,max=64"`
	}

	CreatePlayerTransactionRequest struct {
		PlayerId string  `json:"player_id" validate:"required,max=64"`
		Amount   float64 `json:"amount" validate:"required"`
	}
	
)
