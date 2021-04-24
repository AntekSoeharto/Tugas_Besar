package model

import "github.com/gorm"

type DaftarPemain struct {
	gorm.Model
	Id     int    `gorm:" primaryKey;autoincrement" form:"id" json:"id"`
	IdFilm int    `form:"IdFilm" json:"IdFilm"`
	Nama   string `form:"Nama" json:"Nama"`
}

type DaftarPemainResponse struct {
	Status  int            `form:"status" json:"status"`
	Message string         `form:"message" json:"message"`
	Data    []DaftarPemain `form:"data" json:"data"`
}
