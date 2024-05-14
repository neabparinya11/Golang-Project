package middlewareusecase

import middlewarerepository "github.com/neabparinya11/Golang-Project/modules/middleware/middlewareRepository"

type (
	MiddlewareUsecaseService interface{}

	MiddlewareUsecase struct {
		middlewareRepository middlewarerepository.MiddlewareRepositoryService
	}
)

func NewMiddlewareUsecase(middlewareRepository middlewarerepository.MiddlewareRepositoryService) MiddlewareUsecaseService{
	return &MiddlewareUsecase{middlewareRepository: middlewareRepository}
}