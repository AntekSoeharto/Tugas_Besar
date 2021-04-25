package model

import (
	"github.com/gorm"
)

type User struct {
	gorm.Model
	Id           int           `gorm:" primaryKey;autoincrement" form:"id" json:"id"`
	Nama         string        `form:"Nama" json:"Nama"`
	Email        string        `form:"Email" json:"Email"`
	Password     string        `form:"Password" json:"Password"`
	TglLahir     string        `form:"Tgllahir" json:"Tgllahir"`
	Jeniskelamin string        `form:"Jeniskelamin" json:"Jeniskelamin"`
	Asalnegara   string        `form:"Asalnegara" json:"Asalnegara"`
	Usertype     int           `form:"Usertype" json:"Usertype"`
	Block        int           `form:"block" json:"block"`
	RiwayatUsers []RiwayatUser `gorm:"foreignKey:UserId"`
	Langganan    Langganan     `gorm:"foreignKey:IdUser"`
}

type UserResponse struct {
	Status  int    `form:"status" json:"status"`
	Message string `form:"message" json:"message"`
	Data    []User `form:"data" json:"data"`
}
