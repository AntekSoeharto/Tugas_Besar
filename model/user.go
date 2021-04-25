package model

import (
	"github.com/gorm"
)

type User struct {
	gorm.Model
	ID           int           `gorm:" primaryKey;autoincrement" form:"id" json:"id"`
	Nama         string        `form:"nama" json:"nama"`
	Email        string        `form:"email" json:"email"`
	Password     string        `form:"password" json:"password"`
	TglLahir     string        `form:"tgl_lahir" json:"tgl_lahir"`
	JenisKelamin string        `form:"jenis_kelamin" json:"jenis_kelamin"`
	AsalNegara   string        `form:"asal_negara" json:"asal_negara"`
	UserType     int           `form:"user_type" json:"user_type"`
	Block        int           `form:"block" json:"block"`
	RiwayatUsers []RiwayatUser `gorm:"foreignKey:UserId" form:"riwayat_users" json:"riwayat_users"`
	Langganan    Langganan     `gorm:"foreignKey:IdUser" form:"langganan" json:"langganan"`
}
