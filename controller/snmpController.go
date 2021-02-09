package controller

import (
	"encoding/json"
	"net/http"
	"test/service"
)

type Request struct {
	Login string `json:"login"`
}

func GetOperStatus(w http.ResponseWriter, r *http.Request) {
	var request Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		WriteError(w, err)
		return
	}

	operStatus, err := service.GetOperStatusService(request.Login)
	if err != nil {
		WriteError(w, err)
		return
	}
	_ = json.NewEncoder(w).Encode(operStatus)
}

func WriteError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	res := map[string]interface{}{"error": err.Error()}
	_ = json.NewEncoder(w).Encode(res)
}
