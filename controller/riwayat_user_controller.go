package controller

import (
	"net/http"
	"time"

	"github.com/Tugas_Besar/model"
	"github.com/gorilla/mux"
)

func GetRiwayatUser(w http.ResponseWriter, r *http.Request) {
	db := connect()

	var riwayatUsers []model.RiwayatUser

	// Get RiwayatUser Query
	if err := db.Where("user_id = ?", getid(r)).Preload("Film").Find(&riwayatUsers).Error; err != nil {
		sendResponse(w, 400, "Failed to Query", nil)
	} else {
		sendResponse(w, 200, "Query Success", riwayatUsers)
	}
}

func NontonFilm(w http.ResponseWriter, r *http.Request) {
	db := connect()

	vars := mux.Vars(r)
	filmId := vars["film_id"]

	var film model.Film
	if err := db.Where("id = ?", filmId).First(&film).Error; err != nil {
		sendResponse(w, 400, "Not Found", nil)
		return
	}

	userId := getid(r)
	var user model.User
	db.Model(model.User{}).Where("id=?", userId).Preload("Langganan").First(&user)
	if film.FilmType > user.Langganan.UserMember {
		sendResponse(w, 401, "Butuh Langganan Lebih Tinggi", nil)
		return
	}

	riwayat := model.RiwayatUser{
		Tanggal: time.Now().UTC(),
		UserId:  userId,
		FilmId:  film.ID,
		Film:    film,
	}

	if err := db.Create(&riwayat).Error; err != nil {
		sendResponse(w, 400, "Riwayat tidak terbuat", nil)
	} else {
		sendResponse(w, 200, "Sedang Menonton", film)
	}
}
