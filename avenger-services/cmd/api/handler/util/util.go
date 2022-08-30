package util

import (
	"encoding/json"
	"log"
	"net/http"
)

// var client = &http.Client{}

type ResponseHandler struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func WebResponse(data interface{}, message string, status bool) []byte {
	jsonString := ResponseHandler{
		status,
		message,
		data,
	}
	result, err := json.MarshalIndent(jsonString, "", "    ")
	if err != nil {
		log.Print(err.Error())
	}
	return result
}

// ErrorResponse is a wrapper function for returning a web response that includes a standard text message.
func ErrorResponse(w http.ResponseWriter, r *http.Request, code int, data interface{}, message string) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // To
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(WebResponse(data, message, false))
}

// SuccessResponse is a wrapper function for returning a web response that includes a standard text message.
func SuccessResponse(w http.ResponseWriter, r *http.Request, code int, data interface{}, message string) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // To
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(WebResponse(data, message, true))
}
