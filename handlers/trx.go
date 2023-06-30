package handlers

import (
	"log"
	"time"

	"github.com/ddiox/evermos_api/config"
	"github.com/ddiox/evermos_api/models"
	"github.com/gofiber/fiber/v2"
)

func CreateTrx(c *fiber.Ctx) error {
	trxRequest := new(models.CreateRequestTrx)
	if err := c.BodyParser(trxRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to parse request body",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	newTrx := models.DetailTrx{
		Id_trx:        trxRequest.IdTrx,
		Id_log_produk: trxRequest.IdLogProduk,
		Id_toko:       trxRequest.IdToko,
		Kuantitas:     trxRequest.Kuantitas,
		Harga_total:   trxRequest.HargaTotal,
		Created_at:    time.Now(),
		Updated_at:    time.Now(),
	}

	err := config.DB.Create(&newTrx).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to store transaction",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Transaction successfully stored",
		"errors":  nil,
		"data":    newTrx,
	})
}

func GetAllTrx(c *fiber.Ctx) error {
	var transactions []models.DetailTrx
	result := config.DB.Find(&transactions)
	if result.Error != nil {
		log.Println(result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to get transactions",
			"errors":  result.Error,
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Successfully retrieved transactions",
		"errors":  nil,
		"data":    transactions,
	})
}

func GetTrxById(c *fiber.Ctx) error {
	var transaction models.DetailTrx
	transactionID := c.Params("id")
	result := config.DB.First(&transaction, "id = ?", transactionID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Transaction does not exist",
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Successfully retrieved transaction",
		"errors":  nil,
		"data":    transaction,
	})
}

func UpdateTrxById(c *fiber.Ctx) error {
	trxRequest := new(models.CreateRequestTrx)
	if err := c.BodyParser(trxRequest); err != nil {
		return err
	}

	transactionID := c.Params("id")

	var transaction models.DetailTrx
	err := config.DB.First(&transaction, "id = ?", transactionID).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Transaction does not exist",
		})
	}

	if trxRequest.IdLogProduk != 0 {
		transaction.Id_log_produk = trxRequest.IdLogProduk
	}
	if trxRequest.IdToko != 0 {
		transaction.Id_toko = trxRequest.IdToko
	}
	if trxRequest.Kuantitas != 0 {
		transaction.Kuantitas = trxRequest.Kuantitas
	}
	if trxRequest.HargaTotal != 0 {
		transaction.Harga_total = trxRequest.HargaTotal
	}

	errUpdate := config.DB.Save(&transaction).Error
	if errUpdate != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"data":    transaction,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Transaction update success",
		"data":    transaction,
	})
}

func DeleteTrxById(c *fiber.Ctx) error {
	transactionID := c.Params("id")

	var transaction models.DetailTrx
	err := config.DB.First(&transaction, "id = ?", transactionID).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Transaction does not exist",
		})
	}

	errDelete := config.DB.Delete(&transaction).Error
	if errDelete != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Transaction was deleted",
	})
}
