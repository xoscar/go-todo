package routers

import (
	"github.com/gorilla/mux"

	"github.com/xoscar/go-todo/repositories"
)

type Router interface {
	GetRouter(router *mux.Router, repositories *repositories.Repositories)
}
