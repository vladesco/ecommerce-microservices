package product

import (
	"context"
)

type ProductRepository interface {
	Save(ctx context.Context, product *Product) error
	Delete(ctx context.Context, productId string) error
	FindOne(ctx context.Context, productId string) (*Product, error)
	GetCatalog(ctx context.Context) ([]*Product, error)
}
