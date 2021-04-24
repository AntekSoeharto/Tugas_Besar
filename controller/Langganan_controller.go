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
	memberType := r.Form.Get("memberType")

	// Set inputted data to object
	langganan := model.Langganan{
		NomorKredit:  nomorKredit,
		MasaBerlaku:  masaBerlaku,
		KodeCVC:      kodeCVC,
		TanggalHabis: time.Now().AddDate(0, 1, 0),
		Usermember:   memberType,
	}

	// Insert object to database
	result := db.Create(&langganan)

	db.Model(model.User{}).Where("id = ?", idUser).Updates(model.User{IdLangganan: langganan.Id})

	// Set response
	var response model.LanggananResponse
	if result.Error == nil {
		// Output to console
		response.Status = 200
		response.Message = "Success Insert Film to Database"
	} else {
		// Output to console
		response.Status = 400
		response.Message = "Insert Failed"
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

	id := r.Form.Get("email")
	if id != "" {
		db.Where("id = ?", id).First(&langganan)
	} else {
		db.Find(&langganan)
	}

	// Set response
	var response model.LanggananResponse
	if langganan != nil {
		// Output to console
		response.Status = 200
		response.Message = "Success Get User Data"
		response.Data = langganan
	} else {
		// Output to console
		response.Status = 204
		response.Message = "Not Found, No Content"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
