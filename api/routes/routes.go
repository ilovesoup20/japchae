package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilovesoup20/japchae/controllers"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/hello", controllers.Hello)
}
