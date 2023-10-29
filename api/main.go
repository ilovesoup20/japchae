package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilovesoup20/japchae/database"
	"github.com/ilovesoup20/japchae/routes"
)

func main() {

	app := fiber.New()

	dbClient, _ := database.InitDB()
	defer dbClient.Close()

	routes.SetupRoutes(app, dbClient)

	app.Listen(":3000")
}
