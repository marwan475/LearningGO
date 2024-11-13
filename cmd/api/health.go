package main

import (
	"net/http"
)

func (app *application) CheckHealth(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}

	err := writeJSON(w, http.StatusOK, data)

	if err != nil {
		writeJSONerror(w, http.StatusInternalServerError, "CheckHealth Error")
	}

}
