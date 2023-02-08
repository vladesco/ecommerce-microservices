package ddd

type Entity interface {
	GetId() string
}

type BaseEntity struct {
	id string
}

func (entity BaseEntity) GetId() string {
	return entity.id
}
