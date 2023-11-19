package interfaces

type RepositoryPostgreSQL[T any] interface {
	Create(T) (int, error)
	FindById(id int) (*T, error)
	FindAll(filter T) ([]T, error)
	Delete(id int) error
	Update(filter T) error
}
