package commands

import (
	"context"
	"fmt"

	"github.com/vladesco/ecommerce-microservices/internal/ddd"
	storeDomain "github.com/vladesco/ecommerce-microservices/stores/domain/store"
)

type CreateStoreParams struct {
	Id       string
	Name     string
	Location string
}

type CreateStoreHandler struct {
	storeRepository storeDomain.StoreRepository
	eventPublisher  ddd.EventPublisher
}

func NewCreateStoreHandler(store storeDomain.StoreRepository, eventPublisher ddd.EventPublisher) CreateStoreHandler {
	return CreateStoreHandler{
		store,
		eventPublisher,
	}
}

func (handler CreateStoreHandler) CreateStore(ctx context.Context, params CreateStoreParams) error {
	store, err := storeDomain.CreateStore(params.Id, params.Name, params.Location)

	if err != nil {
		return fmt.Errorf("create store error while creating store %w", err)
	}

	if err = handler.storeRepository.Save(ctx, store); err != nil {
		return fmt.Errorf("create store error while saving store %w", err)
	}

	if err = handler.eventPublisher.Publish(ctx, store.Events...); err != nil {
		return fmt.Errorf("create store error while publishing events %w", err)
	}

	return nil
}
