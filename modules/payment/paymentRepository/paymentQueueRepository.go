package paymentrepository

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/neabparinya11/Golang-Project/config"
	"github.com/neabparinya11/Golang-Project/modules/player"
	"github.com/neabparinya11/Golang-Project/pkg/queue"
)

func (r *PaymentRepository) DockedPlayerMoney(pctx context.Context, cfg *config.Config, req *player.CreatePlayerTransactionRequest) error {
	reqInByte, err := json.Marshal(req)
	if err != nil {
		return errors.New("error: docked player money failed")
	}

	if err := queue.PushMessageWithKeyToQueue([]string{cfg.Kafka.Url}, cfg.Kafka.ApiKey, cfg.Kafka.Secret, "player", "buy", reqInByte); err != nil {
		return errors.New("error: docked player money failed")
	}
	
	return nil
}

func (r *PaymentRepository) RollbackDockedPlayerMoney(pctx context.Context, cfg *config.Config, req *player.CreatePlayerTransactionRequest) error {
	reqInByte, err := json.Marshal(req)
	if err != nil {
		return errors.New("error: docked player money failed")
	}

	if err := queue.PushMessageWithKeyToQueue([]string{cfg.Kafka.Url}, cfg.Kafka.ApiKey, cfg.Kafka.Secret, "player", "rtransaction", reqInByte); err != nil {
		return errors.New("error: docked player money failed")
	}
	
	return nil
}