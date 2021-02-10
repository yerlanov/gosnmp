package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"test/exception"
	"test/service"
)

type Request struct {
	Login string `json:"login"`
}

func GetOperStatus(w http.ResponseWriter, r *http.Request) {
	var request Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		WriteError(w, r, exception.Error{ErrorType: exception.DecodeError, ErrorMessage: err})
		return
	}

	operStatus, serviceError := service.GetOperStatusService(request.Login)
	if serviceError.ErrorMessage != nil {
		WriteError(w, r, exception.Error{ErrorType: serviceError.ErrorType, ErrorMessage: serviceError.ErrorMessage})
		return
	}
	log.Printf("IP: %s, METHOD: %s, URL:%s, STATUS_CODE: %d, MESSAGE: %s", r.RemoteAddr, r.Method, r.URL, 200, "SUCCESS")
	_ = json.NewEncoder(w).Encode(operStatus)
}

func WriteError(w http.ResponseWriter, r *http.Request, error exception.Error) {
	var code int
	switch error.ErrorType {
	case exception.DecodeError:
		code = http.StatusBadRequest
	case exception.NotFound:
		code = http.StatusNotFound
	default:
		code = http.StatusInternalServerError
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	log.Printf("IP: %s, METHOD: %s, URL:%s, STATUS_CODE: %d, MESSAGE: %s", r.RemoteAddr, r.Method, r.URL, code, error.ErrorMessage)
	res := map[string]interface{}{"error": error.ErrorMessage.Error()}
	_ = json.NewEncoder(w).Encode(res)
}
