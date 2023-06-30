package models

import (
	"time"
)

type User struct {
	Id            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama          string    `gorm:"type:varchar(255)" json:"nama"`
	Kata_sandi    string    `gorm:"column:password" json:"kata_sandi"`
	No_telp       string    `gorm:"type:varchar(255);uniqueIndex" json:"no_telp"`
	Tanggal_lahir string    `json:"tanggal_lahir"`
	Jenis_kelamin string    `gorm:"type:varchar(255)" json:"jenis_kelamin"`
	Tentang       string    `gorm:"type:text" json:"tentang"`
	Pekerjaan     string    `gorm:"type:varchar(255)" json:"pekerjaan"`
	Email         string    `gorm:"type:varchar(255);uniqueIndex" json:"email"`
	Id_provinsi   string    `gorm:"type:varchar(255)" json:"id_provinsi"`
	Id_kota       string    `gorm:"type:varchar(255)" json:"id_kota"`
	IsAdmin       bool      `json:"isAdmin"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
}

type CreateRequestUser struct {
	Nama          string `json:"nama" validate:"required"`
	Kata_sandi    string `json:"kata_sandi" validate:"required"`
	No_telp       string `json:"no_telp"`
	Tanggal_lahir string `json:"tanggal_lahir"`
	Jenis_kelamin string `json:"jenis_kelamin"`
	Tentang       string `json:"tentang"`
	Pekerjaan     string `json:"pekerjaan"`
	Email         string `json:"email" validate:"required,email"`
	Id_provinsi   string `json:"id_provinsi"`
	Id_kota       string `json:"id_kota"`
	IsAdmin       bool   `json:"isAdmin"`
}

type LoginRequest struct {
	Email      string `json:"email" validate:"required,email"`
	Kata_sandi string `json:"kata_sandi" validate:"required"`
}
