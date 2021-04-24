package model

import (
	"time"

	"github.com/gorm"
)

type RiwayatUser struct {
	gorm.Model
	Id      int       `gorm:"primaryKey;autoincrement" form:"id" json:"id"`
	Tanggal time.Time `form:"tanggal" json:"tanggal"`
	UserId  int       `form:"user_id" json:"user_id"`
	FilmId  int       `form:"film_id" json:"film_id"`
}

type RiwayatUserResponse struct {
	Status  int           `form:"status" json:"status"`
	Message string        `form:"message" json:"message"`
	Data    []RiwayatUser `form:"data" json:"data"`
}
