package models

import (
	"time"
)

type Alamat struct {
	Id            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Id_user       uint      `gorm:"type:int;index" json:"id_user"`
	User          User      `gorm:"foreignKey:Id_user" json:"user"`
	Judul_alamat  string    `gorm:"type:varchar(255)" json:"judul_alamat"`
	Nama_penerima string    `gorm:"type:varchar(255)" json:"nama_penerima"`
	No_telp       string    `gorm:"type:varchar(255)" json:"no_telp"`
	Detail_alamat string    `gorm:"type:varchar(255)" json:"detail_alamat"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
}

type CreateAlamatRequest struct {
	Id_user       uint   `json:"id_user"`
	Judul_alamat  string `json:"judul_alamat"`
	Nama_penerima string `json:"nama_penerima"`
	No_telp       string `json:"no_telp"`
	Detail_alamat string `json:"detail_alamat"`
}
