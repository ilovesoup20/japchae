package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilovesoup20/japchae/controllers"
)

// SetupRoutes blah
func SetupRoutes(app *fiber.App) {

	app.Get("/hello", controllers.Hello)

	// TODO
	app.Get("/todo", controllers.ListTodo)
	app.Get("/todo/:id", controllers.GetTodoByID)
	app.Post("/todo", controllers.CreateTodo)
	app.Put("/todo/:id", controllers.UpdateTodo)
	app.Patch("/todo/:id", controllers.PatchTodo)
	app.Delete("/todo/:id", controllers.DeleteTodoByID)

}
