package handlers

import (
	"log"
	"time"

	"github.com/ddiox/evermos_api/config"
	"github.com/ddiox/evermos_api/models"
	"github.com/gofiber/fiber/v2"
)

func CreateAlamat(c *fiber.Ctx) error {
	alamatRequest := new(models.CreateAlamatRequest)
	if err := c.BodyParser(alamatRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to parse request body",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	newAlamat := models.Alamat{
		Id_user:       alamatRequest.Id_user,
		Judul_alamat:  alamatRequest.Judul_alamat,
		Nama_penerima: alamatRequest.Nama_penerima,
		No_telp:       alamatRequest.No_telp,
		Detail_alamat: alamatRequest.Detail_alamat,
		Created_at:    time.Now(),
		Updated_at:    time.Now(),
	}

	err := config.DB.Create(&newAlamat).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to store alamat",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Alamat successfully stored",
		"errors":  nil,
		"data":    newAlamat,
	})
}

func GetAllAlamat(c *fiber.Ctx) error {
	var alamat []models.Alamat
	result := config.DB.Find(&alamat)
	if result.Error != nil {
		log.Println(result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to get alamat",
			"errors":  result.Error,
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to get data",
		"errors":  nil,
		"data":    alamat,
	})
}

func GetAlamatById(c *fiber.Ctx) error {
	var alamat models.Alamat
	alamatID := c.Params("id")
	result := config.DB.First(&alamat, "id = ?", alamatID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Alamat does not exist",
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to get data",
		"errors":  nil,
		"data":    alamat,
	})
}

func UpdateAlamatById(c *fiber.Ctx) error {
	alamatRequest := new(models.CreateAlamatRequest)
	if err := c.BodyParser(alamatRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to parse request body",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	alamatID := c.Params("id")

	var alamat models.Alamat
	err := config.DB.First(&alamat, "id = ?", alamatID).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Alamat does not exist",
		})
	}

	if alamatRequest.Id_user != 0 {
		alamat.Id_user = alamatRequest.Id_user
	}

	if alamatRequest.Judul_alamat != "" {
		alamat.Judul_alamat = alamatRequest.Judul_alamat
	}

	if alamatRequest.Nama_penerima != "" {
		alamat.Nama_penerima = alamatRequest.Nama_penerima
	}

	if alamatRequest.No_telp != "" {
		alamat.No_telp = alamatRequest.No_telp
	}

	if alamatRequest.Detail_alamat != "" {
		alamat.Detail_alamat = alamatRequest.Detail_alamat
	}

	alamat.Updated_at = time.Now()

	errUpdate := config.DB.Save(&alamat).Error
	if errUpdate != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to update alamat",
			"errors":  errUpdate.Error(),
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Alamat successfully updated",
		"errors":  nil,
		"data":    alamat,
	})
}

func DeleteAlamatById(c *fiber.Ctx) error {
	alamatID := c.Params("id")

	var alamat models.Alamat
	err := config.DB.First(&alamat, "id = ?", alamatID).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Alamat does not exist",
		})
	}

	errDelete := config.DB.Delete(&alamat).Error
	if errDelete != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to delete alamat",
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Alamat successfully deleted",
	})
}
