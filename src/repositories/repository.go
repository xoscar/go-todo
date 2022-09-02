package repositories

import "database/sql"

type Repository[T any] interface {
	GetAll() ([]T, error)
	Get(id int) (T, error)
	Create(todo T) (T, error)
	Update(todo T) (T, error)
	Delete(id int) error
}

type Repositories struct {
	TodoRepository TodoRepository
}

func GetRepositories(db *sql.DB) Repositories {
	return Repositories{
		TodoRepository: NewTodoRepository(db),
	}
}
