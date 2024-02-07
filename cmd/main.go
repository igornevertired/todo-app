package main

import (
	"log"

	"github.com/igornevertired/todo-app"
	"github.com/igornevertired/todo-app/pkg/handler"
	"github.com/igornevertired/todo-app/pkg/repository"
	"github.com/igornevertired/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf(err.Error())
	}
}
