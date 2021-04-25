package model

import (
	"time"

	"github.com/gorm"
)

type Langganan struct {
	gorm.Model
	ID           int       `gorm:" primaryKey;autoincrement" form:"id" json:"id"`
	NomorKredit  string    `form:"nomor_kredit" json:"nomor_kredit`
	MasaBerlaku  string    `form:"masa_berlaku" json:"masa_berlaku`
	KodeCVC      int       `form:"kode_CVC" json:"kode_CVC`
	TanggalHabis time.Time `form:"tanggal_habis" json:"tanggal_habi"`
	UserMember   int       `form:"user_member" json:"user_member`
	IdUser       int       `form:"id_user" json:"id_user`
}
