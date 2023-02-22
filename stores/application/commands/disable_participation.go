package commands

import (
	"context"

	"github.com/vladesco/ecommerce-microservices/internal/ddd"
	storeDomain "github.com/vladesco/ecommerce-microservices/stores/domain/store"
)

type DisableParticipationParams struct {
	storeId string
}

type DisableParticipationHandler struct {
	storeRepository storeDomain.StoreRepository
	eventPublisher  ddd.EventPublisher
}

func NewDisableParticipationHandler(storeRepository storeDomain.StoreRepository, eventPublisher ddd.EventPublisher) DisableParticipationHandler {
	return DisableParticipationHandler{
		storeRepository,
		eventPublisher,
	}
}

func (handler DisableParticipationHandler) DisableParticipation(ctx context.Context, params DisableParticipationParams) error {
	store, err := handler.storeRepository.FindOne(ctx, params.storeId)

	if err != nil {
		return err
	}

	if err = store.DisableParticipation(); err != nil {
		return err
	}

	if err = handler.storeRepository.Update(ctx, store); err != nil {
		return err
	}

	return handler.eventPublisher.Publish(ctx, store.GetEvents()...)
}
