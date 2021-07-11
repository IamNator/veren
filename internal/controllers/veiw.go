package controllers

import (
	"net/http"
)

//HomePageHandler serves the first page of the application
func HomePageHandler(w http.ResponseWriter, r *http.Request) {

	//	http.ServeFile(w, r, "/frontend/dist/index.html")

	http.ServeFile(w, r, "./frontend/dist")

}
