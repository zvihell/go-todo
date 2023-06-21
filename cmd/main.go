package main

import (
	"log"
	"net/http"
	"todo-api/internal/controllers"
	"todo-api/internal/repository"
	"todo-api/internal/service"
	"todo-api/pkg"

	"github.com/spf13/viper"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	db := pkg.InitDB(pkg.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetInt("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLmode:  viper.GetString("db.sslmode"),
	})

	repo := repository.NewTodo(db)
	services := service.NewTodo(repo)
	handler := controllers.NewHandler(services)

	s := &http.Server{
		Addr:    ":8080",
		Handler: handler.InitRouter(),
	}

	log.Println("Server started at 8080")
	s.ListenAndServe()

}
