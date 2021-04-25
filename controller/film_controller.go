package controller

import (
	"net/http"
	"strconv"

	"github.com/Tugas_Besar/model"
	"gorm.io/gorm"
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
	filmtype, _ := strconv.Atoi(r.Form.Get("filmtype"))
	sinopsis := r.Form.Get("sinopsis")
	daftarpemain := r.Form.Get("daftarpemain")
	tahunrilis, _ := strconv.Atoi(r.Form.Get("tahunrilis"))

	// Set inputted data to object
	film := model.Film{
		Judul:        judul,
		Genre:        genre,
		Sutradara:    sutradara,
		FilmType:     filmtype,
		Sinopsis:     sinopsis,
		DaftarPemain: daftarpemain,
		TahunRilis:   tahunrilis,
	}

	// Insert object to database
	result := db.Create(&film)

	// Set response
	if result.Error == nil {
		// Output to console
		sendResponse(w, 200, "Success Insert Film to Database", nil)
	} else {
		// Output to console
		sendResponse(w, 400, "Insert Failed", nil)
	}
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

	var result *gorm.DB
	if judul != "" {
		result = db.Where("judul LIKE ?", "%"+judul+"%").Find(&films)
	} else if id != "" {
		result = db.Where("id = ?", id).First(&films)
	} else {
		result = db.Find(&films)
	}

	// Set response
	if result.Error != nil {
		sendResponse(w, 400, "Query Failed", nil)
	} else {
		// Output to console
		sendResponse(w, 200, "Success Get User Data", films)
	}
}

func UpdateFilm(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	id := r.Form.Get("id")
	judul := r.Form.Get("judul")
	genre := r.Form.Get("genre")
	sutradara := r.Form.Get("sutradara")
	filmtype, _ := strconv.Atoi(r.Form.Get("filmtype"))
	sinopsis := r.Form.Get("sinopsis")
	daftarpemain := r.Form.Get("daftarpemain")
	tahunrilis, _ := strconv.Atoi(r.Form.Get("tahunrilis"))

	result := db.Model(model.Film{}).Where("id = ?", id).Updates(model.Film{Judul: judul, Genre: genre, Sutradara: sutradara, FilmType: filmtype, Sinopsis: sinopsis, DaftarPemain: daftarpemain, TahunRilis: tahunrilis})

	if result.Error == nil {
		sendResponse(w, 200, "Success Update Data", nil)
	} else {
		sendResponse(w, 400, "Failed Update Data", nil)
	}
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
	tahunrilis := r.Form.Get("tahunrilis")
	genre := r.Form.Get("genre")
	sinopsis := r.Form.Get("sinopsis")
	daftarpemain := r.Form.Get("daftarpemain")
	query := ""

	// Query conditions
	if judul != "" {
		query += "judul LIKE '%" + judul + "%'"
	}

	if sutradara != "" {
		if judul != "" {
			query += " AND "
		}
		query += "sutradara LIKE '%" + sutradara + "%'"
	}

	if tahunrilis != "" {
		if judul != "" || sutradara != "" {
			query += " AND "
		}
		query += "tahun_rilis = " + tahunrilis
	}

	if genre != "" {
		if judul != "" || sutradara != "" || tahunrilis != "" {
			query += " AND "
		}
		query += "genre = " + genre
	}

	if sinopsis != "" {
		if judul != "" || sutradara != "" || tahunrilis != "" || genre != "" {
			query += " AND "
		}
		query += "sinopsis LIKE '%" + sinopsis + "%'"
	}

	if daftarpemain != "" {
		if judul != "" || sutradara != "" || tahunrilis != "" || genre != "" || sinopsis != "" {
			query += " AND "
		}
		query += "daftar_pemain LIKE '%" + daftarpemain + "%'"
	}

	result := db.Where(query).Find(&films)

	// Set response
	if result.Error != nil {
		sendResponse(w, 400, "Query Failed", nil)
	} else {
		// Output to console
		sendResponse(w, 200, "Success Get Films", films)
	}
}
