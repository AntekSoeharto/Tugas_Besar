package model

import "github.com/gorm"

type Film struct {
	gorm.Model
	Id           int    `gorm:" primaryKey;autoincrement" form:"id" json:"id"`
	Judul        string `form:"judul" json:"judul"`
	TahunRilis   int    `form:"tahunRIlis" json:"tahunRIlis"`
	Genre        string `form:"genre" json:"genre"`
	Sutradara    string `form:"sutradara" json:"sutradara"`
	Sinopsis     string `form:"sinopsis" json:"sinopsis"`
	Filmtype     int    `form:"filmtype" json:"filmtype"`
	DaftarPemain string `form:"daftarpemain" json:"daftarpemain"`
}
