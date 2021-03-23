package controllers

import (
	"encoding/json"
	"go-todo/config"
	"go-todo/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetTodo(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo

	rows, err := config.DBClient.Query("SELECT * FROM todos")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var Id int
		var Description string
		rows.Scan(&Id, &Description)

		todos = append(todos, models.Todo{
			Id:          Id,
			Description: Description,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	todo := &models.Todo{}

	err := json.NewDecoder(r.Body).Decode(todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := 0
	err = config.DBClient.QueryRow("INSERT INTO todos (description) VALUES ($1) RETURNING id", todo.Description).Scan(&id)
	if err != nil {
		panic(err)
	}

	todo.Id = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.Todo
	r.ParseForm()

	err := config.DBClient.QueryRow("UPDATE todos SET description = $1 WHERE id = $2 RETURNING *", r.FormValue("description"), params["id"]).Scan(&todo.Id, &todo.Description)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&todo)

}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.Todo

	err := config.DBClient.QueryRow("DELETE FROM todos WHERE id = $1 RETURNING *", params["id"]).Scan(&todo.Id, &todo.Description)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&todo)
}
