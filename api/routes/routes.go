package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilovesoup20/japchae/controllers"
	"github.com/ilovesoup20/japchae/ent"
	"github.com/ilovesoup20/japchae/repository"
	"gorm.io/gorm"
)

// SetupRoutes blah
func SetupRoutes(app *fiber.App, entClient *ent.Client, gormClient *gorm.DB) {

	app.Get("/hello", controllers.Hello)

	userRepo := repository.UserRepositoryImpl{DB: gormClient}
	authController := controllers.NewAuthController(&userRepo)
	app.Post("/login", authController.Login)
	app.Post("/register", authController.RegisterUser)

	userController := controllers.NewUserController(&userRepo)
	app.Get("/user", userController.List)

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
