package product

import (
	"errors"

	"github.com/vladesco/ecommerce-microservices/internal/ddd"
)

var (
	ErrProductMissingLocation      = errors.New("product name can't be empty")
	ErrProductPriceLessOrEqualZero = errors.New("product can't have price less or equal to zero")
)

type Product struct {
	ddd.BaseAggregate
	StoreId     string
	Name        string
	Description string
	SKU         string
	Price       float64
}

func CreateProduct(id, storeId, name, description, sku string, price float64) (product *Product, err error) {
	if name == "" {
		return nil, ErrProductMissingLocation
	}

	if price <= 0 {
		return nil, ErrProductPriceLessOrEqualZero
	}

	product = &Product{
		BaseAggregate: ddd.BaseAggregate{
			Id: id,
		},
		StoreId:     storeId,
		Name:        name,
		Description: description,
		SKU:         sku,
		Price:       price,
	}

	product.AddEvent(&ProductAdded{Product: product})

	return
}

func (product *Product) Remove() error {
	product.AddEvent(&ProductRemoved{Product: product})

	return nil
}
