package main

import (
	"flag"
	"log"

	"github.com/xoscar/go-todo/app"
	"github.com/xoscar/go-todo/config"
	"github.com/xoscar/go-todo/connectors"
)

var configFlag = flag.String("config", "./config.yaml", "path to the config file")

func main() {
	flag.Parse()

	config, err := config.FromFile(*configFlag)
	if err != nil {
		log.Fatal(err)
	}

	postgres, err := connectors.NewPostgres(config.PostgresConnString, config.MigrationsFolder)

	if err != nil {
		log.Fatal(err)
	}

	app := app.New(&config, &postgres)
	app.Start()
}
