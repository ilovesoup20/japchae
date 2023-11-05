package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ilovesoup20/japchae/auth"
	"golang.org/x/crypto/bcrypt"
)

// User .
type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
	Salt     string
}

var users = map[string]User{
	"john": {
		Username: "john",
	},
	"jane": {
		Username: "jane",
	},
	"charles": {
		Username: "charles",
	},
}

// Login .
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

	token, err := auth.GenerateToken()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate JWT token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

// RegisterUser .
func RegisterUser(c *fiber.Ctx) error {
	var newUser User
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	salt := generateSalt()
	newUser.Salt = salt

	hashedPassword := hashPassword(newUser.Password, salt)
	newUser.Password = hashedPassword

	result := db.Create(&newUser)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": result.Error.error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newUser)
}

func generateSalt() string {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(salt)
}

func hashPassword(password, salt string) string {
	saltedPassword := []byte(password + salt)
	hashedPassword, err := bcrypt.GenerateFromPassword(saltedPassword, bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}

	return string(hashedPassword)
}

func validatePassword(inputPassword, storedHashedPassword, salt string) bool {
	hashedInputPassword := hashPassword(inputPassword, salt)
	return hashedInputPassword == storedHashedPassword
}
func verifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
