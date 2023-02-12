package ddd

type Entity interface {
	GetId() string
}

type BaseEntity struct {
	Id string
}

func (entity BaseEntity) GetId() string {
	return entity.Id
}
