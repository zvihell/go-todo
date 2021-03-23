package main

import (
	"go-todo/config"
	"go-todo/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.InitDB()
	initRouter()

}

func initRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/todos/", controllers.GetTodo).Methods("GET")
	r.HandleFunc("/todos/", controllers.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", controllers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", controllers.DeleteTodo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
