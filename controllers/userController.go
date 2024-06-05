package controllers

import "github.com/gofiber/fiber/v2"

func CreateUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Success",
		"data":    nil,
	})
}

func GetAllUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Success",
		"data":    nil,
	})
}

func GetUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Success",
		"data":    nil,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Success",
		"data":    nil,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Success",
		"data":    nil,
	})
}
