package commands

import (
	"context"
	"fmt"

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
		return fmt.Errorf("remove product error while finding product %w", err)
	}

	if err = product.Remove(); err != nil {
		return fmt.Errorf("remove product error while updating product %w", err)
	}

	if err = handler.productRepository.Delete(ctx, product.Id); err != nil {
		return fmt.Errorf("remove product error while saving deleting product %w", err)
	}

	if err = handler.eventPublisher.Publish(ctx, product.GetEvents()...); err != nil {
		return fmt.Errorf("remove product error while publishing events %w", err)
	}

	return nil
}
