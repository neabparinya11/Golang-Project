package middlewarerepository

import "context"

type (
	MiddlewareRepositoryService interface{}

	MiddlewareRepository struct {
	}
)

func NewMiddlewareRepository() MiddlewareRepositoryService {
	return &MiddlewareRepository{}
}

func (r *MiddlewareRepository) AccessTokenSearch(pctx context.Context, grpcUrl, accessToken string) bool {
	return true
}