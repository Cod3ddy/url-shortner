package main

import "net/http"

func (app *application) home(w http.ResponseWriter, r *http.Request){
	data := app.newTemplateData(r)
	app.render(w, r, http.StatusOK, "home.html", data)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
	
}