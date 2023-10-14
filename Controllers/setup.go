package Controllers

import (
	"github.com/gofiber/fiber/v2"
	"fmt"
)

const Port = 3021

func Setup() {
	app := fiber.New()
	app.Post("/api/RegisterItem", RegisterItem)
	app.Post("/api/RegisterCategory", RegisterCategory)
	app.Get("/api/FetchCategories", FetchCategories)
	app.Static("/ItemPhotos", "./ItemPhotos")
	if err := app.Listen(fmt.Sprintf(":%v", Port)); err != nil {
		panic(fmt.Sprintf("Couldn't Listen On Port: %v", Port))
	}
}
