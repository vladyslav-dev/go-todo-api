package main

import (
	"log"

	"github.com/vladyslav-dev/go-todo-api"
	"github.com/vladyslav-dev/go-todo-api/pkg/handler"
)

func main() {
	srv := new(todo.Server)
	handlers := new(handler.Handler)
	intilizedHandlers := handlers.InitRoutes()
	if err := srv.Run("8080", intilizedHandlers); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
