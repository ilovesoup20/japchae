package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilovesoup20/japchae/database"
	"github.com/ilovesoup20/japchae/routes"
)

func main() {

	app := fiber.New()

	entClient, _ := database.InitEntDB()
	defer entClient.Close()

	gormClient, _ := database.InitGormDB()
	// defer gormClient.Close()

	routes.SetupRoutes(app, entClient, gormClient)

	app.Listen(":3000")
}
