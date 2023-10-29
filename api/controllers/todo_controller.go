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

// List .
func (c *TodoController) List(fctx *fiber.Ctx) error {
	ctx := context.Background()
	dbTodos, err := c.Client.Todo.Query().All(ctx)

	if err != nil {
		return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Sprint(err),
		})
	}

	todos := make([]Todo, len(dbTodos))

	// var todo Todo
	// if err := mapstructure.Decode(entTodo, &todo); err != nil {
	// 	return fiberCtx.Status(fiber.StatusInternalServerError).SendString("Mapping failed")
	// }
	for i, todo := range dbTodos {
		err := mapstructure.Decode(todo, &todos[i])
		if err != nil {
			fmt.Println("mapping error")
		}
		// todos[i] = Todo{
		// 	ID:    todo.ID,
		// 	Title: todo.Title,
		// 	Done:  todo.Done,
		// }
	}
	return fctx.JSON(todos)
}

// GetByID .
func (c *TodoController) GetByID(fiberCtx *fiber.Ctx) error {
	id, _ := fiberCtx.ParamsInt("id")

	ctx := context.Background()

	entTodo, err := c.Client.Todo.Get(ctx, id)

	c.Client.Todo.Create().Save(ctx)

	if err != nil {
		return fiberCtx.Status(fiber.StatusNotFound).SendString("Todo not found")
	}

	var todo Todo
	if err := mapstructure.Decode(entTodo, &todo); err != nil {
		return fiberCtx.Status(fiber.StatusInternalServerError).SendString("Mapping failed")
	}
	return fiberCtx.JSON(todo)
}

// func (c *TodoController) CreateTodo(ctx *fiber.Ctx) error {
// 	todo, err := c.Client.Todo.Create().Save(ctx)

// 	if err != nil {
// 		return nil, fmt.Errorf("Failed creating todo: %w", err)
// 	}

// 	fmt.Println("New Todo was created.")
// 	return todo, nil
// }

/*
 *	OLDDDDDDD
 */
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
