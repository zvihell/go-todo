package models

type Todo struct {
	Id          int    `json: "id"`
	Description string `json:"description"`
}
