package paymentrepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/neabparinya11/Golang-Project/config"
	itemPb "github.com/neabparinya11/Golang-Project/modules/item/itemPb"
	"github.com/neabparinya11/Golang-Project/modules/models"
	"github.com/neabparinya11/Golang-Project/modules/player"
	"github.com/neabparinya11/Golang-Project/pkg/grpccon"
	"github.com/neabparinya11/Golang-Project/pkg/jwtauth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	PaymentRepositoryService interface{
		FindItemInIds(pctx context.Context, grpcUrl string, req *itemPb.FindItemInIdsRequest) (*itemPb.FindItemInIdsResponse , error)
		GetOffset(pctx context.Context) (int64, error)
		UpserOffset(pctx context.Context, kafkaOffset int64) error
		DockedPlayerMoney(pctx context.Context, cfg *config.Config, req *player.CreatePlayerTransactionRequest) error
		RollbackDockedPlayerMoney(pctx context.Context, cfg *config.Config, req *player.CreatePlayerTransactionRequest) error
	}

	PaymentRepository struct {
		db *mongo.Client
	}
)

func NewPaymentRepository(db *mongo.Client) PaymentRepositoryService{
	return &PaymentRepository{db: db}
}

func (pay *PaymentRepository) paymentDbConn(pctx context.Context) *mongo.Database{
	_ = pctx
	return pay.db.Database("payment_db")
}

func (r *PaymentRepository) FindItemInIds(pctx context.Context, grpcUrl string, req *itemPb.FindItemInIdsRequest) (*itemPb.FindItemInIdsResponse , error) {
	ctx, cancel := context.WithTimeout(pctx, 30*time.Second)
	defer cancel()

	jwtauth.SetApiKeyInContext(&ctx)
	conn, err := grpccon.NewGrpcClient(grpcUrl)
	if err != nil {
		return nil, errors.New("error: gRPC connected failed")
	}

	result, err := conn.Item().FindItemInIds(ctx, req)
	if err != nil {
		return nil, errors.New("error: Item not found")
	}

	if result == nil {
		return nil, errors.New("error: Item not found")
	}

	if len(result.Items) == 0 {
		return nil, errors.New("error: Item not found")
	}

	return result, nil
}

func (r *PaymentRepository) GetOffset(pctx context.Context) (int64, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.paymentDbConn(ctx)
	col := db.Collection("payment_queue")

	result := new(models.KafkaOffset)
	if err := col.FindOne(ctx, bson.M{}).Decode(result); err != nil {
		return -1, errors.New("error")
	}

	return result.Offset, nil
}

func (r *PaymentRepository) UpserOffset(pctx context.Context, kafkaOffset int64) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.paymentDbConn(ctx)
	col := db.Collection("payment_queue")

	result, err := col.UpdateOne(ctx, bson.M{}, bson.M{"$set": bson.M{"offset": kafkaOffset}})
	if err != nil {
		return errors.New("error: Upseroffset failed")
	}

	log.Printf("Info: Upseroffset result: %v", result)

	return nil
}