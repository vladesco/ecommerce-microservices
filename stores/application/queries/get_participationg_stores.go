package queries

import (
	"context"

	storeDomain "github.com/vladesco/ecommerce-microservices/stores/domain/store"
)

type GetParticipatingStoresHandler struct {
	participationStoreRepository storeDomain.ParticipationStoreRepository
}

func NewGetParticipatingStoresHandler(participationStoreRepository storeDomain.ParticipationStoreRepository) GetParticipatingStoresHandler {
	return GetParticipatingStoresHandler{
		participationStoreRepository,
	}
}

func (handler GetParticipatingStoresHandler) GetParticipatingStores(ctx context.Context) ([]*storeDomain.Store, error) {
	return handler.participationStoreRepository.FindAll(ctx)
}
