package Controllers

import (
	"fmt"
	"github.com/Shawket4/FlavourFusion/Models"
	"github.com/gofiber/fiber/v2"
	"encoding/json"
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
		item.ImageURL = "ItemPhotos/"+item.Name+file.Filename
	}
	if err := Models.DB.Create(&item).Error; err != nil {
		return c.JSON(fiber.Map{"message": "Couldn't Register Item"})
	}

	return c.JSON(fiber.Map{"message": "Item Registered Successfully"})
}
