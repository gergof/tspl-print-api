package main

import (
	"encoding/json"
	"net/http"
)

func jsonResponse(w http.ResponseWriter, code int, message string) {
	jsonResponseObject(w, code, map[string]string{
		"message": message,
	})
}

func jsonResponseObject(w http.ResponseWriter, code int, data map[string]string) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
