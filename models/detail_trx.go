package models

import (
	"time"
)

type DetailTrx struct {
	Id            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Id_trx        uint      `gorm:"type:int;index" json:"id_trx"`
	Trx           Trx       `gorm:"foreignKey:Id_trx" json:"trx"`
	Id_log_produk uint      `gorm:"type:int;index" json:"Id_log_produk"`
	Log_produk    LogProduk `gorm:"foreignKey:Id_log_produk" json:"log_produk"`
	Id_toko       uint      `gorm:"type:int;index" json:"id_toko"`
	Toko          Toko      `gorm:"foreignKey:Id_toko" json:"toko"`
	Kuantitas     int       `gorm:"type:int" json:"kuantitas"`
	Harga_total   int       `gorm:"type:int" json:"harga_total"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
}

type CreateRequestTrx struct {
	IdTrx       uint `json:"id_trx"`
	IdLogProduk uint `json:"id_log_produk"`
	IdToko      uint `json:"id_toko"`
	Kuantitas   int  `json:"kuantitas"`
	HargaTotal  int  `json:"harga_total"`
}
