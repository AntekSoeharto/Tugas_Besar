package model

import (
	"time"

	"github.com/gorm"
)

type Langganan struct {
	gorm.Model
	Id           int       `gorm:" primaryKey;autoincrement" form:"id" json:"id"`
	NomorKredit  string    `form:"nomorKredit" json:"nomorKredit"`
	MasaBerlaku  string    `form:"masaBerlaku" json:"masaBerlaku"`
	KodeCVC      int       `form:"kodeCVC" json:"kodeCVC"`
	TanggalHabis time.Time `form:"tanggalHabis" json:"tanggalHabis"`
	Usermember   string    `form:"Usermember" json:"Usermember"`
	IdUser       int       `form:"iduser" json:"iduser"`
}

type LanggananResponse struct {
	Status  int         `form:"status" json:"status"`
	Message string      `form:"message" json:"mesage"`
	Data    []Langganan `form:"data" json:"data"`
}
