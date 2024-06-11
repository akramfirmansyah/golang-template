package controllers

import (
	"errors"

	"github.com/akramfirmansyah/golang-template/database"
	"github.com/akramfirmansyah/golang-template/database/model"
	"github.com/akramfirmansyah/golang-template/utils"
	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateUser(c *fiber.Ctx) error {
	var mysqlErr *mysql.MySQLError

	pass, err := utils.HashPassword(c.FormValue("password"))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Couldn't hash password",
			"data":    nil,
		})
	}

	data := model.User{
		Username: c.FormValue("username"),
		Email:    c.FormValue("email"),
		Password: pass,
	}

	if err := database.DB.Create(&data).Error; err != nil {
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return c.Status(500).JSON(fiber.Map{
				"status":  fiber.StatusInternalServerError,
				"message": "User already exist!",
				"data":    nil,
			})
		}

		return c.Status(500).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Can't create user!",
			"data":    nil,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  fiber.StatusCreated,
		"message": "Success",
		"data": model.UserRespons{
			ID:       data.ID,
			Username: data.Username,
			Email:    data.Email,
		},
	})
}

func GetAllUser(c *fiber.Ctx) error {
	var users []model.UserRespons

	if err := database.DB.Model(&model.User{}).Find(&users).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Failed to find users!",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Success",
		"data":    users,
	})
}

func GetUser(c *fiber.Ctx) error {
	var user model.UserRespons

	id := c.Params("id")

	result := database.DB.Model(&model.User{}).Where("id = ?", id).Find(&user)

	if err := result.Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Failed to find users!",
			"data":    nil,
		})
	}

	if result.RowsAffected == 0 {
		return c.Status(422).JSON(fiber.Map{
			"status":  fiber.StatusUnprocessableEntity,
			"message": "User not found!. Please check the input again",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Success",
		"data":    user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user model.User

	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"status":  fiber.StatusNotFound,
				"message": "Failed to find users!",
				"data":    nil,
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to update users!",
			"data":    nil,
		})
	}

	if c.FormValue("username") != "" {
		user.Username = c.FormValue("username")
	}
	if c.FormValue("email") != "" {
		user.Email = c.FormValue("email")
	}

	if c.FormValue("password") != "" {
		pass, err := utils.HashPassword(c.FormValue("password"))
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"status":  fiber.StatusInternalServerError,
				"message": "Couldn't hash password",
				"data":    nil,
			})
		}

		user.Password = pass
	}

	if err := database.DB.Model(&user).Updates(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to update users!",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Success",
		"data": model.UserRespons{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user model.User

	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return c.Status(422).JSON(fiber.Map{
			"status":  fiber.StatusUnprocessableEntity,
			"message": "User not found!",
			"data":    nil,
		})
	}

	if err := database.DB.Where("id = ?", id).Delete(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Failed to delete user",
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Success delete user",
		"data":    nil,
	})
}
