package controllers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ilovesoup20/japchae/ent"
	"github.com/mitchellh/mapstructure"
)

// TodoController .
type TodoController struct {
	Client *ent.TodoClient
}

// NewTodoController .
func NewTodoController(client *ent.TodoClient) *TodoController {
	return &TodoController{Client: client}
}

// ListTodos .
func (c *TodoController) ListTodos(fctx *fiber.Ctx) error {
	ctx := context.Background()
	dbTodos, err := c.Client.Query().All(ctx)

	if err != nil {
		return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprint(err),
		})
	}

	todos := make([]ent.Todo, len(dbTodos))

	for i, todo := range dbTodos {
		err := mapstructure.Decode(todo, &todos[i])
		if err != nil {
			fmt.Println("mapping error")
		}
	}
	return fctx.JSON(todos)
}

// GetTodoByID .
func (c *TodoController) GetTodoByID(fiberCtx *fiber.Ctx) error {
	id, _ := fiberCtx.ParamsInt("id")

	ctx := context.Background()

	entTodo, err := c.Client.Get(ctx, id)

	c.Client.Create().Save(ctx)

	if err != nil {
		return fiberCtx.Status(fiber.StatusNotFound).SendString("Todo not found")
	}

	var todo ent.Todo
	if err := mapstructure.Decode(entTodo, &todo); err != nil {
		return fiberCtx.Status(fiber.StatusInternalServerError).SendString("Mapping failed")
	}
	return fiberCtx.JSON(todo)
}

// CreateTodo .
func (c *TodoController) CreateTodo(fiberCtx *fiber.Ctx) error {
	var todo ent.Todo
	if err := fiberCtx.BodyParser(&todo); err != nil {
		return fiberCtx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	fmt.Println(todo)
	ctx := context.Background()
	dbTodo, err := c.Client.
		Create().
		SetTitle(todo.Title).
		Save(ctx)

	if err != nil {
		fmt.Println("error?")
	}

	if err := mapstructure.Decode(dbTodo, &todo); err != nil {
		return fiberCtx.Status(fiber.StatusInternalServerError).SendString("mapping failed")
	}
	fmt.Println(dbTodo)
	return fiberCtx.JSON(dbTodo)
}

/*
 *	OLDDDDDDD
 */
var todos = []ent.Todo{
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
	var todo ent.Todo
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
	var updatedTodo ent.Todo
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
	var patchedTodo ent.Todo

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
