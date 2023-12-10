package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ilovesoup20/japchae/repository"
)

type UserController struct {
	User *repository.UserRepositoryImpl
}

func NewUserController(userRepository *repository.UserRepositoryImpl) *UserController {
	return &UserController{User: userRepository}
}

func (uc *UserController) List(c *fiber.Ctx) error {
	users, err := uc.User.FindAll()
	if err != nil {
		fmt.Println("error")
	}
	return c.JSON(users)
}
