package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks = []Task{
	{ID: 1, Title: "Buy groceries", Done: false},
	{ID: 2, Title: "Exercise", Done: false},
}

func TaskList(c *fiber.Ctx) error {
	return c.JSON(tasks)
}
