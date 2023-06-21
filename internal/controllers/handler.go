package controllers

import (
	"context"
	"net/http"
	"strconv"
	"todo-api/internal/models"

	"github.com/gin-gonic/gin"
)

type Todo interface {
	Create(ctx context.Context, todo models.Todo) error
	GetAll(ctx context.Context) ([]models.Todo, error)
	GetByID(ctx context.Context, id int) (models.Todo, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, todo models.Todo) error
}

type Handler struct {
	todoService Todo
}

func NewHandler(todo Todo) *Handler {
	return &Handler{
		todoService: todo,
	}
}

func (h *Handler) InitRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/todo", h.CreateTodo)
	r.GET("/todos", h.getAllTodo)
	r.GET("/todo/:id", h.getTodoByID)
	r.DELETE("/todo/:id", h.DeleteTodo)
	r.PUT("/todo/:id", h.UpdateTodo)

	return r
}

func (h *Handler) CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.todoService.Create(context.TODO(), todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, "OK")
}

func (h *Handler) getAllTodo(c *gin.Context) {
	todos, err := h.todoService.GetAll(context.TODO())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, todos)
}

func (h *Handler) getTodoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid id param")
		return
	}

	todo, err := h.todoService.GetByID(context.TODO(), id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, todo)
}

func (h *Handler) DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.todoService.Delete(context.TODO(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "OK")
}

func (h *Handler) UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid id param")
		return
	}
	var todo models.Todo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.todoService.Update(context.TODO(), id, todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, "OK")
}
