package handlers

import (
	"log"

	"github.com/ddiox/evermos_api/config"
	"github.com/ddiox/evermos_api/models"
	"github.com/gofiber/fiber/v2"
)

func CreateCategory(c *fiber.Ctx) error {
	categoryRequest := new(models.CreateRequestCategory)
	if err := c.BodyParser(categoryRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to parse request body",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	newCategory := models.Category{
		Nama_category: categoryRequest.Nama_category,
	}

	err := config.DB.Create(&newCategory).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to store category",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Category successfully stored",
		"errors":  nil,
		"data":    newCategory,
	})
}

func GetAllCategory(c *fiber.Ctx) error {
	var categories []models.Category
	result := config.DB.Find(&categories)
	if result.Error != nil {
		log.Println(result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to get categories",
			"errors":  result.Error,
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to get data",
		"errors":  nil,
		"data":    categories,
	})
}

func GetCategoryById(c *fiber.Ctx) error {
	var category models.Category
	categoryID := c.Params("id")
	result := config.DB.First(&category, "id = ?", categoryID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Category does not exist",
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to get data",
		"errors":  nil,
		"data":    category,
	})
}

func UpdateCategoryById(c *fiber.Ctx) error {
	categoryRequest := new(models.CreateRequestCategory)
	if err := c.BodyParser(categoryRequest); err != nil {
		return err
	}

	categoryID := c.Params("id")

	var categories models.Category
	err := config.DB.First(&categories, "id = ?", categoryID).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Category does not exist",
		})
	}

	if categoryRequest.Nama_category != "" {
		categories.Nama_category = categoryRequest.Nama_category
	}

	errUpdate := config.DB.Save(&categories).Error
	if errUpdate != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"data":    categories,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Category update success",
		"data":    categories,
	})
}

func DeleteCategoryById(c *fiber.Ctx) error {
	categoryID := c.Params("id")

	var categories models.Category
	err := config.DB.First(&categories, "id = ?", categoryID).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Category does not exist",
		})
	}

	errDelete := config.DB.Delete(&categories).Error
	if errDelete != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Category was deleted",
	})
}
