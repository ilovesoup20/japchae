package controllers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ilovesoup20/japchae/ent"
	"github.com/mitchellh/mapstructure"
)

// Todo .
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// TodoController .
type TodoController struct {
	Client *ent.Client
}

// NewTodoController .
func NewTodoController(client *ent.Client) *TodoController {
	return &TodoController{Client: client}
}

// GetByID .
func (c *TodoController) GetByID(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")

	entTodo, err := c.Client.Todo.Get(context.Background(), id)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).SendString("Todo not found")
	}

	var todo Todo
	if err := mapstructure.Decode(entTodo, &todo); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Mapping failed")
	}
	return ctx.JSON(todo)
}

var todos = []Todo{
	{ID: 1, Title: "Buy groceries", Done: false},
	{ID: 2, Title: "Exercise", Done: false},
}

// ListTodo .
func ListTodo(c *fiber.Ctx) error {
	fmt.Println("1" == "1")
	return c.JSON(todos)
}

// GetTodoByID .
func GetTodoByID(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	for _, todo := range todos {
		if id == todo.ID {
			return c.JSON(todo)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "todo not found",
	})
}

// CreateTodo .
func CreateTodo(c *fiber.Ctx) error {
	var todo Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	todo.ID = len(todos) + 1
	todos = append(todos, todo)
	return c.Status(fiber.StatusCreated).JSON(todo)
}

// UpdateTodo .
func UpdateTodo(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var updatedTodo Todo
	if err := c.BodyParser(&updatedTodo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	for i, todo := range todos {
		if id == todo.ID {
			todos[i] = updatedTodo
			return c.JSON(updatedTodo)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Task not found",
	})
}

// PatchTodo .
func PatchTodo(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var patchedTodo Todo

	if err := c.BodyParser(&patchedTodo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	fmt.Println(id)

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Todo not found",
	})
}

// DeleteTodoByID .
func DeleteTodoByID(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	for i, todo := range todos {
		if id == todo.ID {
			todos = append(todos[:i], todos[i+1:]...)
			return c.Status(fiber.StatusNoContent).Send(nil)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Todo not found",
	})
}
