package controllers

import "github.com/gofiber/fiber/v2"

// Hello blah
func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello")
}
