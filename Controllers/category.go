package Controllers

import (
	"github.com/Shawket4/FlavourFusion/Models"
	"github.com/gofiber/fiber/v2"
)

func RegisterCategory(c *fiber.Ctx) error {
	var category Models.Category

	if err := c.BodyParser(&category); err != nil {
		return c.JSON(fiber.Map{"message": "Couldn't Parse Data"})
	}

	if err := Models.DB.Create(&category).Error; err != nil {
		return c.JSON(fiber.Map{"message": "Couldn't Register Item"})
	}

	return c.JSON(fiber.Map{"message": "Category Registered Successfully"})
}

func FetchCategories(c *fiber.Ctx) error {
	var categories []Models.Category
	if err := Models.DB.Model(&Models.Category{}).Preload("Items").Find(&categories).Error; err != nil {
		return c.JSON(fiber.Map{"message": "Couldn't Fetch Categories"})
	}
	return c.JSON(categories)
}

func DeleteCategory(c *fiber.Ctx) error {
	var input struct {
		ID uint `json:"id"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := Models.DB.Model(&Models.Category{}).Delete("id = ?", input.ID).Error; err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Category Deleted Successfully",
	})
}
