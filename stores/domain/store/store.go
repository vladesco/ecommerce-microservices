package store

import (
	"errors"

	"github.com/vladesco/ecommerce-microservices/internal/ddd"
)

var (
	ErrStoreMissingName                  = errors.New("store name can't be empty")
	ErrStoreMissingLocation              = errors.New("store location can't be empty")
	ErrStoreParticipationAlreadyEnabled  = errors.New("store participation already enabled")
	ErrStoreParticipationAlreadyDisabled = errors.New("store participation already disabled")
)

type Store struct {
	ddd.BaseAggregate
	Name                 string
	Location             string
	ParticipationEnabled bool
}

func CreateStore(id, name, location string) (store *Store, err error) {
	if name == "" {
		return nil, ErrStoreMissingName
	}

	if location == "" {
		return nil, ErrStoreMissingLocation
	}

	store = &Store{
		BaseAggregate: ddd.BaseAggregate{
			Id: id,
		},
		Name:     name,
		Location: location,
	}

	store.AddEvent(&StoreCreated{Store: store})

	return
}

func (store *Store) EnableParticipation() (err error) {
	if store.ParticipationEnabled {
		return ErrStoreParticipationAlreadyEnabled
	}

	store.ParticipationEnabled = true
	store.AddEvent(&StoreParticipationEnabled{Store: store})

	return
}

func (store *Store) DisableParticipation() (err error) {
	if !store.ParticipationEnabled {
		return ErrStoreParticipationAlreadyDisabled
	}

	store.ParticipationEnabled = false
	store.AddEvent(&StoreParticipationDisabled{Store: store})

	return
}
