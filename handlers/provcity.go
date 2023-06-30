package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetListProvince(c *fiber.Ctx) error {
	resp, err := http.Get("https://emsifa.github.io/api-wilayah-indonesia/api/provinces.json")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch provinces",
		})
	}
	defer resp.Body.Close()

	var provinces []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&provinces); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to parse provinces",
		})
	}

	return c.JSON(fiber.Map{
		"message":   "Success",
		"provinces": provinces,
	})
}

func GetDetailProvince(c *fiber.Ctx) error {
	provinceID := c.Params("id")
	url := fmt.Sprintf("https://emsifa.github.io/api-wilayah-indonesia/api/province/%s.json", provinceID)

	resp, err := http.Get(url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch province",
		})
	}
	defer resp.Body.Close()

	var province struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&province); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to parse province",
		})
	}

	return c.JSON(fiber.Map{
		"message":  "Success",
		"province": province,
	})
}

func GetListCity(c *fiber.Ctx) error {
	provinceID := c.Params("id")
	url := fmt.Sprintf("https://emsifa.github.io/api-wilayah-indonesia/api/regencies/%s.json", provinceID)

	resp, err := http.Get(url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch cities",
		})
	}
	defer resp.Body.Close()

	var cities []struct {
		ID         string `json:"id"`
		ProvinceID string `json:"province_id"`
		Name       string `json:"name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&cities); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to parse cities",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Success",
		"cities":  cities,
	})
}

func GetDetailCity(c *fiber.Ctx) error {
	cityID := c.Params("city_id")
	url := fmt.Sprintf("https://emsifa.github.io/api-wilayah-indonesia/api/regency/%s.json", cityID)

	resp, err := http.Get(url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch city",
		})
	}
	defer resp.Body.Close()

	var city struct {
		ID         string `json:"id"`
		ProvinceID string `json:"province_id"`
		Name       string `json:"name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&city); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to parse city",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Success",
		"city":    city,
	})
}
