package main

import (
	"github.com/gofiber/fiber/v2"
	"shopnilashik/config"
	"shopnilashik/handlers"
)

func main() {
	app := fiber.New()
	config.Connect()
	app.Post("/add", handlers.Create)
	app.Get("/posts", handlers.GetPosts)
	app.Get("/post/:id", handlers.GetPost)
	app.Put("/post/:id", handlers.UpdatePost)
	app.Delete("post/:id", handlers.DeletePost)
	app.Listen(":8000")

}
