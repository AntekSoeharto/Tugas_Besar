package controller

import (
	"encoding/json"
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
		Usermember:   memberType,
		IdUser:       idUser,
	}

	// Insert object to database
	result := db.Create(&langganan)

	// Set response
	var response model.LanggananResponse
	if result.Error == nil {
		// Output to console
		response.Status = 200
		response.Message = "Success Upgrade Membreship"
	} else {
		// Output to console
		response.Status = 400
		response.Message = "Upgrade Failed"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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
	var response model.LanggananResponse
	if langganan != nil {
		// Output to console
		response.Status = 200
		response.Message = "Success Get Data"
		response.Data = langganan
	} else {
		// Output to console
		response.Status = 204
		response.Message = "Not Found, No Content"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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

	var response model.LanggananResponse
	if err := db.Where("id_user = ?", idUser).Delete(&langganan).Error; err != nil {
		response.Status = 400
		response.Message = "Failed Stopping Membership"
	} else {
		response.Status = 200
		response.Message = "Success Stopping Membership"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
