package commands

import (
	"context"

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
		return err
	}

	if err = handler.storeRepository.Save(ctx, store); err != nil {
		return err
	}

	return handler.eventPublisher.Publish(ctx, store.Events...)
}
