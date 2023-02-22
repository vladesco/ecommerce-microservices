package queries

import (
	"context"

	storeDomain "github.com/vladesco/ecommerce-microservices/stores/domain/store"
)

type GetStoresHandler struct {
	storeRepository storeDomain.StoreRepository
}

func NewGetStoresHandler(storeRepository storeDomain.StoreRepository) GetStoreHandler {
	return GetStoreHandler{
		storeRepository,
	}
}

func (handler GetStoresHandler) GetStores(ctx context.Context) ([]*storeDomain.Store, error) {
	return handler.storeRepository.FindAll(ctx)
}
