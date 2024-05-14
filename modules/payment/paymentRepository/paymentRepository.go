package paymentrepository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	PaymentRepositoryService interface{}

	PaymentRepository struct {
		db *mongo.Client
	}
)

func NewPaymentRepository(db *mongo.Client) PaymentRepositoryService{
	return &PaymentRepository{db: db}
}

func (pay *PaymentRepository) paymentDbConn(pctx context.Context) *mongo.Database{
	return pay.db.Database("payment_db")
}