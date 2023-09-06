package main

import (
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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
	if err:= godotenv.Load(); err != nil {
		log.Fatalf("error loading env varible :%s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:      viper.GetString("db.port"),
		Username:  viper.GetString("db.username"),
		Password:  viper.GetString("db.password"),
		DBName:    viper.GetString("db.dbname"),
		SSLMode:   viper.GetString("db.sslmode"),
	})
	if err != nil {
		// log.Fatalf("failed to initiliazed db:%s", err.Error())
	}
	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8080"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
