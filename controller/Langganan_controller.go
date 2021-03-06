package controller

import (
	"strconv"
	"time"

	"net/http"

	"github.com/Tugas_Besar/model"
)

func InsertLangganan(w http.ResponseWriter, r *http.Request) {
	db := connect()

	//defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	// Get data from user
	idUser := getid(r)
	// Cek sedang langganan
	var user model.User
	if db.Where("id = ?", idUser).Preload("Langganan").First(&user); user.Langganan.ID != 0 {
		sendResponse(w, 200, "Masih dalam langganan", user.Langganan)
		return
	}

	nomorKredit := r.Form.Get("nomorKredit")
	masaBerlaku := r.Form.Get("masaBerlaku")
	kodeCVC, _ := strconv.Atoi(r.Form.Get("kodeCVC"))
	memberType, _ := strconv.Atoi(r.Form.Get("memberType"))

	// Set inputted data to object
	langganan := model.Langganan{
		NomorKredit:  nomorKredit,
		MasaBerlaku:  masaBerlaku,
		KodeCVC:      kodeCVC,
		TanggalHabis: time.Now().AddDate(0, 1, 0),
		UserMember:   memberType,
		IdUser:       idUser,
	}

	// Insert object to database
	result := db.Create(&langganan)

	// Set response
	if result.Error == nil {
		// Output to console
		sendResponse(w, 200, "Success Upgrade Membreship", nil)
	} else {
		// Output to console
		sendResponse(w, 400, "Upgrade Failed", nil)
	}
}

func StopMembership(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	var langganan model.Langganan
	idUser := getid(r)

	if err := db.Where("id_user = ?", idUser).Delete(&langganan).Error; err != nil {
		sendResponse(w, 400, "Failed Stopping Membership", nil)
	} else {
		sendResponse(w, 200, "Success Stopping Membership", nil)
	}
}
