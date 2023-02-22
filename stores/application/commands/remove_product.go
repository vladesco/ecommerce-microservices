package commands

import (
	"context"

	"github.com/vladesco/ecommerce-microservices/internal/ddd"
	productDomain "github.com/vladesco/ecommerce-microservices/stores/domain/product"
)

type RemoveProductParams struct {
	Id string
}

type RemoveProductHandler struct {
	productRepository productDomain.ProductRepository
	eventPublisher    ddd.EventPublisher
}

func NewRemoveProductHandler(
	productRepository productDomain.ProductRepository,
	eventPublisher ddd.EventPublisher,
) RemoveProductHandler {
	return RemoveProductHandler{
		productRepository,
		eventPublisher,
	}
}

func (handler RemoveProductHandler) RemoveProduct(ctx context.Context, params RemoveProductParams) error {
	product, err := handler.productRepository.FindOne(ctx, params.Id)

	if err != nil {
		return err
	}

	if err = product.Remove(); err != nil {
		return err
	}

	if err = handler.productRepository.Delete(ctx, product.Id); err != nil {
		return err
	}

	return handler.eventPublisher.Publish(ctx, product.GetEvents()...)
}
