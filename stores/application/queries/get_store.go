package queries

import (
	"context"

	storeDomain "github.com/vladesco/ecommerce-microservices/stores/domain/store"
)

type GetStoreParams struct {
	storeId string
}

type GetStoreHandler struct {
	storeRepository storeDomain.StoreRepository
}

func NewGetStoreHandler(storeRepository storeDomain.StoreRepository) GetStoreHandler {
	return GetStoreHandler{
		storeRepository,
	}
}

func (handler GetStoreHandler) GetStore(ctx context.Context, params GetStoreParams) (*storeDomain.Store, error) {
	return handler.storeRepository.FindOne(ctx, params.storeId)
}
