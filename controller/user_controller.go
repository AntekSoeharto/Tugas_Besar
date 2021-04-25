package controller

import (
	"fmt"

	"net/http"

	"github.com/Tugas_Besar/model"
)

func GetMember(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	var user model.User
	var users []model.User
	var langganan model.Langganan

	email := r.Form.Get("email")
	db.Where("email", email).First(&user)
	db.Where("id_user", user.ID).First(&langganan)
	user.Langganan = langganan
	users = append(users, user)

	// Set response
	if len(users) > 0 {
		// Output to console
		sendResponse(w, 200, "Success Get User Data", users)
	} else {
		// Output to console
		sendResponse(w, 204, "Not Found, No Content", nil)
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
	usertype := 1

	// Set inputted data to object
	user := model.User{
		Nama:         nama,
		Email:        email,
		Password:     password,
		TglLahir:     tgllahir,
		JenisKelamin: jeniskelamin,
		AsalNegara:   asalnegara,
		UserType:     usertype,
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

	// We can chaining method like this
	db.Where("email = ? and password = ?", email, password).First(&user)

	// Set response
	if user.Block == 1 {
		sendResponse(w, 400, "Akun Anda Sedang Ditangguhkan", nil)
	} else {
		if user.Email != "" {
			generateToken(w, user.ID, user.Nama, user.UserType)
			fmt.Println(user.UserType + 7)
			sendResponse(w, 200, "Success Login", nil)
		} else {
			sendResponse(w, 204, "No Content (Email and Password doesn't match)", nil)
		}
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

	db.Model(model.User{}).Where("id = ?", id).Updates(model.User{Block: 1})

	var user model.User
	db.Where("id = ?", id).First(&user)

	if user.Block == 1 {
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

	if nama != user.Nama {
		db.Model(model.User{}).Where("id = ?", id).Updates(model.User{Nama: nama})
	}
	if tgllahir != user.TglLahir {
		db.Model(model.User{}).Where("id = ?", id).Updates(model.User{TglLahir: tgllahir})
	}
	if jeniskelamin != user.JenisKelamin {
		db.Model(model.User{}).Where("id = ?", id).Updates(model.User{JenisKelamin: jeniskelamin})
	}

	db.Where("id = ?", id).First(&user)
	if user.Nama == nama || user.TglLahir == tgllahir || user.JenisKelamin == jeniskelamin {
		sendResponse(w, 200, "Success Update Data", nil)
	} else {
		sendResponse(w, 200, "Success Update Data", nil)
	}
}
