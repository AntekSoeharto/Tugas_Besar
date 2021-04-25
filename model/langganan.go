package model

import (
	"time"

	"github.com/gorm"
)

type Langganan struct {
	gorm.Model
	Id           int `gorm:" primaryKey;autoincrement" form:"id" json:"id"`
	Deleted      gorm.DeletedAt
	NomorKredit  string    `form:"nomorKredit" json:"nomorKredit"`
	MasaBerlaku  string    `form:"masaBerlaku" json:"masaBerlaku"`
	KodeCVC      int       `form:"kodeCVC" json:"kodeCVC"`
	TanggalHabis time.Time `form:"tanggalHabis" json:"tanggalHabis"`
	Usermember   int       `form:"Usermember" json:"Usermember"`
	IdUser       int       `form:"iduser" json:"iduser"`
}
