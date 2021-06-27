package apiHelpers

import (
	"encoding/json"
	"net/http"
	tools "restful-api-kit/utilities"
)

//ResponseData structure
type (
	ResponseData struct {
		Data interface{} `json:"data"`
		Meta interface{} `json:"meta"`
	}
	Response struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

func MakeRes(code int, data ...interface{}) *Response {
	res := &Response{
		Code:    code,
		Message: CODE2MSG[code],
		Data:    data,
	}
	return res
}

//Message returns map data
func Message(status int, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

//Respond returns basic response structure
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	tools.CheckErr(err)
}

func RespondString(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	tools.CheckErr(err)
}
