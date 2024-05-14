package paymenthandler

import (
	"github.com/neabparinya11/Golang-Project/config"
	paymentusecase "github.com/neabparinya11/Golang-Project/modules/payment/paymentUsecase"
)

type (
	PaymentQueueHandlerService interface{}

	PaymentQueueHandler struct {
		cfg            *config.Config
		paymentUsecase paymentusecase.PaymentUsecaseService
	}
)

func NewPaymentQueueHandler(cfg *config.Config, paymentUsecase paymentusecase.PaymentUsecaseService) PaymentQueueHandlerService{
	return &PaymentQueueHandler{cfg: cfg, paymentUsecase: paymentUsecase}
}