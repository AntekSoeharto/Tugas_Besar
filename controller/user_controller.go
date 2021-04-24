package controller

import (
	"encoding/json"
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

	var users []model.User

	email := r.Form.Get("email")
	if email != "" {
		db.Where("email = ?", email).First(&users)
	} else {
		db.Find(&users).Where("usertype = ?", 1)
	}

	// Set response
	var response model.UserResponse
	if len(users) > 0 {
		// Output to console
		response.Status = 200
		response.Message = "Success Get User Data"
		response.Data = users
	} else {
		// Output to console
		response.Status = 204
		response.Message = "Not Found, No Content"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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
	usermember := r.Form.Get("usermember")

	// Set inputted data to object
	user := model.User{
		Nama:         nama,
		Email:        email,
		Password:     password,
		TglLahir:     tgllahir,
		Jeniskelamin: jeniskelamin,
		Asalnegara:   asalnegara,
		Usertype:     usertype,
		Usermember:   usermember,
	}

	// Insert object to database
	result := db.Create(&user)

	// Set response
	var response model.UserResponse
	if result.Error == nil {
		// Output to console
		response.Status = 200
		response.Message = "Ragister Failed"
	} else {
		// Output to console
		response.Status = 400
		response.Message = "Register Failed"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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
	var response model.UserResponse
	if user.Usermember == "Ditangguhkan" {
		response.Status = 400
		response.Message = "Akun Anda Sedang Ditangguhkan"
	} else {
		if user.Email != "" {
			generateToken(w, user.Id, user.Nama, user.Usertype)
			fmt.Println(user.Usertype + 7)
			response.Status = 200
			response.Message = "Success Login"
		} else {
			response.Status = 204
			response.Message = "No Content (Email and Password doesn't match)"
		}
	}

	w.Header().Set("Content-Type", "application/json")
	// casting dari type go, ke json
	json.NewEncoder(w).Encode(response)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	resetUsersToken(w)

	var response model.UserResponse
	response.Status = 200
	response.Message = "Log-Out Success"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func TangguhkanMember(w http.ResponseWriter, r *http.Request) {
	db := connect()
	//defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	email := r.Form.Get("email")

	db.Model(model.User{}).Where("email = ?", email).Updates(model.User{Usermember: "Ditangguhkan"})

	var user model.User
	db.Where("email = ?", email).First(&user)

	var response model.UserResponse
	if user.Usermember == "Ditangguhkan" {
		response.Status = 200
		response.Message = "Berhasil Menangguhkan"
	} else {
		response.Status = 400
		response.Message = "Gagal Menangguhkan"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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
	if jeniskelamin != user.Jeniskelamin {
		db.Model(model.User{}).Where("id = ?", id).Updates(model.User{Jeniskelamin: jeniskelamin})
	}

	db.Where("id = ?", id).First(&user)
	var response model.FilmResponse
	if user.Nama == nama || user.TglLahir == tgllahir || user.Jeniskelamin == jeniskelamin {
		response.Status = 200
		response.Message = "Success Update Data"
	} else {
		response.Status = 200
		response.Message = "Success Update Data"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendUnAuthorizedResponse(w http.ResponseWriter) {
	var response model.UserResponse
	response.Status = 401
	response.Message = "Unauthorized Access"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
