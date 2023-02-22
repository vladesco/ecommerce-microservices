package commands

import (
	"context"
	"fmt"

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
		return fmt.Errorf("enable participation error while finding store %w", err)
	}

	if err = store.EnableParticipation(); err != nil {
		return fmt.Errorf("enable participation error while updating store %w", err)
	}

	if err = handler.storeRepository.Update(ctx, store); err != nil {
		return fmt.Errorf("enable participation error while saving updated store %w", err)
	}

	if err = handler.eventPublisher.Publish(ctx, store.GetEvents()...); err != nil {
		return fmt.Errorf("enable participation error while publishing events %w", err)
	}

	return nil
}
