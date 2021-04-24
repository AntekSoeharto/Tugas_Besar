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
	genre := r.Form.Get("genre")
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
	var response model.FilmResponse
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

func GetFilmsbyAdmin(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	var films []model.Film

	judul := r.Form.Get("judul")
	id := r.Form.Get("id")
	if judul != "" {
		db.Where("judul = ?", judul).First(&films)
	} else {
		if id != "" {
			db.Where("id = ?", id).First(&films)
		} else {
			db.Find(&films)
		}
	}

	// Set response
	var response model.FilmResponse
	if len(films) > 0 {
		// Output to console
		response.Status = 200
		response.Message = "Success Get User Data"
		response.Data = films
	} else {
		// Output to console
		response.Status = 204
		response.Message = "Not Found, No Content"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateFilm(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	var film model.Film

	id := r.Form.Get("id")

	db.Where("id = ?", id).First(&film)

	judul := r.Form.Get("judul")
	genre := r.Form.Get("genre")
	sutradara := r.Form.Get("sutradara")
	filmtype := r.Form.Get("filmtype")
	sinopsis := r.Form.Get("sinopsis")
	daftarpemain := r.Form.Get("daftarpemain")
	tahunrilis, _ := strconv.Atoi(r.Form.Get("tahunrilis"))

	if judul != film.Judul {
		db.Model(model.Film{}).Where("id = ?", id).Updates(model.Film{Judul: judul})
	}
	if genre != film.Genre {
		db.Model(model.Film{}).Where("id = ?", id).Updates(model.Film{Genre: genre})
	}
	if sutradara != film.Sutradara {
		db.Model(model.Film{}).Where("id = ?", id).Updates(model.Film{Sutradara: sutradara})
	}
	if filmtype != film.Filmtype {
		db.Model(model.Film{}).Where("id = ?", id).Updates(model.Film{Filmtype: filmtype})
	}
	if sinopsis != film.Sinopsis {
		db.Model(model.Film{}).Where("id = ?", id).Updates(model.Film{Sinopsis: sinopsis})
	}
	if daftarpemain != film.DaftarPemain {
		db.Model(model.Film{}).Where("id = ?", id).Updates(model.Film{DaftarPemain: daftarpemain})
	}
	if tahunrilis != film.TahunRilis {
		db.Model(model.Film{}).Where("id = ?", id).Updates(model.Film{TahunRilis: tahunrilis})
	}

	db.Where("id = ?", id).First(&film)
	var response model.FilmResponse
	if film.Judul == judul || film.Genre == genre || film.Sutradara == sutradara || film.Filmtype == filmtype || film.Sinopsis == sinopsis || film.DaftarPemain == daftarpemain || film.TahunRilis == tahunrilis {
		response.Status = 200
		response.Message = "Success Update Data"
	} else {
		response.Status = 400
		response.Message = "Failed Update Data"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func FindFilms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	var films []model.Film

	judul := r.Form.Get("judul")
	sutradara := r.Form.Get("sutradara")
	tahunrilis, _ := strconv.Atoi(r.Form.Get("tahunrilis"))
	sinopsis := r.Form.Get("sinopsis")
	daftarpemain := r.Form.Get("daftarpemain")

	if judul != "" {
		db.Where("judul LIKE ?", "%"+judul+"%").Find(&films)
	} else {
		if sutradara != "" {
			db.Where("sutradara LIKE ?", "%"+sutradara+"%").Find(&films)
		} else {
			if tahunrilis != 0 {
				db.Where("tahun_rilis = ?", tahunrilis).Find(&films)
			} else {
				if sinopsis != "" {
					db.Where("sinopsis LIKE ?", "%"+sinopsis+"%").Find(&films)
				} else {
					if daftarpemain != "" {
						db.Where("daftar_pemain LIKE ?", "%"+daftarpemain+"%").Find(&films)
					} else {
						db.Find(&films)
					}
				}
			}
		}
	}

	// Set response
	var response model.FilmResponse
	if len(films) > 0 {
		// Output to console
		response.Status = 200
		response.Message = "Success Get Films"
		response.Data = films
	} else {
		// Output to console
		response.Status = 204
		response.Message = "Not Found, No Content"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
