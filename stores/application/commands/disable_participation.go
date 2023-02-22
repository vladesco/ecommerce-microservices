package commands

import (
	"context"
	"fmt"

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
		return fmt.Errorf("disable participation error while finding store %w", err)
	}

	if err = store.DisableParticipation(); err != nil {
		return fmt.Errorf("disable participation error while updating store %w", err)
	}

	if err = handler.storeRepository.Update(ctx, store); err != nil {
		return fmt.Errorf("disable participation error while saving updated store %w", err)
	}

	if err = handler.eventPublisher.Publish(ctx, store.GetEvents()...); err != nil {
		return fmt.Errorf("disable participation error while publishing events %w", err)
	}

	return nil
}
