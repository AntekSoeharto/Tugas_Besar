package model

import "github.com/gorm"

type Film struct {
	gorm.Model
	ID           int    `gorm:" primaryKey;autoincrement" form:"id" json:"id"`
	Judul        string `form:"judul" json:"judul"`
	TahunRilis   int    `form:"tahun_rilis" json:"tahun_rilis"`
	Genre        string `form:"genre" json:"genre"`
	Sutradara    string `form:"sutradara" json:"sutradara"`
	Sinopsis     string `form:"sinopsis" json:"sinopsis"`
	FilmType     int    `form:"film_type" json:"film_type"`
	DaftarPemain string `form:"daftar_pemain" json:"daftar_pemain"`
}
