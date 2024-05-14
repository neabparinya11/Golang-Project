package paymentusecase

import paymentrepository "github.com/neabparinya11/Golang-Project/modules/payment/paymentRepository"

type (
	PaymentUsecaseService interface{}

	PaymentUsecase struct {
		paymentRepository paymentrepository.PaymentRepositoryService
	}
)

func NewPaymentUsecase(paymentRepository paymentrepository.PaymentRepositoryService) PaymentUsecaseService{
	return &PaymentUsecase{paymentRepository: paymentRepository}
}