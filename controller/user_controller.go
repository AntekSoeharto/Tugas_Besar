package controller

import (
	"net/http"

	"github.com/Tugas_Besar/model"
	"gorm.io/gorm"
)

func GetMember(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	email := r.Form.Get("email")

	var users []model.User
	var result *gorm.DB
	if email != "" {
		result = db.Where("email = ?", email).Preload("Langganan").First(&users)
	} else {
		result = db.Find(&users)
	}

	// Set response
	if result.Error != nil {
		// Output to console
		sendResponse(w, 400, "Query Failed", nil)
	} else if len(users) < 1 {
		// Output to console
		sendResponse(w, 204, "Not Found, No Content", nil)
	} else {
		sendResponse(w, 200, "Success Get User Data", users)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	// Get data from user
	nama := r.Form.Get("nama")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	tgllahir := r.Form.Get("tgllahir")
	jeniskelamin := r.Form.Get("jeniskelamin")
	asalnegara := r.Form.Get("asalnegara")

	// Set inputted data to object
	user := model.User{
		Nama:         nama,
		Email:        email,
		Password:     password,
		TglLahir:     tgllahir,
		JenisKelamin: jeniskelamin,
		AsalNegara:   asalnegara,
		UserType:     1,
	}

	// Insert object to database
	result := db.Create(&user)

	// Set response
	if result.Error == nil {
		// Output to console
		sendResponse(w, 200, "Register Success", nil)
	} else {
		// Output to console
		sendResponse(w, 400, "Register Failed", nil)
	}
}

func LogIn(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	var user model.User
	result := db.Where("email = ? and password = ?", email, password).First(&user)

	// Set response
	if result.Error != nil {
		sendResponse(w, 400, "Query Failed", nil)
	} else if user.ID == 0 {
		sendResponse(w, 204, "No Content (Email and Password doesn't match)", nil)
	} else if user.Block == 1 {
		sendResponse(w, 403, "Akun Anda Sedang Ditangguhkan", nil)
	} else {
		generateToken(w, user.ID, user.Nama, user.UserType)
		sendResponse(w, 200, "Success Login", nil)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	resetUsersToken(w)
	sendResponse(w, 200, "Log-Out Success", nil)
}

func TangguhkanMember(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	id := r.Form.Get("id")

	result := db.Model(model.User{}).Where("id = ?", id).Updates(model.User{Block: 1})

	if result.Error != nil {
		sendResponse(w, 200, "Berhasil Menangguhkan", nil)
	} else {
		sendResponse(w, 400, "Gagal Menangguhkan", nil)
	}
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	var user model.User
	id := getid(r)
	db.Where("id = ?", id).First(&user)

	nama := r.Form.Get("nama")
	tgllahir := r.Form.Get("tgllahir")
	jeniskelamin := r.Form.Get("jeniskelamin")

	result := db.Model(&user).Where("id = ?", id).Updates(model.User{Nama: nama, TglLahir: tgllahir, JenisKelamin: jeniskelamin})

	if result.Error == nil {
		sendResponse(w, 200, "Success Update Data", nil)
	} else {
		sendResponse(w, 400, "Failure Update Data", nil)
	}
}
