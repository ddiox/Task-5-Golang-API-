package handlers

import (
	"log"

	"github.com/ddiox/evermos_api/config"
	"github.com/ddiox/evermos_api/helpers"
	"github.com/ddiox/evermos_api/middlewares"
	"github.com/ddiox/evermos_api/models"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	user := new(models.CreateRequestUser)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Storing data is failed",
			"error":   errValidate.Error(),
		})
	}

	newUser := models.User{
		Nama:          user.Nama,
		No_telp:       user.No_telp,
		Tanggal_lahir: user.Tanggal_lahir,
		Jenis_kelamin: user.Jenis_kelamin,
		Tentang:       user.Tentang,
		Pekerjaan:     user.Pekerjaan,
		Email:         user.Email,
		Id_provinsi:   user.Id_provinsi,
		Id_kota:       user.Id_kota,
		IsAdmin:       user.IsAdmin,
	}

	encryptedPassword, err := helpers.Encrypt(user.Kata_sandi, "your-secret-key")
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	newUser.Kata_sandi = encryptedPassword

	errCreateUser := config.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Storing data is failed",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data successfully stored",
		"data":    newUser,
	})
}

func Login(c *fiber.Ctx) error {
	req := new(models.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(req)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Login failed",
			"error":   errValidate.Error(),
		})
	}

	var user models.User
	err := config.DB.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Login failed",
			"error":   "Invalid credentials",
		})
	}

	decryptedPassword, err := helpers.Decrypt(user.Kata_sandi, "your-secret-key")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	if decryptedPassword != req.Kata_sandi {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Login failed",
			"error":   "Invalid credentials",
		})
	}

	token, err := middlewares.GenerateToken(int(user.Id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
	})
}
