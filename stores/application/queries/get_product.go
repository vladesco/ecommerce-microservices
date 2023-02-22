package queries

import (
	"context"

	productDomain "github.com/vladesco/ecommerce-microservices/stores/domain/product"
)

type GetProductParams struct {
	productId string
}

type GetProductHandler struct {
	productRepository productDomain.ProductRepository
}

func NewGetProductHandlerHandler(productRepository productDomain.ProductRepository) GetProductHandler {
	return GetProductHandler{
		productRepository,
	}
}

func (handler GetProductHandler) GetProduct(ctx context.Context, params GetProductParams) (*productDomain.Product, error) {
	return handler.productRepository.FindOne(ctx, params.productId)
}
