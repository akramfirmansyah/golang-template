package controllers

import (
	"github.com/akramfirmansyah/golang-template/database"
	"github.com/akramfirmansyah/golang-template/database/model"
	"github.com/akramfirmansyah/golang-template/utils"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	var user model.User

	err := database.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return fiber.ErrUnauthorized
	}

	if !utils.CheckPassword(password, string(user.Password)) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Invalid password!",
			"data":    nil,
		})
	}

	token, err := utils.GetToken(user.Username)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Success",
		"data":    token,
	})
}
