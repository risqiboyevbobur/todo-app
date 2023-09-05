package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	"github.com/risqiboyevbobur/todo-app.git"
	"github.com/risqiboyevbobur/todo-app.git/pkg/handler"
	"github.com/risqiboyevbobur/todo-app.git/pkg/repository"
	"github.com/risqiboyevbobur/todo-app.git/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initilization: %s", err.Error())
	}
	repos := repository.NewRepository()
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8080"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running: %s", err.Error())
	}
	fmt.Println("Hello")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
