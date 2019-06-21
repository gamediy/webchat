package utils

import (
	"net/http"
	"encoding/json"
)

type JsonModel struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func JsonResp(w http.ResponseWriter, code int, message string,data interface{})  {
	bytes, _ := json.Marshal(JsonModel{
		Code: code,
		Message:message,
		Data:data,
	})
	w.Write(bytes)
}
