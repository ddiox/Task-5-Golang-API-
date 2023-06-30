package handlers

import (
	"log"

	"github.com/ddiox/evermos_api/config"
	"github.com/ddiox/evermos_api/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllUser(c *fiber.Ctx) error {
	var users []models.User
	result := config.DB.Find(&users)
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
		"data":    users,
	})
}

func GetUserById(c *fiber.Ctx) error {
	var user models.User
	userID := c.Params("id")
	result := config.DB.First(&user, "id = ?", userID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "User does not exist",
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Succeed to get data",
		"errors":  nil,
		"data":    user,
	})
}

func UpdateUserById(c *fiber.Ctx) error {
	userRequest := new(models.CreateRequestUser)
	if err := c.BodyParser(userRequest); err != nil {
		return err
	}

	userID := c.Params("id")

	var user models.User
	err := config.DB.First(&user, "id = ?", userID).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User does not exist",
		})
	}

	if userRequest.Nama != "" {
		user.Nama = userRequest.Nama
	}
	if userRequest.Kata_sandi != "" {
		user.Kata_sandi = userRequest.Kata_sandi
	}
	if userRequest.No_telp != "" {
		user.No_telp = userRequest.No_telp
	}
	if userRequest.Tanggal_lahir != "" {
		user.Tanggal_lahir = userRequest.Tanggal_lahir
	}
	if userRequest.Jenis_kelamin != "" {
		user.Jenis_kelamin = userRequest.Jenis_kelamin
	}
	if userRequest.Tentang != "" {
		user.Tentang = userRequest.Tentang
	}
	if userRequest.Pekerjaan != "" {
		user.Pekerjaan = userRequest.Pekerjaan
	}
	if userRequest.Email != "" {
		user.Email = userRequest.Email
	}
	if userRequest.Id_provinsi != "" {
		user.Id_provinsi = userRequest.Id_provinsi
	}
	if userRequest.Id_kota != "" {
		user.Id_kota = userRequest.Id_kota
	}

	errUpdate := config.DB.Save(&user).Error
	if errUpdate != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"data":    user,
		})
	}

	return c.JSON(fiber.Map{
		"message": "User update success",
		"data":    user,
	})
}
