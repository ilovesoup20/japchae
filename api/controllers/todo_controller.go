package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos = []Todo{
	{ID: 1, Title: "Buy groceries", Done: false},
	{ID: 2, Title: "Exercise", Done: false},
}

func ListTodo(c *fiber.Ctx) error {
	fmt.Println("1" == "1")
	return c.JSON(todos)
}

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

func DeleteTodoById(c *fiber.Ctx) error {
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
