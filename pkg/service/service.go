package service

import "github.com/vladyslav-dev/go-todo-api/pkg/repository"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func newService(repos *repository.Repository) *Service {
	return &Service{}
}
