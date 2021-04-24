package controller

import (
	"encoding/json"
	"strconv"

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

	// Set inputted data to object
	langganan := model.Langganan{
		IdUser:      idUser,
		NomorKredit: nomorKredit,
		MasaBerlaku: masaBerlaku,
		KodeCVC:     kodeCVC,
	}

	// Insert object to database
	result := db.Create(&langganan)

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
