package main

import (
	"log"

	"github.com/vladyslav-dev/go-todo-api"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
