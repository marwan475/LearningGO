package main

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func readJSON(w http.ResponseWriter, r *http.Request, data any) error {

	// max 1mb json data to prevent DDOS
	maxBytes := 1048578

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func writeJSONerror(w http.ResponseWriter, status int, msg string) error {

	type jsonError struct {
		Error string `json:"error"`
	}

	return writeJSON(w, status, &jsonError{Error: msg})
}
