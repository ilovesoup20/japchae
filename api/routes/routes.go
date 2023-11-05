package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilovesoup20/japchae/controllers"
	"github.com/ilovesoup20/japchae/ent"
)

// SetupRoutes blah
func SetupRoutes(app *fiber.App, entClient *ent.Client) {

	app.Get("/hello", controllers.Hello)

	app.Post("/auth", controllers.Login)

	// TODO
	todoController := controllers.NewTodoController(entClient.Todo)
	app.Get("/todo", todoController.ListTodos)
	app.Get("/todo/:id", todoController.GetTodoByID)
	app.Post("/todo", todoController.CreateTodo)
	app.Put("/todo/:id", todoController.UpdateTodo)

	// app.Get("/todo", controllers.ListTodo)
	// app.Get("/todo/:id", controllers.GetTodoByID)
	// app.Post("/todo", controllers.CreateTodo)
	// app.Put("/todo/:id", controllers.UpdateTodo)
	// app.Patch("/todo/:id", controllers.PatchTodo)
	// app.Delete("/todo/:id", controllers.DeleteTodoByID)
}
