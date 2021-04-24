package model

import "github.com/gorm"

type Langganan struct {
	gorm.Model
	Id          int    `gorm:" primaryKey;autoincrement" form:"id" json:"id"`
	IdUser      int    `gorm:"foreignKey" form:"idUser" json:"idUser"`
	NomorKredit string `form:"nomorKredit" json:"nomorKredit"`
	MasaBerlaku string `form:"masaBerlaku" json:"masaBerlaku"`
	KodeCVC     int    `form:"kodeCVC" json:"kodeCVC"`
}

type LanggananResponse struct {
	Status  int         `form:"status" json:"status"`
	Message string      `form:"message" json:"message"`
	Data    []Langganan `form:"data" json:"data"`
}
