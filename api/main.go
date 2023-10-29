package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ilovesoup20/japchae/controllers"
	"github.com/ilovesoup20/japchae/database"
	"github.com/ilovesoup20/japchae/routes"
)

func main() {
	fmt.Println("hello world")
	app := fiber.New()

	dbClient, _ := database.InitDB()
	defer dbClient.Close()

	todoController := controllers.NewTodoController(dbClient)

	app.Get("/todo2", todoController.List)
	app.Get("/todo2/:id", todoController.GetByID)

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
