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
	Film    Film      `gorm:"foreignKey:FilmId" form:"film" json:"film"`
}
