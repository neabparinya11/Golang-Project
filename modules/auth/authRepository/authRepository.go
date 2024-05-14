package authrepository

import (
	"context"
	"errors"
	"time"

	"github.com/neabparinya11/Golang-Project/pkg/grpccon"
	"github.com/neabparinya11/Golang-Project/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/neabparinya11/Golang-Project/modules/auth"
	playerPb "github.com/neabparinya11/Golang-Project/modules/player/playerPb"
)

type (
	AuthRepositoryService interface {
		CredentialSearch(pctx context.Context, grpcUrl string, req *playerPb.CreadentialSearchRequest) (*playerPb.PlayerProfile, error)
		InsertOnePlayerCredential(pctx context.Context, req *auth.Credential) (primitive.ObjectID, error)
		FindOnePlayerCredential(pctx context.Context, credentialId string) (*auth.Credential, error)
		FindOnePlayerProfileToRefresh(pctx context.Context, grpcUrl string, req *playerPb.FindOnePlayerProfileToRefreshRequest) (*playerPb.PlayerProfile, error)
		UpdateOnePlayerCredential(pctx context.Context, creadentialId string, req *auth.UpdateRefreshTokenRequest) error
		DeleteOnePlayerCredential(pctx context.Context, credentialId string) (int64, error)
	}

	AuthRepository struct {
		db *mongo.Client
	}
)

func NewAuthRepository(db *mongo.Client) AuthRepositoryService {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) authDbConn(pctx context.Context) *mongo.Database {
	_ = pctx
	return r.db.Database("auth_db")
}

func (r *AuthRepository) CredentialSearch(pctx context.Context, grpcUrl string, req *playerPb.CreadentialSearchRequest) (*playerPb.PlayerProfile, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	conn, err := grpccon.NewGrpcClient(grpcUrl)
	if err != nil {
		return nil, errors.New("error: gRPC connected failed")
	}

	result, err := conn.Player().CreadentialSearch(ctx, req)
	if err != nil {
		return nil, errors.New("error: email or password is incorrect")
	}

	return result, nil
}

func (r *AuthRepository) InsertOnePlayerCredential(pctx context.Context, req *auth.Credential) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.authDbConn(ctx)
	col := db.Collection("auth")

	result, err := col.InsertOne(ctx, req)
	if err != nil {
		return primitive.NilObjectID, errors.New("error: insert one player credential failed")
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *AuthRepository) FindOnePlayerCredential(pctx context.Context, credentialId string) (*auth.Credential, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.authDbConn(ctx)
	col := db.Collection("auth")

	result := new(auth.Credential)

	if err := col.FindOne(ctx, bson.M{"credential_id": utils.ConvertToObjectId(credentialId)}).Decode(result); err != nil {
		return nil, errors.New("error: Find one player credential failed")
	}

	return result, nil
}

func (r *AuthRepository) FindOnePlayerProfileToRefresh(pctx context.Context, grpcUrl string, req *playerPb.FindOnePlayerProfileToRefreshRequest) (*playerPb.PlayerProfile, error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	conn, err := grpccon.NewGrpcClient(grpcUrl)
	if err != nil {
		return nil, errors.New("error: gRPC connected failed")
	}

	result, err := conn.Player().FindOnePlayerProfileToRefresh(ctx, req)
	if err != nil {
		return nil, errors.New("error: Player profile not founded")
	}

	return result, nil
}

func (r *AuthRepository) UpdateOnePlayerCredential(pctx context.Context, creadentialId string, req *auth.UpdateRefreshTokenRequest) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.authDbConn(ctx)
	col := db.Collection("auth")

	_, err := col.UpdateOne(
		ctx,
		bson.M{"_id": utils.ConvertToObjectId(creadentialId)},
		bson.M{"$set": bson.M{
			"player_id":     req.PlayerId,
			"access_token":  req.AccessToken,
			"refresh_token": req.RefreshToken,
			"update_at":     req.UpdateAt,
		}},
	)
	if err != nil {
		return errors.New("error: Update one player credential failed")
	}

	return nil
}

func (r *AuthRepository) DeleteOnePlayerCredential(pctx context.Context, credentialId string) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.authDbConn(ctx)
	col := db.Collection("auth")

	result, err := col.DeleteOne(pctx, bson.M{"_id": utils.ConvertToObjectId(credentialId)})
	if err != nil {
		return -1, errors.New("error: Cannot delete credential")
	}

	return result.DeletedCount, nil
}