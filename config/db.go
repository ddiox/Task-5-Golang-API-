package config

import (
	"log"

	"github.com/ddiox/evermos_api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:123456@tcp(localhost:3306)/evermos?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Database Connection Failed to Open")
	}

	err = database.AutoMigrate(&models.User{}, &models.Alamat{}, &models.Toko{}, &models.Category{},
		&models.Produk{}, &models.FotoProduk{}, &models.LogProduk{}, &models.Trx{},
		&models.DetailTrx{})
	if err != nil {
		log.Println("Failed to perform database migration!")
	}

	log.Println("Database Connection Established")

	DB = database
}
