package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xoscar/go-todo/config"
	"github.com/xoscar/go-todo/connectors"
	"github.com/xoscar/go-todo/middlewares"
	"github.com/xoscar/go-todo/repositories"
	"github.com/xoscar/go-todo/routers"
)

type App struct {
	config       *config.Config
	postgres     *connectors.Postgres
	repositories *repositories.Repositories
}

func New(config *config.Config, postgres *connectors.Postgres) App {
	repositories := repositories.GetRepositories(postgres.DB)
	return App{config: config, repositories: &repositories, postgres: postgres}
}

func (app App) Start() {
	router := mux.NewRouter()
	router.Use(middlewares.CommonMiddleware.GetMiddleware)
	routers.TodoRouter.GetRouter(router, app.repositories)

	http.Handle("/", router)

	log.Printf("Listening on port %d", app.config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", app.config.Port), nil))
}
