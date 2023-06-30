package handlers

import (
	"time"

	"github.com/ddiox/evermos_api/config"
	"github.com/ddiox/evermos_api/models"
	"github.com/gofiber/fiber/v2"
)

func CreateToko(c *fiber.Ctx) error {
	tokoRequest := new(models.Toko)
	if err := c.BodyParser(tokoRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to parse request body",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	newToko := models.Toko{
		Id_user:    tokoRequest.Id_user,
		Nama_toko:  tokoRequest.Nama_toko,
		Url_foto:   tokoRequest.Url_foto,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	err := config.DB.Create(&newToko).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to store toko",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Toko successfully stored",
		"errors":  nil,
		"data":    newToko,
	})
}

func GetAllToko(c *fiber.Ctx) error {
	var tokos []models.Toko
	err := config.DB.Find(&tokos).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to get categories",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to get data",
		"errors":  nil,
		"data":    tokos,
	})
}

func GetTokoById(c *fiber.Ctx) error {
	var toko models.Toko
	tokoID := c.Params("id")
	result := config.DB.First(&toko, "id = ?", tokoID)
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
		"data":    toko,
	})
}

func UpdateTokoById(c *fiber.Ctx) error {
	tokoRequest := new(models.Toko)
	if err := c.BodyParser(tokoRequest); err != nil {
		return err
	}

	tokoID := c.Params("id")

	var toko models.Toko
	err := config.DB.First(&toko, "id = ?", tokoID).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Toko does not exist",
		})
	}

	if tokoRequest.Nama_toko != "" {
		toko.Nama_toko = tokoRequest.Nama_toko
	}
	if tokoRequest.Url_foto != "" {
		toko.Url_foto = tokoRequest.Url_foto
	}

	errUpdate := config.DB.Save(&toko).Error
	if errUpdate != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"data":    toko,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Category update success",
		"data":    toko,
	})
}

func DeleteTokoById(c *fiber.Ctx) error {
	tokoID := c.Params("id")

	var toko models.Toko
	err := config.DB.First(&toko, "id = ?", tokoID).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Toko does not exist",
		})
	}

	errDelete := config.DB.Delete(&toko).Error
	if errDelete != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Toko was deleted",
	})
}
