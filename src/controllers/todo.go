package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/xoscar/go-todo/models"
	"github.com/xoscar/go-todo/repositories"
)

var TodoController Controller = todoController{}

type todoController struct {
	todoRepository *repositories.TodoRepository
}

func NewTodoController(repositories *repositories.Repositories) todoController {
	return todoController{todoRepository: &repositories.TodoRepository}
}

func (t todoController) GetAll(w http.ResponseWriter, r *http.Request) {
	todoList, error := t.todoRepository.GetAll()

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(todoList)
}

func (t todoController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	id, _ := strconv.Atoi(key)

	todo, err := t.todoRepository.Get(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func (t todoController) Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)

	todo, error := t.todoRepository.Create(todo)

	if error != nil {
		log.Println(error)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func (t todoController) Update(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)

	todo, error := t.todoRepository.Update(todo)

	if error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func (t todoController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	id, _ := strconv.Atoi(key)

	err := t.todoRepository.Delete(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
