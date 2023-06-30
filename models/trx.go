package models

import (
	"time"
)

type Trx struct {
	Id                uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Id_user           uint      `gorm:"type:int;index" json:"id_user"`
	User              User      `gorm:"foreignKey:Id_user" json:"user"`
	Alamat_pengiriman string    `gorm:"type:varchar(255)" json:"alamat_pengiriman"`
	Harga_total       int       `gorm:"type:int" json:"harga_total"`
	Kode_invoice      string    `gorm:"type:varchar(255)" json:"kode_invoice"`
	Method_bayar      string    `gorm:"type:varchar(255)" json:"method_bayar"`
	Created_at        time.Time `json:"created_at"`
	Updated_at        time.Time `json:"updated_at"`
}
