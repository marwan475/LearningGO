package main

import "net/http"

func (app *application) CheckHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
