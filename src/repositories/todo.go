package repositories

import (
	"database/sql"
	"fmt"

	"github.com/xoscar/go-todo/connectors"
	"github.com/xoscar/go-todo/models"
)

var TRepository Repository[models.Todo] = TodoRepository{}

type TodoRepository struct {
	db *sql.DB
}

const getAllQuery = `
		SELECT id, title, status
		from todos
	`

func NewTodoRepository(db *sql.DB) TodoRepository {
	return TodoRepository{db: db}
}

func (repository TodoRepository) GetAll() ([]models.Todo, error) {
	rows, err := repository.db.Query(getAllQuery)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	todoList := []models.Todo{}

	for rows.Next() {
		todo, err := scanRow(rows)

		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		todoList = append(todoList, todo)
	}

	return todoList, nil
}

func (repository TodoRepository) Get(id int) (models.Todo, error) {
	query := getAllQuery + `
		WHERE id = $1
	`

	rows, err := repository.db.Query(query, id)
	if err != nil {
		return models.Todo{}, fmt.Errorf("query: %w", err)
	}

	rows.Next()
	todo, err := scanRow(rows)

	if err != nil {
		return models.Todo{}, fmt.Errorf("scan row: %w", err)
	}

	return todo, nil
}

func (repository TodoRepository) Create(todo models.Todo) (models.Todo, error) {
	insert := `
		INSERT INTO todos VALUES (default, $1, $2) RETURNING id
	`
	var id int
	err := repository.db.QueryRow(insert, todo.Title, todo.Status).Scan(&id)

	if err != nil {
		return models.Todo{}, fmt.Errorf("scan: %w", err)
	}

	todo.ID = id
	return todo, nil
}

func (repository TodoRepository) Update(todo models.Todo) (models.Todo, error) {
	oldTodo, err := repository.Get(todo.ID)

	if err != nil {
		return models.Todo{}, fmt.Errorf("get: %w", err)
	}

	oldTodo.Status = todo.Status
	oldTodo.Title = todo.Title

	update := `
		UPDATE todos
		SET title = $2, status = $3
		WHERE id = $1
	`

	_, err = repository.db.Query(update, todo.ID, todo.Title, todo.Status)
	if err != nil {
		return models.Todo{}, fmt.Errorf("query: %w", err)
	}

	oldTodo.Status = todo.Status
	oldTodo.Title = todo.Title

	return oldTodo, nil
}

func (repository TodoRepository) Delete(id int) error {
	delete := `
		DELETE FROM todos
		WHERE id = $1
	`

	_, err := repository.db.Exec(delete, id)

	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}

func scanRow(rows connectors.Scanner) (models.Todo, error) {
	var todo models.Todo

	error := rows.Scan(&todo.ID, &todo.Title, &todo.Status)
	if error != nil {
		return models.Todo{}, fmt.Errorf("scan: %w", error)
	}

	return todo, nil
}
