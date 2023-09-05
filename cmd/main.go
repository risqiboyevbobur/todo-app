package main

import (
	"log"
	"fmt"

	"github.com/risqiboyevbobur/todo-app.git"
	"github.com/risqiboyevbobur/todo-app.git/pkg/handler"
)

func main()  {
	handlers := new(handler.Handler)
	srv:= new(todo.Server)
	if err := srv.Run("8080",handlers.InitRoutes()); err != nil{
		log.Fatalf("error while running: %s", err.Error())
	}
	fmt.Println("Hello")
}