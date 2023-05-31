package storage

type IStorage[Entity any] interface {
	GetAll() ([]Entity, error)
	Create(Entity) error
	Delete(Entity) error
}
