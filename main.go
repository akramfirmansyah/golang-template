package main

import (
	"os"

	"github.com/akramfirmansyah/golang-template/config"
	"github.com/akramfirmansyah/golang-template/database"
	"github.com/akramfirmansyah/golang-template/database/migration"
	"github.com/akramfirmansyah/golang-template/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jsoniter "github.com/json-iterator/go"
)

func init() {
	config.LoadEnveronmentVariabel()
	database.ConnectDatabase()
	migration.MigrateModel()
}

func main() {
	app := fiber.New(fiber.Config{
		AppName:     "Golang API Template",
		JSONEncoder: jsoniter.Marshal,
		JSONDecoder: jsoniter.Unmarshal,
	})

	app.Use(logger.New())
	app.Use(cors.New())

	app.Static("/public", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/api")
	})

	api := app.Group("/api", func(c *fiber.Ctx) error {
		return c.SendString("Golang API Template by Akram Firmansyah")
	})

	routes.RegisterUserRoutes(api.Group("/users"))

	_ = app.Listen(":" + os.Getenv("PORT"))
}
