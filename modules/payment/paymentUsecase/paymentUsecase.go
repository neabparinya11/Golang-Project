package paymentusecase

import (
	"context"
	"errors"
	"log"

	"github.com/IBM/sarama"
	"github.com/neabparinya11/Golang-Project/config"
	"github.com/neabparinya11/Golang-Project/modules/item"
	itemPb "github.com/neabparinya11/Golang-Project/modules/item/itemPb"
	"github.com/neabparinya11/Golang-Project/modules/payment"
	paymentrepository "github.com/neabparinya11/Golang-Project/modules/payment/paymentRepository"
	"github.com/neabparinya11/Golang-Project/pkg/queue"
)

type (
	PaymentUsecaseService interface{
		FindItemInIds(pctx context.Context, grpcUrls string, req []*payment.ItemServiceRequestDatum) error
		GetOffset(pctx context.Context) (int64, error)
		UpserOffset(pctx context.Context, kafkaOffset int64) error
		BuyItem(pctx context.Context, cfg *config.Config, playerId string, req *payment.ItemServiceRequest) (*payment.PaymentTranferResponse, error)
		SellItem(pctx context.Context, cfg *config.Config, playerId string, req *payment.ItemServiceRequest) (*payment.PaymentTranferResponse, error)
	}

	PaymentUsecase struct {
		paymentRepository paymentrepository.PaymentRepositoryService
	}
)

func NewPaymentUsecase(paymentRepository paymentrepository.PaymentRepositoryService) PaymentUsecaseService{
	return &PaymentUsecase{paymentRepository: paymentRepository}
}

func (u *PaymentUsecase) FindItemInIds(pctx context.Context, grpcUrls string, req []*payment.ItemServiceRequestDatum) error {

	setIds := make(map[string]bool)
	for _, v := range req {
		if !setIds[v.ItemId]{
			setIds[v.ItemId] = true
		}
	}

	itemData, err := u.paymentRepository.FindItemInIds(pctx, grpcUrls, &itemPb.FindItemInIdsRequest{
		Ids: func() []string {
			itemIds := make([]string, 0)
			for k, _ := range setIds {
				itemIds = append(itemIds, k)
			}
			return itemIds
		}(),
	})
	if err != nil {
		return errors.New("error: Item not found")
	}

	itemMap := make(map[string]*item.ItemShowCase)
	for _, v := range itemData.Items {
		itemMap[v.Id] = &item.ItemShowCase{
			ItemId: v.Id,
			Title: v.Title,
			Price: v.Price,
			Damage: int(v.Damage),
			ImageUrl: v.ImageUrl,
		}
	}

	for i := range req {
		if _, ok := itemMap[req[i].ItemId]; !ok {
			return errors.New("error: Item not found")
		}
		req[i].Price = itemMap[req[i].ItemId].Price
	}

	return nil
}

func (u *PaymentUsecase) GetOffset(pctx context.Context) (int64, error) {
	return u.paymentRepository.GetOffset(pctx)
}

func (u *PaymentUsecase) UpserOffset(pctx context.Context, kafkaOffset int64) error {
	return u.paymentRepository.UpserOffset(pctx, kafkaOffset)
}

func (u *PaymentUsecase) BuyItem(pctx context.Context, cfg *config.Config, playerId string, req *payment.ItemServiceRequest) (*payment.PaymentTranferResponse, error) {
	if err := u.FindItemInIds(pctx, cfg.Grpc.ItemUrl, req.Items); err != nil {
		return nil, err
	}
	return nil, nil
}

func (u *PaymentUsecase) SellItem(pctx context.Context, cfg *config.Config, playerId string, req *payment.ItemServiceRequest) (*payment.PaymentTranferResponse, error) {
	if err := u.FindItemInIds(pctx, cfg.Grpc.ItemUrl, req.Items); err != nil {
		return nil, err
	}
	return nil, nil
}

func (u *PaymentUsecase) BuyItemRes(pctx context.Context, cfg *config.Config, resCh chan<- payment.PaymentTranferResponse) {

}

func (u *PaymentUsecase) PaymentConsumer(pctx context.Context, cfg *config.Config) (sarama.PartitionConsumer, error) {
	worker, err := queue.ConnectConsumer([]string{cfg.Kafka.Url}, cfg.Kafka.ApiKey, cfg.Kafka.Secret)
	if err != nil {
		return nil, err
	}

	offset, err := u.GetOffset(pctx)
	if err != nil {
		return nil, err
	}

	consumer, err := worker.ConsumePartition("payment", 0, offset)
	if err != nil {
		_, err := worker.ConsumePartition("payment", 0, 0)
		if err != nil {
			return nil, err
		}
	}
	return consumer, nil
}

func (u *PaymentUsecase) BuyOrSellConsumer(pctx context.Context, cfg *config.Config, resCh chan<- *payment.PaymentTranferResponse) {
	consumer, err := u.PaymentConsumer(pctx, cfg)
	if err != nil {
		resCh <- nil
		return
	}

	log.Println("Start BuyOrSellConsumer...")

	select {
	case err:= <- consumer.Errors():
		log.Println("Error: BuyOrSellConsumer failed: ", err.Error())
		resCh <- nil
		return
	case msg := <- consumer.Messages():
		if string(msg.Key) == "buy" {
			u.UpserOffset(pctx, msg.Offset+1)

			req := new(payment.PaymentTranferResponse)

			if err := queue.DecodeMessage(req, msg.Value); err != nil {
				resCh <- nil
				return
			}

			resCh <- req
			log.Printf("BuyOrSellConsumer | Topics(%s) | offset(%d) | Message(%s) \n", msg.Topic, msg.Offset, msg.Value)
		}
	}
}