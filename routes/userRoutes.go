package routes

import (
	"github.com/akramfirmansyah/golang-template/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(route fiber.Router) {
	route.Post("/", controllers.CreateUser)
	route.Get("/", controllers.GetAllUser)
	route.Get("/:id", controllers.GetUser)
	route.Put("/:id", controllers.UpdateUser)
	route.Delete("/:id", controllers.DeleteUser)
}
