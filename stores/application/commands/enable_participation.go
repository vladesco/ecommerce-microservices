package commands

import (
	"context"

	"github.com/vladesco/ecommerce-microservices/internal/ddd"
	storeDomain "github.com/vladesco/ecommerce-microservices/stores/domain/store"
)

type EnableParticipationParams struct {
	storeId string
}

type EnableParticipationHandler struct {
	storeRepository storeDomain.StoreRepository
	eventPublisher  ddd.EventPublisher
}

func NewEnableParticipationHandler(storeRepository storeDomain.StoreRepository, eventPublisher ddd.EventPublisher) EnableParticipationHandler {
	return EnableParticipationHandler{
		storeRepository,
		eventPublisher,
	}
}

func (handler EnableParticipationHandler) EnableParticipation(ctx context.Context, params EnableParticipationParams) error {
	store, err := handler.storeRepository.FindOne(ctx, params.storeId)

	if err != nil {
		return err
	}

	if err = store.EnableParticipation(); err != nil {
		return err
	}

	if err = handler.storeRepository.Update(ctx, store); err != nil {
		return err
	}

	return handler.eventPublisher.Publish(ctx, store.GetEvents()...)
}
