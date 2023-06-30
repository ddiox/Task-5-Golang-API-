package handlers

import (
	"log"
	"time"

	"github.com/ddiox/evermos_api/config"
	"github.com/ddiox/evermos_api/models"
	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
	productRequest := new(models.CreateRequestProduk)
	if err := c.BodyParser(productRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to parse request body",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	newProduct := models.Produk{
		Nama_produk:    productRequest.Nama_produk,
		Slug:           productRequest.Slug,
		Harga_reseller: productRequest.Harga_reseller,
		Harga_konsumen: productRequest.Harga_konsumen,
		Stok:           productRequest.Stok,
		Deskripsi:      productRequest.Deskripsi,
		Id_toko:        productRequest.Id_toko,
		Id_category:    productRequest.Id_category,
		Created_at:     time.Now(),
		Updated_at:     time.Now(),
	}

	err := config.DB.Create(&newProduct).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to store product",
			"errors":  err.Error(),
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Product successfully stored",
		"errors":  nil,
		"data":    newProduct,
	})
}

func GetAllProducts(c *fiber.Ctx) error {
	var products []models.Produk
	result := config.DB.Find(&products)
	if result.Error != nil {
		log.Println(result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to get products",
			"errors":  result.Error,
			"data":    nil,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Successfully retrieved products",
		"errors":  nil,
		"data":    products,
	})
}

func GetProductById(c *fiber.Ctx) error {
	var product models.Produk
	productID := c.Params("id")
	result := config.DB.First(&product, "id = ?", productID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  false,
			"message": "Product does not exist",
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Successfully retrieved product",
		"errors":  nil,
		"data":    product,
	})
}

func UpdateProductById(c *fiber.Ctx) error {
	productRequest := new(models.CreateRequestProduk)
	if err := c.BodyParser(productRequest); err != nil {
		return err
	}

	productID := c.Params("id")

	var product models.Produk
	err := config.DB.First(&product, "id = ?", productID).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product does not exist",
		})
	}

	if productRequest.Nama_produk != "" {
		product.Nama_produk = productRequest.Nama_produk
	}

	errUpdate := config.DB.Save(&product).Error
	if errUpdate != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
			"data":    product,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Product update success",
		"data":    product,
	})
}

func DeleteProductById(c *fiber.Ctx) error {
	productID := c.Params("id")

	var product models.Produk
	err := config.DB.First(&product, "id = ?", productID).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Product does not exist",
		})
	}

	errDelete := config.DB.Delete(&product).Error
	if errDelete != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Product has been deleted",
	})
}
