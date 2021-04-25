package controller

import (
	"net/http"
	"time"

	"github.com/Tugas_Besar/model"
	"github.com/gorilla/mux"
)

func NontonFilm(w http.ResponseWriter, r *http.Request) {
	db := connect()

	vars := mux.Vars(r)
	filmId := vars["film_id"]

	var film model.Film
	if err := db.Where("id = ?", filmId).First(&film).Error; err != nil {
		sendResponse(w, 204, "Not Found, No Content", nil)
		return
	}

	userId := getid(r)
	var user model.User
	db.Model(model.User{}).Where("id=?", userId).First(&user)
	if film.Filmtype > user.Langganan.Usermember {
		sendResponse(w, 401, "Butuh Langganan Lebih Tinggi", nil)
		return
	}

	riwayat := model.RiwayatUser{
		Tanggal: time.Now().UTC(),
		UserId:  userId,
		FilmId:  film.Id,
		Film:    film,
	}

	if err := db.Create(&riwayat).Error; err != nil {
		sendResponse(w, 400, "Riwayat tidak terbuat", nil)
	} else {
		sendResponse(w, 200, "Sedang Menonton", film)
	}
}
