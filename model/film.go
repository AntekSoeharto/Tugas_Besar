package model

import "github.com/gorm"

type Film struct {
	gorm.Model
	Id           int            `gorm:" primaryKey;autoincrement" form:"id" json:"id"`
	Judul        string         `form:"judul" json:"judul"`
	TahunRilis   int            `form:"tahunRIlis" json:"tahunRIlis"`
	Genre        string         `form:"genre" json:"genre"`
	Sutradara    string         `form:"sutradara" json:"sutradara"`
	Sinopsis     string         `form:"sutradara" json:"sutradara"`
	Filmtype     string         `form:"filmtype" json:"filmtype"`
	DaftarPemain []DaftarPemain `gorm:"foreignKey:IdFilm;references:Id"`
}

type FIlmResponse struct {
	Status  int    `form:"status" json:"status"`
	Message string `form:"message" json:"message"`
	Data    []Film `form:"data" json:"data"`
}
