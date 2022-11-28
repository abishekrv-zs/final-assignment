package store

type Store interface {
	Get(id string) ([]any, error)
	Create(obj any) (any, error)
	Update(obj any) (any, error)
	Delete(id string) (any, error)
}
