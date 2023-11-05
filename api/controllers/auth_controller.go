package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ilovesoup20/japchae/auth"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string // Unique username
	Password string // Hashed password
}

var users = map[string]User{
	"john": {
		Username: "john",
		Password: hashPassword("password123"),
	},
	"jane": {
		Username: "jane",
		Password: hashPassword("securepass"),
	},
}

func Login(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, found := users[username]

	if !found {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid username or password",
		})
	}

	err := verifyPassword(user.Password, password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid username or password",
		})
	}

	token, err := auth.GenerateToken(username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate JWT token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func hashPassword(password string) string {
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hashedPw)
}
func verifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
