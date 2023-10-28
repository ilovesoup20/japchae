package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ilovesoup20/japchae/routes"
)

func main() {
	fmt.Println("hello world")
	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
