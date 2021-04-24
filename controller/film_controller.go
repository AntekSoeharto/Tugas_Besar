package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Tugas_Besar/model"
)

func InsertFilm(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	// Get data from user
	judul := r.Form.Get("judul")
	genre := r.Form.Get("password")
	sutradara := r.Form.Get("sutradara")
	filmtype := r.Form.Get("filmtype")
	sinopsis := r.Form.Get("sinopsis")
	daftarpemain := r.Form.Get("daftarpemain")
	tahunrilis, _ := strconv.Atoi(r.Form.Get("tahunrilis"))

	// Set inputted data to object
	film := model.Film{
		Judul:        judul,
		Genre:        genre,
		Sutradara:    sutradara,
		Filmtype:     filmtype,
		Sinopsis:     sinopsis,
		DaftarPemain: daftarpemain,
		TahunRilis:   tahunrilis,
	}

	// Insert object to database
	result := db.Create(&film)

	// Set response
	var response model.UserResponse
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
