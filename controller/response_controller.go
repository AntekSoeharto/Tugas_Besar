package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Tugas_Besar/model"
)

func sendResponse(w http.ResponseWriter, status int, msg string, data interface{}) {
	var response model.ResponseData
	response.Status = status
	response.Message = msg
	response.Data = data

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendUnAuthorizedResponse(w http.ResponseWriter) {
	var response model.ResponseData
	response.Status = 401
	response.Message = "Unauthorized Access"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
