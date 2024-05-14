package auth

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Credential struct {
		Id           primitive.ObjectID `bson:"_id,omitempty"`
		PlayerId     string             `bson:"player_id"`
		RoleCode     int                `bson:"role_code"`
		AccessToken  string             `bson:"access_token"`
		RefreshToken string             `bson:"refresh_token"`
		CreateAt     time.Time          `bson:"create_at"`
		UpdateAt     time.Time          `bson:"update_at"`
	}

	Role struct {
		Id    primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
		Title string             `json:"title" bson:"title"`
		Code  int                `json:"code" bson:"code"`
	}

	UpdateRefreshTokenRequest struct {
		PlayerId     string    `bson:"player_id"`
		AccessToken  string    `bson:"access_token"`
		RefreshToken string    `bson:"refresh_token"`
		UpdateAt     time.Time `bson:"update_at"`
	}
)
