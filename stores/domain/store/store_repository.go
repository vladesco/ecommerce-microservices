package store_domain

import (
	"context"
)

type StoreRepository interface {
	Save(ctx context.Context, store *Store) error
	Update(ctx context.Context, store *Store) error
	Delete(ctx context.Context, store *Store) error
	FindOne(ctx context.Context, storeId string) (*Store, error)
	FindAll(ctx context.Context) ([]*Store, error)
}

type ParticipationStoreRepository interface {
	FindAll(ctx context.Context) ([]*Store, error)
}
