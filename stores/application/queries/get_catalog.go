package queries

import (
	"context"

	productDomain "github.com/vladesco/ecommerce-microservices/stores/domain/product"
)

type GetCatalogParams struct {
	storeId string
}

type GetCatalogHandler struct {
	productRepository productDomain.ProductRepository
}

func NewGetCatalogHandler(productRepository productDomain.ProductRepository) GetCatalogHandler {
	return GetCatalogHandler{
		productRepository,
	}
}

func (handler GetCatalogHandler) GetCatalog(ctx context.Context, params GetCatalogParams) ([]*productDomain.Product, error) {
	return handler.productRepository.GetCatalog(ctx, params.storeId)
}
