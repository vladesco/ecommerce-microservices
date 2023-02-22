package commands

import (
	"context"

	"github.com/vladesco/ecommerce-microservices/internal/ddd"
	productDomain "github.com/vladesco/ecommerce-microservices/stores/domain/product"
	storeDomain "github.com/vladesco/ecommerce-microservices/stores/domain/store"
)

type AddProductParams struct {
	Id          string
	StoreId     string
	Name        string
	Description string
	SKU         string
	Price       float64
}

type AddProductHandler struct {
	storeRepository   storeDomain.StoreRepository
	productRepository productDomain.ProductRepository
	eventPublisher    ddd.EventPublisher
}

func NewAddProductHandler(
	storeRepository storeDomain.StoreRepository,
	productRepository productDomain.ProductRepository,
	eventPublisher ddd.EventPublisher,
) AddProductHandler {
	return AddProductHandler{
		storeRepository,
		productRepository,
		eventPublisher,
	}
}

func (handler AddProductHandler) AddProduct(ctx context.Context, params AddProductParams) error {
	if _, err := handler.storeRepository.FindOne(ctx, params.Id); err != nil {
		return err
	}

	product, err := productDomain.CreateProduct(params.Id, params.StoreId, params.Name, params.Description, params.SKU, params.Price)

	if err != nil {
		return err
	}

	if err = handler.productRepository.Save(ctx, product); err != nil {
		return err

	}

	return handler.eventPublisher.Publish(ctx, product.GetEvents()...)
}
