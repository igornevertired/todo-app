package main

import (
	"log"

	"github.com/igornevertired/todo-app"
	"github.com/igornevertired/todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf(err.Error())
	}
}
