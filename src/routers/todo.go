package routers

import (
	"github.com/gorilla/mux"
	"github.com/xoscar/go-todo/controllers"
	"github.com/xoscar/go-todo/repositories"
)

var TodoRouter Router = todoRouter{}

type todoRouter struct{}

func (t todoRouter) GetRouter(router *mux.Router, repositories *repositories.Repositories) {
	controller := controllers.NewTodoController(repositories)

	router.HandleFunc("/todos", controller.Create).Methods("POST")
	router.HandleFunc("/todos", controller.Update).Methods("PUT")
	router.HandleFunc("/todos/{id}", controller.Delete).Methods("DELETE")
	router.HandleFunc("/todos", controller.GetAll).Methods("GET")
	router.HandleFunc("/todos/{id}", controller.Get).Methods("GET")
}
