package handlers

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, object struct{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(object)
}

func DecodeBody(r http.Request, model *struct{}) error {
	err := json.NewDecoder(r.Body).Decode(model)
	if err != nil {
		return err
	}
	return nil
}
