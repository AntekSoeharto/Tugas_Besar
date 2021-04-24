package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Tugas_Besar/model"
)

func GetRiwayatUser(w http.ResponseWriter, r *http.Request) {
	db := connect()

	var riwayatUsers []model.RiwayatUser

	// Get RiwayatUser Query
	if err := db.Find(&riwayatUsers).Where("iduser = ?", getid(r)).Error; err != nil {
		sendResponse(w, 400, "Failed to Query", nil)
	} else if len(riwayatUsers) == 0 {
		sendResponse(w, 204, "Not Found, No Content", nil)
	} else {
		sendResponse(w, 200, "Query Success", riwayatUsers)
	}
}

func InsertRiwayatUser(w http.ResponseWriter, r *http.Request) {
	db := connect()

	err := r.ParseForm()
	if err != nil {
		return
	}

	filmId, _ := strconv.Atoi(r.Form.Get("filmId"))

	riwayatUser := model.RiwayatUser{
		Tanggal: time.Now().UTC(),
		UserId:  getid(r),
		FilmId:  filmId,
	}

	if err := db.Create(&riwayatUser).Error; err != nil {
		sendResponse(w, 400, "Failed to Insert", nil)
	} else {
		sendResponse(w, 200, "Insert Success", []model.RiwayatUser{riwayatUser})
	}
}

func sendResponse(w http.ResponseWriter, status int, msg string, data []model.RiwayatUser) {
	var response model.RiwayatUserResponse
	response.Status = status
	response.Message = msg
	response.Data = data

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
