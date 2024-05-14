package server

import (
	paymenthandler "github.com/neabparinya11/Golang-Project/modules/payment/paymentHandler"
	paymentrepository "github.com/neabparinya11/Golang-Project/modules/payment/paymentRepository"
	paymentusecase "github.com/neabparinya11/Golang-Project/modules/payment/paymentUsecase"
)

func (s *Server) PaymentService() {
	repository := paymentrepository.NewPaymentRepository(s.db)
	usecase := paymentusecase.NewPaymentUsecase(repository)
	httpHandler := paymenthandler.NewPaymentHttpHandler(s.cfg, usecase)
	queueHandler := paymenthandler.NewPaymentQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = queueHandler

	//Route path
	payment := s.app.Group("/payment_v1")

	//Health check
	payment.GET("/health", s.HealthCheckService)
}