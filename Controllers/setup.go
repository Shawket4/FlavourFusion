package Controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const Port = 3021

func Setup() {
	app := fiber.New()
	app.Post("/api/RegisterItem", RegisterItem)
	app.Post("/api/RegisterCategory", RegisterCategory)
	app.Get("/api/FetchCategories", FetchCategories)
	app.Post("/api/DeleteCategory", DeleteCategory)
	app.Post("/api/DeleteItem", DeleteItem)
	app.Post("/api/UpdateItem", UpdateItem)
	app.Static("/ItemPhotos", "./ItemPhotos")
	if err := app.Listen(fmt.Sprintf(":%v", Port)); err != nil {
		panic(fmt.Sprintf("Couldn't Listen On Port: %v", Port))
	}
}
