package paymenthandler

import (
	"github.com/neabparinya11/Golang-Project/config"
	paymentusecase "github.com/neabparinya11/Golang-Project/modules/payment/paymentUsecase"
)

type (
	PaymentHttpHandlerService interface{}

	PaymentHttpHandler struct {
		cfg            *config.Config
		paymentUsecase paymentusecase.PaymentUsecaseService
	}
)

func NewPaymentHttpHandler(cfg *config.Config, paymentUsecase paymentusecase.PaymentUsecaseService) PaymentHttpHandlerService{
	return &PaymentHttpHandler{cfg: cfg, paymentUsecase: paymentUsecase}
}