package Controllers

import (
	"encoding/json"
	"fmt"

	"github.com/Shawket4/FlavourFusion/Models"
	"github.com/gofiber/fiber/v2"
)

func RegisterItem(c *fiber.Ctx) error {
	var item Models.Item
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(fiber.Map{"message": "Couldn't Parse Data"})
	}
	fmt.Println(form.Value["item"][0])

	if json.Unmarshal([]byte(form.Value["item"][0]), &item); err != nil {
		return c.JSON(fiber.Map{"message": "Couldn't Parse Data"})
	}
	// if err := c.BodyParser(&item); err != nil {
	// 	return c.JSON(fiber.Map{"message": "Couldn't Parse Data"})
	// }
	files := form.File["photos"]
	for _, file := range files {
		if err := c.SaveFile(file, fmt.Sprintf("./ItemPhotos/%s%s", item.Name, file.Filename)); err != nil {
			return err
		}
		item.ImageURL = "ItemPhotos/" + item.Name + file.Filename
	}
	if err := Models.DB.Create(&item).Error; err != nil {
		return c.JSON(fiber.Map{"message": "Couldn't Register Item"})
	}

	return c.JSON(fiber.Map{"message": "Item Registered Successfully"})
}

func DeleteItem(c *fiber.Ctx) error {
	var input struct {
		ID uint `json:"id"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := Models.DB.Model(&Models.Item{}).Delete("id = ?", input.ID).Error; err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Item Deleted Successfully",
	})
}

func UpdateItem(c *fiber.Ctx) error {
	var input Models.Item
	if err := c.BodyParser(&input); err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := Models.DB.Save(&input).Error; err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Item Updated Successfully",
	})
}
