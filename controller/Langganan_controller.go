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
	nomorKredit := r.Form.Get("nomorKredit")
	masaBerlaku := r.Form.Get("masaBerlaku")
	kodeCVC, _ := strconv.Atoi(r.Form.Get("kodeCVC"))
	memberType, _ := strconv.Atoi(r.Form.Get("kodeCVC"))

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

func GetLangganan(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	var langganan []model.Langganan

	id := getid(r)
	db.Where("id_user = ", id).First(&langganan)

	// Set response
	if langganan != nil {
		// Output to console
		sendResponse(w, 200, "Success Get Data", langganan)
	} else {
		// Output to console
		sendResponse(w, 204, "Not Found, No Content", nil)
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
